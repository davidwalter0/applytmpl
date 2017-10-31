package main

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"strconv"
	"strings"
	"text/template"
)

var TemplateFunctions = template.FuncMap{
	"add":            Add,
	"atoi":           Atoi,
	"base64Decode":   Base64Decode,
	"base64Encode":   Base64Encode,
	"cat":            Cat,
	"curl":           Curl,
	"delimit":        Delimit, // replace space with ,
	"div":            Div,
	"env":            Env,
	"file":           File,
	"first":          First,
	"generate":       Generate,
	"generateInt":    GenerateInt,
	"generateN":      GenerateN,
	"charAGenerator": GeneratorChar,
	"intAGenerator":  Generator,
	"join":           Join,
	"get":            HTTPGet,
	"in":             In,
	"index":          Index,
	"lower":          Lower,
	"mod":            Mod,
	"mult":           Mult,
	"nth":            Nth,
	"set":            Set,
	"split":          Split,
	"sub":            Sub,
	"tostring":       ByteArrayToString,
	"trim":           Trim,
	"upper":          Upper,
	"upperCase":      UCase,
	"zip":            Zip,
	"zipPrefix":      ZipPrefix,
	"zipSuffix":      ZipSuffix,
	"zipprefix":      ZipPrefix,
	"zipsuffix":      ZipSuffix,
}

var debug bool
var debugText string

func trace() {
	pc := make([]uintptr, 10)
	runtime.Callers(10, pc)
	for i := 0; i < 10; i++ {
		if pc[i] == 0 {
			break
		}
		f := runtime.FuncForPC(pc[i])
		file, line := f.FileLine(pc[i])
		fmt.Printf("error: %s:%d: %s\n", file, line, f.Name())
	}
}

func isHttps(url string) bool {
	return len(strings.Split(url, "https://")) > 1
}

func recoverWithMessage(step string, exitOnException bool, failureExitCode int) {
	if r := recover(); r != nil {
		fmt.Printf("error: Recovered step[%s] with info\n-----\n%v\n-----\n", step, r)
		trace()
		pc := make([]uintptr, 10)
		runtime.Callers(5, pc)
		f := runtime.FuncForPC(pc[1])
		file, line := f.FileLine(pc[1])
		fmt.Printf("error: %s:%d: %s call failed at or near\n", file, line, f.Name())
		if exitOnException {
			os.Exit(failureExitCode)
		}
	}
}

// HTTPGet return text for uri
func HTTPGet(uri string) (text []byte, err error) {
	defer recoverWithMessage("HTTPGet", false, 5)
	var response *http.Response
	if debug {
		log.Printf("uri: %v\n>%v\n", uri, err)
	}
	if isHttps(uri) {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		response, err = client.Get(uri)
	} else {
		response, err = http.Get(uri)
	}

	if response != nil && (response.StatusCode < 200 || response.StatusCode > 399) {
		if debug {
			log.Fatalf("uri: %v\n>%v %v %v\n", uri, response.Status, response.StatusCode, err)
		}
		log.Fatalf("uri: %v\n>%v %v %v\n", uri, response.Status, response.StatusCode, err)
	}
	if err != nil {
		log.Fatalf("uri: %v\n>%v\n", uri, err)
	} else {
		defer response.Body.Close()
		text, err = ioutil.ReadAll(response.Body)
		if err != nil {
			if debug {
				log.Fatalf("uri: %v\n>%v\n", uri, err)
			}
			log.Fatalf("uri: %v\n>%v\n", uri, err)
			os.Exit(1)
		}
	}
	return
}

// Base64Encode transform input string to base64 encoded data
func Base64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

// Base64Decode transform input from base64 to string
func Base64Decode(text string) string {
	lhs, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		log.Fatalf("base64 decode error: %v\n>%v\n", text, err)
	}
	return string(lhs)
}

// Split a string to an array of strings on a space character
func Split(text string) []string {
	return strings.Split(Trim(text), " ")
}

// Join an array of strings to a single string delimited by a space
func Join(text []string) string {
	return strings.Join(text, " ")
}

// First item in an array split on spaces, using Split
func First(text string) string {
	array := Split(text)

	if len(array) > 0 {
		text = array[0]
	} else {
		text = ""
	}
	return text
}

// Nth zero offset item in the array after the text argument is split on spaces
func Nth(nstr string, text string) string {
	n, _ := strconv.Atoi(nstr)
	array := Split(text)
	if len(array) > n {
		return array[n]
	}
	return ""
}

// Trim spaces from a string
func Trim(text string) string {
	return strings.Trim(text, " ")
}

// Delimit a space separated string with delimiter [ default comma ',' ]
func Delimit(text string, delimiter string) (o string) {
	if len(delimiter) == 0 {
		delimiter = ","
	}
	array := Split(Trim(text))
	for i, x := range array {
		if i > 0 {
			o += delimiter
		}
		o += x
	}
	return o
}

// Zip 2 space separated lists with a separator char like "."
// "a b c" "1 2 3" "." -> "a.1 a.2 a.3 b.1 b.2 b.3"
// split list1 list2 and append with separator
func Zip(list1, list2, separator string) string {
	l1 := strings.Split(Trim(list1), " ")
	l2 := strings.Split(Trim(list2), " ")
	if len(separator) == 0 {
		separator = "-"
	}
	text := ""
	for _, x := range l2 {
		for j, y := range l1 {
			if j < len(l1) {
				text += " "
			}
			text += x + separator + y
		}
	}
	return text
}

// Index return the array index of [find] from in the text
func Index(find, in string) (text string) {
	array := strings.Split(Trim(in), " ")
	for i, x := range array {
		if find == x {
			text = strconv.Itoa(i)
			break
		}
	}
	return text
}

// ZipPrefix split text on space and zip with prefix
// "a b c" "node" "-" -> node-a node-b node-c
func ZipPrefix(text, prefix, separator string) []string {
	if len(separator) == 0 {
		separator = "-"
	}
	array := Split(Trim(text))
	text = ""
	for i, x := range array {
		array[i] = prefix + separator + x
	}
	return array
}

// ZipSuffix split text on space and zip with suffix
// "a b c" "node" "-" -> a-node b-node c-node
func ZipSuffix(text, suffix, separator string) []string {
	if len(separator) == 0 {
		separator = "-"
	}
	array := Split(Trim(text))

	for i, x := range array {
		array[i] = x + separator + suffix
	}
	return array
}

// Cat concatenate sequence of strings
func Cat(in ...string) string {
	text := ""
	for _, x := range in {
		text += x
	}
	return text
}

// errorStrings list of problems during template processing
var errorStrings []string

// Env lookup name return value
func Env(name string) (text string) {
	text = os.Getenv(name)
	if len(text) == 0 {
		if len(errorStrings) == 0 {
			errorStrings = append(errorStrings, fmt.Sprintf("Template Processing Error"))
		}
		errorStrings = append(errorStrings, fmt.Sprintf("env var unset: %s", name))
	}
	return
}

// File name loaded to byte array
func File(name string) []byte {
	return Load(name)
}

// ByteArrayToString from byte array
func ByteArrayToString(bytes []byte) string {
	return string(bytes)
}

// Load helper function from file to []byte
func Load(filename string) []byte {
	if len(filename) == 0 {
		panic(fmt.Sprintf("Can't Load() a file with an empty name"))
	}

	var err error
	var text []byte
	if len(filename) > 0 {
		text, err = ioutil.ReadFile(filename)
		if err != nil {
			log.Printf("%v.\n", err)
			os.Exit(3)
		}
	}
	return text
}

// GenerateInt an integer array from [0..n]
func GenerateInt(n int) (ints []int) {
	var i int
	ints = make([]int, 0)
	if i >= 0 {
		for i = 0; i < n; i++ {
			ints = append(ints, i)
		}
	}

	return
}

// Generate an integer array from [0..n] for
// consistent name extension use
func Generate(n int, zerofill bool) (text []string) {
	var i int
	suffix := " "
	text = make([]string, 0)
	if i >= 0 {
		var width = int(math.Log10(float64(n))) + 1
		for i = 0; i < n; i++ {
			if zerofill {
				text = append(text, fmt.Sprintf("%0.*d%s", width, i, suffix))
			} else {
				text = append(text, fmt.Sprintf("%d%s", i, suffix))
			}
			if i == n-1 {
				suffix = ""
			}
		}
	}
	return
}

// Generate an array of from [0..n] for consistent name extension use
// func Generate(n string) (text []string) {
//   return GenerateN(n)
// }
// GenerateN an array of from [0..n] for consistent name extension use
func GenerateN(nS string) (text []string) {
	n, _ := strconv.Atoi(nS)
	var i int
	text = make([]string, 0)
	if i >= 0 {
		for i = 0; i < n; i++ {
			text = append(text, fmt.Sprintf("%d", i))
		}
	}

	return
}

// Curl pulls a value using http(s)
func Curl(name string) string {
	defer recoverWithMessage("Curl", false, 4)
	bytes, err := HTTPGet(name)
	if err != nil {
		fmt.Println(err)
		panic(fmt.Sprintf("%v", err))
	}
	return string(bytes)
}

// Atoi convert a string to a base 10 integer
func Atoi(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Printf("problem %v\n", err)
	}
	return int(n)
}

// UCase the first character of string
func UCase(s string) string {
	s = Trim(s)
	if len(s) > 0 {
		return strings.ToUpper(s[0:1]) + s[1:]
	}
	return s
}

// Lower downcase string
func Lower(s string) string {
	s = Trim(s)
	if len(s) > 0 {
		return strings.ToLower(s)
	}
	return s
}

// Upper upcase string
func Upper(s string) string {
	s = Trim(s)
	if len(s) > 0 {
		return strings.ToUpper(s)
	}
	return s
}

// In array returns find if present, else return an empty string
// Calling split on a string converts to an array to preprocess the
// string for array operations like In.
// {{ split "a b c"|in "a" }} returns a
func In(find string, in []string) string {
	for _, x := range in {
		if find == x {
			return find
		}
	}
	return ""
}

var envMap = template.FuncMap{
	"env": Env,
}

// func SplitDigits(lhs, rhs string) (l, r int) {
// 	l, _ = strconv.Atoi(lhs)
// 	r, _ = strconv.Atoi(rhs)
// 	return
// }

func Add(l, r int) string {
	//	l, r := SplitDigits(lhs, rhs)
	return fmt.Sprintf("%d", l+r)
}
func Sub(l, r int) string {
	//	l, r := SplitDigits(lhs, rhs)
	return fmt.Sprintf("%d", l-r)
}
func Div(l, r int) string {
	//	l, r := SplitDigits(lhs, rhs)
	return fmt.Sprintf("%d", l/r)
}
func Mult(l, r int) string {
	//	l, r := SplitDigits(lhs, rhs)
	return fmt.Sprintf("%d", l*r)
}
func Mod(l, r int) string {
	//	l, r := SplitDigits(lhs, rhs)
	return fmt.Sprintf("%d", l%r)
}

var Environment map[string]string = make(map[string]string, 0)

func loadEnv() {
	env_array := os.Environ()
	for _, env := range env_array {
		parts := strings.SplitN(env, "=", 2)
		k, v := string(parts[0]), string(parts[1])
		Environment[k] = v

	}
}

// upperCase text
func upperCase(text string) string {
	var upperCased = text
	if len(text) > 0 {
		upperCased = strings.ToUpper(text[0:1]) + strings.ToLower(text[1:])
	}
	return upperCased
}

// downCase text
func downCase(text string) string {
	var downCased = Trim(text)
	if len(text) > 0 {
		downCased = strings.ToLower(text)
	}
	return downCased
}

// Set a var in the internal map
// func Set(k, v string) []string {
func Set(k, v string) string {
	k, v = Trim(k), Trim(v)
	EnvironmentKV[k] = v
	// return []string{k, v}
	return ""
}

// LCamelCase string delimited by "_" AA_bb_Cc return aaBbCc
func LCamelCase(arg string) (text string) {
	words := strings.Split(Trim(arg), "_")
	for i, word := range words {
		if i == 0 {
			text += downCase(word)
		} else {
			text += upperCase(word)
		}
	}
	return
}

// UCamelCase string delimited by "_" AA_bb_Cc return AaBbCc
func UCamelCase(arg string) (text string) {
	words := strings.Split(Trim(arg), "_")
	for _, word := range words {
		text += upperCase(word)
	}
	return
}

func ToString(k interface{}) (s string, e error) {
	switch k.(type) {
	case string:
		s = k.(string)
	case int32, int64, int:
		s = strconv.FormatInt(k.(int64), 10)
	case uint32, uint64, uint:
		s = strconv.FormatUint(k.(uint64), 10)
	case bool:
		s = strconv.FormatBool(k.(bool))
	case float32:
		s = strconv.FormatFloat(k.(float64), 'f', -1, 32)
	case float64:
		s = strconv.FormatFloat(k.(float64), 'f', -1, 32)
	case interface{}:
		s = fmt.Sprintf("%v", k)
	default:
		text := fmt.Sprintf("Expected one of [string, {u,}int{,32,64} float{32,64}] but got: %T", k)
		e = errors.New(text)
	}
	return s, e
}

// davidwalter0/api-driver/dispatch/jpath.go
func ToInt(k interface{}) (i int) {
	var s string
	switch k.(type) {
	case string:
		s = k.(string)
	case uint8:
		return int(k.(uint8))
	case int32:
		return int(k.(int32))
	case int64:
		return int(k.(int64))
	case int:
		return k.(int)
	case uint32:
		return int(k.(uint32))
	case uint64:
		return int(k.(uint64))
	case uint:
		return int(k.(uint))
	case bool:
		if k.(bool) {
			return 1
		} else {
			return 0
		}
	case float32:
		return int(k.(float32))
	case float64:
		return int(k.(float64))
	case interface{}:
		s = fmt.Sprintf("%v", k)
	default:
		text := fmt.Sprintf("Expected one of [string, {u,}int{,32,64} float{32,64}] but got: %T", k)
		err = errors.New(text)
	}
	if err == nil {
		i, err = atoi(s)
	}

	if err != nil {
		fmt.Errorf("%s", err)
	}
	return i
}

func atoi(s string) (int, error) {
	i, e := strconv.Atoi(s)
	return int(i), e
}

type char byte

func (c char) String() string {
	return string(c)
}

// Generator create array from, to (inclusive), using step
// If only one argument then assume from 0, to but excluding (args[0]) step 1
// that is [0..args[0])
// otherwise [from..to) if step, increment by step for each element
// from 1st arg - default 0
// to   2nd arg - default 0
// step 3rd arg - default 1
func Generator(args ...interface{}) (result []int) {
	if len(args) == 0 {
		return
	}
	var from, to, step, v int = 0, 0, 1, 0
	for i := 0; i < len(args); i++ {
		v = ToInt(args[i])
		switch i {
		case 0:
			if len(args) == 1 {
				to = v
			} else {
				from = v
			}
		case 1:
			to = v
		case 2:
			step = v
		}
	}

	for i := 0; i < to; i += step {
		result = append(result, from+i)
	}
	return
}

// GeneratorChar create array from, to (inclusive), using step
// If only one argument then assume from 0, to but excluding (args[0]) step 1
// that is [0..args[0])
// otherwise [from..to] if step, increment by step for each element
// from 1st arg - default 0
// to   2nd arg - default 0
// step 3rd arg - default 1
func GeneratorChar(args ...interface{}) (result []string) {
	if len(args) == 0 {
		return
	}
	var from, to, step, v int = 0, 0, 1, 0
	for i := 0; i < len(args); i++ {
		switch args[i].(type) {
		case string:
			a := byte("a"[0])
			z := byte("z"[0])
			s := []byte(args[i].(string))
			if len(s) == 1 && s[0] >= a && s[0] <= z {
				v = int(s[0] - a)
			} else {
				v = ToInt(args[i])
			}
		default:
			v = ToInt(args[i])
		}
		switch i {
		case 0:
			if len(args) == 1 {
				to = v
			} else {
				from = v
			}
		case 1:
			to = v
		case 2:
			step = v
		}
	}

	for i := 0; i < to; i += step {
		result = append(result, string('a'+from+i))
	}
	return
}
