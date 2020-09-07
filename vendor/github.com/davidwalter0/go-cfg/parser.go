package cfg

import (
	"encoding/json"
	"log"
	"reflect"
	"strings"
	//	"strconv"
)

const (
	cfgEnvKeyPrefix = "CFG_KEY_PREFIX"
	cfgDecorate     = "CFG_DECORATE"
)

/*
// Arg to parse
type Arg []string

// Key struct name
type Key string

// KV map of key(struct element) + tags
type KV map[Key]*Field

// Parser recursively parses a struct
type Parser struct {
	ptr interface{}
	kv  KV
	mgr Mgr
}

// KV returns a map of object definitions
func (parser *Parser) KV() KV {
	return parser.kv
}

// NewParser return an initialized parser using args
func NewParser(ptr interface{}) (*Parser, error) {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		return nil, ErrInvalidArgPointerRequired
	}
	return &Parser{
		ptr: ptr,
		kv:  make(KV),
		mgr: NewCache(),
	}, nil
}
*/
var emptyStructField = reflect.StructField{}

// Store persistable representation
var Store = NewStor()

type SFWrap struct {
	reflect.StructField
}

// GetName from reflect struct field tag or struct name
func (sfwrap *SFWrap) GetName() string {
	name := sfwrap.Tag.Get("json")
	i := strings.Index(name, ",")
	if i > 0 {
		name = name[0:i]
	}
	if len(name) > 0 {
		return name
	}
	return sfwrap.Name
}

// Enter recursively processes object configurations
func Enter(args *Arg, ptr interface{}) error {
	var err error
	elem := reflect.ValueOf(ptr).Elem()
	etype := elem.Type()
	name := etype.Name()

	Store[name] = ptr

	if args.Prefixed && len(args.Prefix) > 0 {
		name = args.Prefix + "_" + name
	}

	if !decorate {
		name = ""
	}

	err = ParseStruct(args, ptr, name, emptyStructField)
	if err != nil {
		log.Printf("error parsing %s %+v\n", etype.Name(), err)
	}

	return err
}

// ParseStruct recursively processes object configurations
func ParseStruct(args *Arg, ptr interface{}, prefix string, structField reflect.StructField) error {
	depth := args.Depth
	var err error
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		panic(ErrInvalidArgPointerRequired)
	}
	kind := reflect.ValueOf(ptr).Elem().Kind()
	switch kind {
	case reflect.Struct:
		elem := reflect.ValueOf(ptr).Elem()
		etype := elem.Type()
		indirect := reflect.Indirect(reflect.ValueOf(ptr))
		for i := 0; i < etype.NumField(); i++ {
			var name string
			ptr := elem.Field(i).Addr().Interface()
			sfwrap := &SFWrap{indirect.Type().Field(i)}
			if args.Prefixed && len(prefix) != 0 {
				name = Capitalize(prefix) + "_" + sfwrap.GetName()
				// name = Capitalize(prefix) + "_" + Capitalize(indirect.Type().Field(i).Name)
			} else {
				name = Capitalize(sfwrap.GetName())
				// name = Capitalize(indirect.Type().Field(i).Name)
			}
			args.Depth++
			err = ParseStruct(args, ptr, name, etype.Field(i))
			if err != nil {
				panic(err)
				break
			}
			args.Depth--
		}
	default:
		elem := reflect.ValueOf(ptr).Elem()
		etype := elem.Type()
		sfwrap := &SFWrap{structField}
		name := sfwrap.GetName()
		var field = &Field{
			StructField: structField,
			FieldPtr:    ptr,
			UseFlags:    args.UseFlags,
			Depth:       depth,
			Name:        name,
			Prefix:      prefix,
			KeyName:     prefix,
			FlagName:    prefix,
			Type:        etype.Name(),
		}

		field.SetField()
		if field.Error != nil {
			log.Println(field.Error)
			panic(field.Error)
			return nil
		}

		if debug {
			byte, err := json.Marshal(field)
			if err != nil {
				log.Println(err)
				panic(err)
			}
			log.Println(">> ", string(byte))
		}

		return err
	}
	return err
}
