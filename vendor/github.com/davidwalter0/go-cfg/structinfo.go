package cfg

import (
	"fmt"
	"github.com/davidwalter0/go-cfg/flag"
	"log"
	"os"
	"reflect"
	"strings"
)

// AppEnvVarPrefixOverrideName environment variable application lookup
// prefix override, default prefix is the struct name
var AppEnvVarPrefixOverrideName = "APP_OVERRIDE_PREFIX"

var helpText string

// HelpText pre write text for help
func HelpText(text string) {
	helpText = text
}

// Usage add usage text for flags/help processing
var Usage = func() {
	var parts = strings.Split(os.Args[0], "/")
	var l = len(parts)
	var Program = parts[l-1]
	fmt.Fprintf(os.Stderr, "\nUsage of %s:\n", Program)
	if len(helpText) > 0 {
		fmt.Fprintf(os.Stderr, "\n%s\n\n", helpText)
	}
	flag.PrintDefaults()
}

func init() {
	flag.Usage = Usage
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "This is not helpful.\n")
	// }
}

// AddPrefixDeferFreeze initialize struct but don't freeze flags
func AddPrefixDeferFreeze(prefix string, sptr interface{}) (err error) {
	return ProcessHoldFlags(prefix, sptr)
}

// AddStruct bootstrap the configuration from environment and flags to
// struct with env var name override with empty prefix, defer freeze
// of flags
func AddStruct(sptr interface{}) (err error) {
	return ProcessHoldFlags("", sptr)
}

var finalized = false

// Finalize the configuration
func Finalize() {
	if !finalized {
		Freeze()
		finalized = true
	}
}

// Process bootstrap the configuration from environment and flags to
// struct with env var name override to replace the prefix of the
// object name. The prefix argument will replace/override the struct
// type name, to default to the use of the struct name call the Parse
// method directly
func Process(prefix string, sptr interface{}) (err error) {
	var sti *StructInfo = &StructInfo{
		StructPtr:    sptr,
		EnvVarPrefix: prefix,
		EmptyPrefix:  len(prefix) == 0,
	}

	if err = sti.Parse(); err != nil { // parse tags, environment, flags
		log.Println(*sti)
		return
	}
	if announceDuplicates {
		fmt.Println()
		os.Exit(1)
	}
	return
}

// ProcessHoldFlags bootstrap the configuration from environment and flags to
// struct with env var name override to replace the prefix of the
// object name. The prefix argument will replace/override the struct
// type name, to default to the use of the struct name call the Parse
// method directly
func ProcessHoldFlags(prefix string, sptr interface{}) (err error) {
	var sti *StructInfo = &StructInfo{
		StructPtr:    sptr,
		EnvVarPrefix: prefix,
		EmptyPrefix:  len(prefix) == 0,
	}

	// parse tags, environment, hold off on flags
	if err = sti.ParseHoldFlags(); err != nil {
		return
	}
	if announceDuplicates {
		fmt.Println()
		os.Exit(1)
	}
	return
}

// Freeze flags with current set
func Freeze() {
	flag.Parse()
}

// Parse bootstrap the configuration from environment and flags
// to struct with env var name override to replace the prefix of the
// object name
func Parse(sptr interface{}) (err error) {
	var sti *StructInfo = &StructInfo{
		StructPtr: sptr,
	}

	if err = sti.Parse(); err != nil { // parse tags, environment, flags
		log.Printf("%v\n", err)
	}
	return err
}

// Parse the struct and flags initializing configuration from tags,
// environment, and flags in order, additional flags must be defined
// prior to Parse call as it calls flag.Parse
func (structInfo *StructInfo) Parse() (err error) {
	if !structInfo.Processed {
		structInfo.Processed = true
		if err = structInfo.process(); err != nil {
			return
		}
		flag.Parse()
	} else {
		panic("Parse called more than once")
	}
	return
}

// ParseHoldFlags the struct and prep flags, but don't call
// flags.Parse, initializing configuration from tags, environment, and
// flags in order, additional flags must be defined prior to Parse
// call as it calls flag.Parse
func (structInfo *StructInfo) ParseHoldFlags() (err error) {
	if !structInfo.Processed {
		structInfo.Processed = true
		if err = structInfo.process(); err != nil {
			return
		}
	} else {
		panic("Parse called more than once")
	}
	return
}

// process initialize with *struct (struct pointer) and use the struct
// name as the app name for the env prefix
// func (structInfo *StructInfo) Init(structPtr interface{}) (err error) {
func (structInfo *StructInfo) process() (err error) {
	if reflect.TypeOf(structInfo.StructPtr).Kind() != reflect.Ptr {
		err = ErrInvalidArgPointerRequired
	} else {
		var depth = 0
		var ok bool
		m := reflect.ValueOf(structInfo.StructPtr).Elem()
		switch m.Kind() {
		case reflect.Struct:
			element := reflect.ValueOf(structInfo.StructPtr).Elem()
			elementType := element.Type()
			AppName := elementType.Name()
			if !structInfo.EmptyPrefix {
				if len(structInfo.EnvVarPrefix) == 0 {
					if structInfo.EnvVarPrefix, ok = lookupEnv(AppEnvVarPrefixOverrideName); !ok {
						structInfo.EnvVarPrefix = elementType.Name()
					} else {
						AppName = structInfo.EnvVarPrefix
					}
				}
			}
			prefix := structInfo.EnvVarPrefix

			for i := 0; i < elementType.NumField(); i++ {
				structField := elementType.Field(i)
				ptr := element.Field(i).Addr().Interface()
				var member = &MemberType{
					AppName:      AppName,
					EnvVarPrefix: prefix,
				}
				if err = member.Parse(prefix, structField, ptr, depth); err != nil {
					return
				}
			}
		}
	}
	return
}
