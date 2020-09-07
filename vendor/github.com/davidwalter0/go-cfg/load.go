package cfg

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
)

// RemovePkg from a type name
func RemovePkg(pkg string) string {
	regx := regexp.MustCompile(`\*?[^.]*\.`)
	text := regx.ReplaceAll([]byte(pkg), []byte{})
	return string(text)
}

var ErrInvalidArgStructPointerRequired = errors.New("Argument must be a struct pointer")

func IsStructPtr(config interface{}) bool {
	return reflect.TypeOf(config).Kind() == reflect.Ptr &&
		reflect.ValueOf(config).Elem().Kind() == reflect.Struct
}

func IsMap(config interface{}) bool {
	return reflect.TypeOf(config).Kind() == reflect.Map
}

func Type(config interface{}) string {
	if reflect.TypeOf(config).Kind() == reflect.Ptr {
		return fmt.Sprintf("(%s) %s", reflect.TypeOf(config).Kind(), reflect.ValueOf(config).Elem().Kind())
	}
	return fmt.Sprintf("%s", reflect.ValueOf(config).Elem().Kind())
}

// Copy from the internal store to the config object interface{}
func Copy(config interface{}) {
	if !IsStructPtr(config) {
		panic(ErrInvalidArgStructPointerRequired)
	}
	var err error
	var TypeOf = reflect.TypeOf(config).String()
	lookup := RemovePkg(TypeOf)
	if target, found := Store[lookup]; found {
		switch target.(type) {
		case Stor: // NoOp
		case map[string]interface{}:
			to := Stor{}
			for k, v := range target.(map[string]interface{}) {
				to[k] = v
			}
			target = to
		}
		err = json.Unmarshal(target.(Stor).Bytes(), config)
		if err != nil {
			fmt.Fprintf(os.Stderr, "config to string: %s", err.Error())
			os.Exit(1)
		}
	}
}
