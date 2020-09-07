package cfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

// DeepCopier recurse and iterate over configuration structs internal
// representation Stor
type DeepCopier interface {
	DeepCopyInto(DeepCopier)
}

// Stor configuration representation, restorable object from, saveable
// to persistence.
type Stor map[string]interface{}

// DeepCopyInto ...
func (stor Stor) DeepCopyInto(dst interface{}, depthArgs ...int) {
	const max = 3
	var depth = 0
	if len(depthArgs) > 0 {
		depth = depthArgs[0]
	}
	fmt.Println("depth", depth)
	if depth > max {
		panic(fmt.Errorf("exceeded struct max depth of %d", max))
	}
	for k, v := range stor {
		switch v.(type) {
		case DeepCopier:
			src := v.(DeepCopier)
			if src != nil {
				src.DeepCopyInto(Store[k].(DeepCopier))
			}
		case Stor:
			src := v.(Stor)
			if src != nil {
				src.DeepCopyInto(Store[k], depth)
			}
		}
	}
}

// DeepCopyInto dst from source
func DeepCopyInto(src Stor, dst map[string]interface{}, depthArgs ...int) {
	const max = 3
	var depth = 0
	if len(depthArgs) > 0 {
		depth = depthArgs[0]
	}
	fmt.Println("depth", depth)
	if depth > max {
		panic(fmt.Errorf("exceeded struct max depth of %d", max))
	}
	for k, v := range src {
		switch v.(type) {
		case DeepCopier:
			src := v.(DeepCopier)
			if src != nil {
				src.DeepCopyInto(Store[k].(DeepCopier))
			}
		case Stor:
			src := v.(Stor)
			if src != nil {
				src.DeepCopyInto(Store[k].(Stor), depth)
			}
		}
	}
}

// Loader reads input configuration from a store
type Loader interface {
	Load(filename string)
}

// Storer writes configurations to store
type Storer interface {
	Store(filename string)
}

// Adder appends to configuration objects
type Adder interface {
	AddStor(name string, o interface{})
}

// NewStor returns a Stor object for persistence
func NewStor() Stor {
	return Stor{}
}

// // Read object from io.Reader
// func (stor Stor) Read(r io.Reader) (int, error) {
// 	var data = []byte{}
// 	n, err := r.Read(data)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return n, err
// }

// // Write object from io.Writer
// func (stor Stor) Write(w io.Writer) (int, error) {
// 	data, err := yaml.Marshal(stor)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return w.Write(data)
// }

// AddStor appends an object to a Stor persistence
func (stor Stor) AddStor(name string, o interface{}) {
	stor[name] = o
}

// Load object from persistence
func (stor *Stor) Load(filename string) error {
	var err error
	var data []byte

	// fmt.Println("filename", filename)
	data, err = ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	// fmt.Println(string(data))
	// Save pointers to underlying
	var temp = Stor{}
	for k, v := range Store {
		temp[k] = v
	}
	switch strings.ToLower(path.Ext(filename)) {
	case ".json":
		err = json.Unmarshal(data, temp)
		if err != nil {
			return err
		}
	case ".yaml":
		err = yaml.Unmarshal(data, temp)
		if err != nil {
			return err
		}
	default:
		err = yaml.Unmarshal(data, temp)
		if err != nil {
			return err
		}
	}

	for k, v := range temp {
		CopyOut(v, Store[k])
	}

	return err
}

// CopyOut serializes and deseiralizes an object to copy it out of an
// object
func CopyOut(in, out interface{}) error {
	text, err := yaml.Marshal(in)
	if err != nil {
		log.Println(err)
		return err
	}

	err = yaml.Unmarshal(text, out)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func Dump(o interface{}) string {
	byte, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		log.Println(err)
	}
	return string(byte)
}

///////////////////// func (src Stor) Copy(dst interface{}) {
///////////////////// 	fmt.Printf("Stor type mismatch %T %T", src, dst)
///////////////////// 	src.DeepCopyInto(dst)

///////////////////// 	/////////// for k, v := range src {
///////////////////// 	/////////// 	if IsStructPtr(v) {
///////////////////// 	/////////// 		origin, ok := v.(DeepCopier)
///////////////////// 	/////////// 		if !ok {
///////////////////// 	/////////// 			panic(fmt.Errorf("Stor type mismatch %+T %+T", src, dst))
///////////////////// 	/////////// 		}
///////////////////// 	/////////// 		origin.DeepCopyInto(dst.(DeepCopier))
///////////////////// 	/////////// 	} else {
///////////////////// 	/////////// 		// panic(fmt.Errorf("Stor types %+T %+T", v, dst))
///////////////////// 	/////////// 		// panic(ErrInvalidArgStructPointerRequired)
///////////////////// 	/////////// 	}
///////////////////// 	/////////// }
///////////////////// }

/////////// func DeepCopy(src Stor, dst map[string]interface{}){
/////////// // DeepCopy from the internal store to the config object's map representation
/////////// func DeepCopy(src Stor, dst interface{}) {
/////////// 	if !IsStructPtr(dst) {
/////////// 		panic(ErrInvalidArgStructPointerRequired)
/////////// 	}
/////////// 	var err error
/////////// 	var TypeOf = reflect.TypeOf(dst).String()
/////////// 	lookup := RemovePkg(TypeOf)
/////////// 	if target, found := src[lookup]; found {
/////////// 		switch target.(type) {
/////////// 		case Stor: // NoOp
/////////// 		case map[string]interface{}:
/////////// 			to := Stor{}
/////////// 			for k, v := range target.(map[string]interface{}) {
/////////// 				to[k] = v
/////////// 			}
/////////// 			target = to
/////////// 		}
/////////// 		err = json.Unmarshal(target.(Stor).Bytes(), dst)
/////////// 		if err != nil {
/////////// 			fmt.Fprintf(os.Stderr, "config to string: %s", err.Error())
/////////// 			os.Exit(1)
/////////// 		}
/////////// 	}
///////////       }

// Stor object to persistence
func (stor Stor) Stor(filename string) error {
	return stor.Save(filename)
}

// Save object to persistence
func (stor Stor) Save(filename string) error {
	var err error
	var data []byte
	switch strings.ToLower(path.Ext(filename)) {
	case ".json":
		data, err = json.Marshal(stor)
		if err != nil {
			return err
		}
	case ".yaml":
		data, err = yaml.Marshal(stor)
		if err != nil {
			return err
		}
	default:
		data, err = yaml.Marshal(stor)
		if err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// String stor interfaceable
func (stor Stor) Bytes() []byte {
	text, err := json.Marshal(stor)
	if err != nil {
		fmt.Fprintf(os.Stderr, "stor to bytes: %s", err.Error())
		return nil
	}
	return text
}

// String stor interfaceable
func (stor Stor) String() string {
	return string(stor.Bytes())
}
