package flag // import "github.com/davidwalter0/go-flag"

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ErrInvalidArgPointerRequired pointer error message
var ErrInvalidArgPointerRequired = fmt.Errorf("Argument must be a pointer")

// MapFromText interface text to value
func MapFromText(arg /*map type*/ interface{}, textArg ...string) (err error) {
	var text = ""
	if reflect.TypeOf(arg).Kind() != reflect.Ptr {
		err = ErrInvalidArgPointerRequired
	} else {
		if len(textArg) > 0 {
			text = textArg[0]

			m := reflect.ValueOf(arg).Elem()
			if len(text) != 0 {
				switch m.Kind() {
				case reflect.String:
					m.SetString(text)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					var lhs int64
					if m.Kind() == reflect.Int64 &&
						m.Type().PkgPath() == "time" &&
						m.Type().Name() == "Duration" {
						var t time.Duration
						t, err = time.ParseDuration(text)
						if err != nil {
							return
						}
						lhs = (int64)(t)
					} else {
						lhs, err = strconv.ParseInt(text, 0, m.Type().Bits())
						if err != nil {
							return
						}
					}
					m.SetInt(lhs)
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					var lhs uint64
					lhs, err = strconv.ParseUint(text, 0, m.Type().Bits())
					if err != nil {
						return
					}
					m.SetUint(lhs)
				case reflect.Bool:
					var lhs bool
					lhs, err = strconv.ParseBool(text)
					if err != nil {
						return
					}
					m.SetBool(lhs)
				case reflect.Float32, reflect.Float64:
					var lhs float64
					lhs, err = strconv.ParseFloat(text, m.Type().Bits())
					if err != nil {
						return
					}
					m.SetFloat(lhs)
				case reflect.Slice:
					values := strings.Split(text, ",")
					if len(values) > 0 {
						T := reflect.TypeOf(arg).Elem().Elem()
						// T := reflect.MakeSlice(reflect.TypeOf(arg).Elem())
						// var slice = reflect.ValueOf(arg).Elem()
						var slice = reflect.MakeSlice(reflect.TypeOf(arg).Elem(), 0, 0)
						for _, v := range values {
							if v = strings.TrimSpace(v); len(v) > 0 {
								svalue := reflect.New(T)
								value := svalue.Elem().Addr().Interface()
								if err := MapFromText(value, v); err != nil {
									return err
								}
								slice.Set(reflect.Append(slice, svalue.Elem()))
							}
						}
						reflect.ValueOf(arg).Elem().Set(slice)
					}
				case reflect.Map:
					pairs := strings.Split(text, ",")
					if len(pairs) > 0 {
						mp := reflect.MakeMap(reflect.TypeOf(arg).Elem())
						for _, pair := range pairs {
							kvpair := strings.Split(pair, ":")
							if len(kvpair) != 2 || (len(kvpair[0]) == 0 && len(kvpair[1]) == 0) {
								return fmt.Errorf("Map argument requires non empty pairs key:value text source: %s", pair)
							}
							v := reflect.TypeOf(arg).Elem().Elem()
							k := reflect.TypeOf(arg).Elem().Key()

							key := reflect.New(k).Elem()
							{ // scoped assignment
								lhs := kvpair[0]
								MapFromText(key.Addr().Interface(), lhs)
							}
							value := reflect.New(v).Elem()
							{ // scoped assignment
								lhs := kvpair[1]
								MapFromText(value.Addr().Interface(), lhs)
							}
							// mp.SetMapIndex(key, value)
							mp.SetMapIndex(key, value)
						}
						reflect.ValueOf(arg).Elem().Set(mp)
					}
				}
			}
		}
	}
	return
}
