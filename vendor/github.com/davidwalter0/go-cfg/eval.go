/*
Package cfg provides facilities to ingest structs as configuration
elements via reflection. Struct names must be exported. Beyond that
recursive structs are supported with options prefixing with the
outermost struct name. Name overrides via struct tags can be used
to rename configuration options as well as (optionally) set
defaults.


 Configurations share these options

  - A configuration is one or more struct pointers
  - Each struct member's data type is used for type conversion during assignment
  - Primitive types are evaluated directly
  - Aggregate slice can be set with comma delibmited syntax
  - Aggregate map can be set with comma delibmited syntax

  map example
	Map  map[string]float64 `short:"m" default:"π:3.14159,ξ:1,ρ:.01,φ:1.2,β:3,α:.01,δ:3,ε:.001,φ:.1,ψ:.9,ω:2.1"`


  - Decorate

Modes of configuration


  - Wrap(prefix, ptrs...) options with a prefix argument to each flag or environment variable
  - Nest(ptrs...) each struct member with the name of the the parent struct(s)
  - Unwrap( ptr ... interface{} ) error
  - Flags and Env Vars are unwrapped - no prefix or argument name added
  - Call with one or more structures with uniquely named members
  - One flag + env variable for each entry named
  - When called with duplicate members in one or more structs flags and
    env vars will conflict and error
*/
package cfg

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"strconv"
)

// decorate var setting, disable prefixes when undecorated
var decorate bool = true

// Arg is the settings passed for recursion
type Arg struct {
	Depth    int
	Prefix   string
	Prefixed bool
	UseFlags bool
}

// NewArg sets values in Arg for cfg to process structs
func NewArg(name string) *Arg {
	var prefixed bool
	var prefix string
	var text string
	var ok bool
	if text, ok = LookupEnv(cfgDecorate); ok {
		if v, err := strconv.ParseBool(text); err != nil {
			log.Println(err)
		} else {
			decorate = v
		}
	}
	if decorate {
		prefix, ok = LookupEnv(cfgEnvKeyPrefix)
		if ok && len(prefix) > 0 {
			prefix = prefix + "_"
			prefixed = true
		}
		if len(name) > 0 {
			prefixed = true
			prefix = prefix + name
		}
	}
	return &Arg{Depth: 0, Prefixed: prefixed, Prefix: prefix, UseFlags: true}
}

// Undecorate structs with prefix
func Undecorate() bool {
	decorate = false
	os.Setenv(cfgDecorate, "false")
	return decorate
}

// Decorate structs with prefix
func Decorate() bool {
	decorate = true
	os.Setenv(cfgDecorate, "true")
	return decorate
}

func Unprefix() {
	os.Unsetenv(cfgEnvKeyPrefix)
}

// Eval one or more configuration structures
func Eval(ptrs ...interface{}) error {
	args := NewArg("")
	return Run(args, ptrs...)
}

// Init from a list of struct pointers
func Init(ptrs ...interface{}) error {
	args := NewArg("")
	err := Run(args, ptrs...)
	Freeze()
	return err
}

// EvalName one or more configuration structures overriding the name
func EvalName(name string, ptrs ...interface{}) error {
	args := NewArg(name)
	err := Run(args, ptrs...)
	Freeze()
	return err
}

func Run(args *Arg, ptrs ...interface{}) error {
	var err error
	err = CheckArgs(ptrs...)
	if err != nil {

		os.Exit(1)
	}
	for _, ptr := range ptrs {
		err = Enter(args, ptr)
		if err != nil {
			return err
		}
	}
	return err
}

// SimpleFlags don't prefix with base struct prefix, and they are
// undecorated.
//
// type T struct {
//    I int
// }
// var ptr = &T{}
// Default
// Eval(ptr) is decorated/prefixed with the struct name
// type T struct {
// ...
//    I int
// }
// type T struct {
// Eval
// EvalName("D", ptr) is decorated/prefixed with flag "--d-i"
// EvalName("D", ptr) is decorated/prefixed with env var "D_I"
//
// export CFG_KEY_PREFIX=KP
// Eval(p) is decorated/prefixed with flag "--kp-d-i"
// EvalName("D", p) is decorated/prefixed with env var "KP_D_I"
//
// export CFG_KEY_PREFIX=KP
// export CFG_DECORATE=false
// ignore prefix and struct name prefixing
// Eval(p) is decorated/prefixed with flag "--kp-d-i"
// EvalName("D", p) is decorated/prefixed with env var "KP_D_I"

// Simple create env vars prefices
func Simple(ptrs ...interface{}) error {
	Unprefix()
	Undecorate()
	args := NewArg("")
	return Run(args, ptrs...)
}

// SimpleFlags create env vars and flags without prefices
func SimpleFlags(ptrs ...interface{}) error {
	err := Simple(ptrs...)
	Freeze()
	return err
}

// Unwrap alias of Simple create env vars without prefices
func Unwrap(ptrs ...interface{}) error {
	Unprefix()
	Undecorate()
	args := NewArg("")
	err := Run(args, ptrs...)
	return err
}

// Final freezes calling flag.Parse, no more additions to the
// configuration after Final
func Final() {
	Freeze()
}

// Bare alias of Simple create env vars without prefices
func Bare(ptrs ...interface{}) error {
	Unprefix()
	Undecorate()
	args := NewArg("")
	err := Run(args, ptrs...)
	return err
}

// Wrap (optionally) with prefix and struct names create env vars without prefices
func Wrap(name string, ptrs ...interface{}) error {
	args := NewArg(name)
	err := Run(args, ptrs...)
	return err
}

// Add alias of Eval
func Add(ptrs ...interface{}) error {
	args := NewArg("")
	return Run(args, ptrs...)
}

var Package = func() string {
	type Empty struct{}
	return reflect.TypeOf(Empty{}).PkgPath()
}()

// Flags alias of Eval
func Flags(ptrs ...interface{}) error {
	args := NewArg("")
	err := Run(args, ptrs...)
	Freeze()
	return err
}

// Nest uses struct names recursively to prefix each flag or
// environment variable name as it descends the configuration object
// heirarchy. For exqmple, type B struct{ I int }; type A struct{ B }
// will create a flag with prefixes a and b, i.e. --a-b-i
// --a-b-i int
//   	 Env _A_B_I                           : (I) (int)

func Nest(ptrs ...interface{}) error {
	args := NewArg("")
	args.Prefixed = true
	err := Run(args, ptrs...)
	return err
}

// NestWrap objects retaining object hierarchy with prefix
func NestWrap(prefix string, ptrs ...interface{}) error {
	// fmt.Println("NestWrap", prefix)
	// fmt.Println(ptrs...)
	args := NewArg(prefix)
	args.Prefixed = true
	err := Run(args, ptrs...)
	return err
}

// // https://graphemica.com/%E2%9C%93
// // https://graphemica.com/%E2%9C%97
// log.Println("\n\u2713 text\n\u2716 text\n")

// CheckArgs validate that the pointers are pointers to struct
func CheckArgs(ptrs ...interface{}) error {
	checkArgs := []string{}
	ok := true
	for i, ptr := range ptrs {
		typeOf := reflect.Indirect(reflect.ValueOf(ptr)).Type()
		kind := reflect.TypeOf(ptr).Kind()
		switch kind {
		case reflect.Ptr:
			elem := reflect.TypeOf(ptr).Elem()
			isPtr := reflect.Ptr == elem.Kind()
			if isPtr {
				ok = false
				checkArgs = append(checkArgs, fmt.Sprintf("%s !ok %2d [%v] is a pointer to %v.", Red("\u2716"), i, reflect.TypeOf(ptr), typeOf))
			} else {
				checkArgs = append(checkArgs, fmt.Sprintf("%s  ok %2d [%v] is a pointer to %v. expected.", Green("\u2713"), i, reflect.TypeOf(ptr), typeOf))
			}
		default:
			ok = false
			name := reflect.TypeOf(ptr).Name()
			checkArgs = append(checkArgs, fmt.Sprintf("%s !ok %2d [%v] is not a pointer.", Red("\u2716"), i, name))
		}
	}
	var err error
	var str string
	if !ok {
		for _, text := range checkArgs {
			str += fmt.Sprintf(text)
		}
		str += fmt.Sprintf("\n%s\n\nThe argument list ptrs []interface{} is a slice of struct pointers (*struct) \n\n", Red(ErrInvalidArgPointerRequired))
		// fmt.Println(CallerAndArgs(ptrs))
		str += fmt.Sprintln(Caller())
		err = fmt.Errorf(str)
	}
	return err
}

// Caller get the calling frame info
func Caller() string {
	var pcs [1]uintptr
	n := runtime.Callers(2, pcs[:])
	text := ""
	{
		var pcs [10]uintptr
		n = runtime.Callers(2, pcs[:])
		frames := runtime.CallersFrames(pcs[:n])
		var frame runtime.Frame
		var more bool = true
		for i := 0; i < n-2 && more; i++ {
			frame, more = frames.Next()
			text += fmt.Sprintf("%s:%d:%s\n", frame.File, frame.Line, frame.Func.Name())
		}
	}
	return text
}
