package flag

////////////////////////////////////////////////////////////////////////
//
////////////////////////////////////////////////////////////////////////
import (
        "fmt"
        "strings"
        "strconv"
        "reflect"
        "time"
)


// sliceDurationValue []time.Duration
type sliceDurationValue []time.Duration

func newsliceDurationValue(val sliceDurationValue, p *sliceDurationValue) *sliceDurationValue {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}


// Set a slice after parsing a string
func (slc *sliceDurationValue) Set(s string) error {
  var T = reflect.TypeOf(sliceDurationValue{}).Elem()
  var debugging = false
  if debugging {
    fmt.Printf("%v %T\n", T,T)
  }
  
	var l = strings.Split(s, ",")

	for _, text := range l {
    if text = strings.TrimSpace(text); len(text)>0 {
      // start time.Duration
      lhs, err := time.ParseDuration(text)
      if err != nil {
        panic(err)
      }
      *slc = append(*slc, lhs)
      // end time.Duration
    }
	}
	return nil
}

// Get get a slice interface from the value
func (slc *sliceDurationValue) Get() interface{} { return ([]time.Duration)(*slc) }

// String join a string from slice
func (slc *sliceDurationValue) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// sliceDurationVar defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) sliceDurationVar(p *sliceDurationValue, name string, value sliceDurationValue, usage string) {
	f.Var(newsliceDurationValue(value, p), name, usage)
}

// sliceDurationVar defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func sliceDurationVar(p *sliceDurationValue, name string, value sliceDurationValue, usage string) {
	CommandLine.Var(newsliceDurationValue(value, p), name, usage)
}

// sliceDuration defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) sliceDuration(name string, value sliceDurationValue, usage string) *sliceDurationValue {
	p := new(sliceDurationValue)
	f.sliceDurationVar(p, name, value, usage)
	return p
}

// sliceDuration defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func sliceDuration(name string, value sliceDurationValue, usage string) *sliceDurationValue {
	return CommandLine.sliceDuration(name, value, usage)
}

// sliceIntValue []int
type sliceIntValue []int

func newsliceIntValue(val sliceIntValue, p *sliceIntValue) *sliceIntValue {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}


// Set a slice after parsing a string
func (slc *sliceIntValue) Set(s string) error {
  var T = reflect.TypeOf(sliceIntValue{}).Elem()
  var debugging = false
  if debugging {
    fmt.Printf("%v %T\n", T,T)
  }
  var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
    if text = strings.TrimSpace(text); len(text)>0 {
      // start int
      n, _ = strconv.ParseInt(text, 0, T.Bits())
      *slc = append(*slc, (int)(n.(int64)))
      // end int
    }
	}
	return nil
}

// Get get a slice interface from the value
func (slc *sliceIntValue) Get() interface{} { return ([]int)(*slc) }

// String join a string from slice
func (slc *sliceIntValue) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// sliceIntVar defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) sliceIntVar(p *sliceIntValue, name string, value sliceIntValue, usage string) {
	f.Var(newsliceIntValue(value, p), name, usage)
}

// sliceIntVar defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func sliceIntVar(p *sliceIntValue, name string, value sliceIntValue, usage string) {
	CommandLine.Var(newsliceIntValue(value, p), name, usage)
}

// sliceInt defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) sliceInt(name string, value sliceIntValue, usage string) *sliceIntValue {
	p := new(sliceIntValue)
	f.sliceIntVar(p, name, value, usage)
	return p
}

// sliceInt defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func sliceInt(name string, value sliceIntValue, usage string) *sliceIntValue {
	return CommandLine.sliceInt(name, value, usage)
}

// sliceInt8Value []int8
type sliceInt8Value []int8

func newsliceInt8Value(val sliceInt8Value, p *sliceInt8Value) *sliceInt8Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}


// Set a slice after parsing a string
func (slc *sliceInt8Value) Set(s string) error {
  var T = reflect.TypeOf(sliceInt8Value{}).Elem()
  var debugging = false
  if debugging {
    fmt.Printf("%v %T\n", T,T)
  }
  var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
    if text = strings.TrimSpace(text); len(text)>0 {
      // start int8
      n, _ = strconv.ParseInt(text, 0, T.Bits())
      *slc = append(*slc, (int8)(n.(int64)))
      // end int8
    }
	}
	return nil
}

// Get get a slice interface from the value
func (slc *sliceInt8Value) Get() interface{} { return ([]int8)(*slc) }

// String join a string from slice
func (slc *sliceInt8Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// sliceInt8Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) sliceInt8Var(p *sliceInt8Value, name string, value sliceInt8Value, usage string) {
	f.Var(newsliceInt8Value(value, p), name, usage)
}

// sliceInt8Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func sliceInt8Var(p *sliceInt8Value, name string, value sliceInt8Value, usage string) {
	CommandLine.Var(newsliceInt8Value(value, p), name, usage)
}

// sliceInt8 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) sliceInt8(name string, value sliceInt8Value, usage string) *sliceInt8Value {
	p := new(sliceInt8Value)
	f.sliceInt8Var(p, name, value, usage)
	return p
}

// sliceInt8 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func sliceInt8(name string, value sliceInt8Value, usage string) *sliceInt8Value {
	return CommandLine.sliceInt8(name, value, usage)
}

// sliceInt16Value []int16
type sliceInt16Value []int16

func newsliceInt16Value(val sliceInt16Value, p *sliceInt16Value) *sliceInt16Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}


// Set a slice after parsing a string
func (slc *sliceInt16Value) Set(s string) error {
  var T = reflect.TypeOf(sliceInt16Value{}).Elem()
  var debugging = false
  if debugging {
    fmt.Printf("%v %T\n", T,T)
  }
  var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
    if text = strings.TrimSpace(text); len(text)>0 {
      // start int16
      n, _ = strconv.ParseInt(text, 0, T.Bits())
      *slc = append(*slc, (int16)(n.(int64)))
      // end int16
    }
	}
	return nil
}

// Get get a slice interface from the value
func (slc *sliceInt16Value) Get() interface{} { return ([]int16)(*slc) }

// String join a string from slice
func (slc *sliceInt16Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// sliceInt16Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) sliceInt16Var(p *sliceInt16Value, name string, value sliceInt16Value, usage string) {
	f.Var(newsliceInt16Value(value, p), name, usage)
}

// sliceInt16Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func sliceInt16Var(p *sliceInt16Value, name string, value sliceInt16Value, usage string) {
	CommandLine.Var(newsliceInt16Value(value, p), name, usage)
}

// sliceInt16 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) sliceInt16(name string, value sliceInt16Value, usage string) *sliceInt16Value {
	p := new(sliceInt16Value)
	f.sliceInt16Var(p, name, value, usage)
	return p
}

// sliceInt16 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func sliceInt16(name string, value sliceInt16Value, usage string) *sliceInt16Value {
	return CommandLine.sliceInt16(name, value, usage)
}

// sliceInt32Value []int32
type sliceInt32Value []int32

func newsliceInt32Value(val sliceInt32Value, p *sliceInt32Value) *sliceInt32Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}


// Set a slice after parsing a string
func (slc *sliceInt32Value) Set(s string) error {
  var T = reflect.TypeOf(sliceInt32Value{}).Elem()
  var debugging = false
  if debugging {
    fmt.Printf("%v %T\n", T,T)
  }
  var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
    if text = strings.TrimSpace(text); len(text)>0 {
      // start int32
      n, _ = strconv.ParseInt(text, 0, T.Bits())
      *slc = append(*slc, (int32)(n.(int64)))
      // end int32
    }
	}
	return nil
}

// Get get a slice interface from the value
func (slc *sliceInt32Value) Get() interface{} { return ([]int32)(*slc) }

// String join a string from slice
func (slc *sliceInt32Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// sliceInt32Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) sliceInt32Var(p *sliceInt32Value, name string, value sliceInt32Value, usage string) {
	f.Var(newsliceInt32Value(value, p), name, usage)
}

// sliceInt32Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func sliceInt32Var(p *sliceInt32Value, name string, value sliceInt32Value, usage string) {
	CommandLine.Var(newsliceInt32Value(value, p), name, usage)
}

// sliceInt32 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) sliceInt32(name string, value sliceInt32Value, usage string) *sliceInt32Value {
	p := new(sliceInt32Value)
	f.sliceInt32Var(p, name, value, usage)
	return p
}

// sliceInt32 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func sliceInt32(name string, value sliceInt32Value, usage string) *sliceInt32Value {
	return CommandLine.sliceInt32(name, value, usage)
}

// sliceInt64Value []int64
type sliceInt64Value []int64

func newsliceInt64Value(val sliceInt64Value, p *sliceInt64Value) *sliceInt64Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}


// Set a slice after parsing a string
func (slc *sliceInt64Value) Set(s string) error {
  var T = reflect.TypeOf(sliceInt64Value{}).Elem()
  var debugging = false
  if debugging {
    fmt.Printf("%v %T\n", T,T)
  }
  var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
    if text = strings.TrimSpace(text); len(text)>0 {
      // start int64
      n, _ = strconv.ParseInt(text, 0, T.Bits())
      *slc = append(*slc, (int64)(n.(int64)))
      // end int64
    }
	}
	return nil
}

// Get get a slice interface from the value
func (slc *sliceInt64Value) Get() interface{} { return ([]int64)(*slc) }

// String join a string from slice
func (slc *sliceInt64Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// sliceInt64Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) sliceInt64Var(p *sliceInt64Value, name string, value sliceInt64Value, usage string) {
	f.Var(newsliceInt64Value(value, p), name, usage)
}

// sliceInt64Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func sliceInt64Var(p *sliceInt64Value, name string, value sliceInt64Value, usage string) {
	CommandLine.Var(newsliceInt64Value(value, p), name, usage)
}

// sliceInt64 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) sliceInt64(name string, value sliceInt64Value, usage string) *sliceInt64Value {
	p := new(sliceInt64Value)
	f.sliceInt64Var(p, name, value, usage)
	return p
}

// sliceInt64 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func sliceInt64(name string, value sliceInt64Value, usage string) *sliceInt64Value {
	return CommandLine.sliceInt64(name, value, usage)
}

// sliceUintValue []uint
type sliceUintValue []uint

func newsliceUintValue(val sliceUintValue, p *sliceUintValue) *sliceUintValue {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}


// Set a slice after parsing a string
func (slc *sliceUintValue) Set(s string) error {
  var T = reflect.TypeOf(sliceUintValue{}).Elem()
  var debugging = false
  if debugging {
    fmt.Printf("%v %T\n", T,T)
  }
  var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
    if text = strings.TrimSpace(text); len(text)>0 {
      // start uint
      n, _ = strconv.ParseUint(text, 0, T.Bits())
      *slc = append(*slc, (uint)(n.(uint64)))
      // end uint
    }
	}
	return nil
}

// Get get a slice interface from the value
func (slc *sliceUintValue) Get() interface{} { return ([]uint)(*slc) }

// String join a string from slice
func (slc *sliceUintValue) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// sliceUintVar defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) sliceUintVar(p *sliceUintValue, name string, value sliceUintValue, usage string) {
	f.Var(newsliceUintValue(value, p), name, usage)
}

// sliceUintVar defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func sliceUintVar(p *sliceUintValue, name string, value sliceUintValue, usage string) {
	CommandLine.Var(newsliceUintValue(value, p), name, usage)
}

// sliceUint defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) sliceUint(name string, value sliceUintValue, usage string) *sliceUintValue {
	p := new(sliceUintValue)
	f.sliceUintVar(p, name, value, usage)
	return p
}

// sliceUint defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func sliceUint(name string, value sliceUintValue, usage string) *sliceUintValue {
	return CommandLine.sliceUint(name, value, usage)
}

// sliceUint8Value []uint8
type sliceUint8Value []uint8

func newsliceUint8Value(val sliceUint8Value, p *sliceUint8Value) *sliceUint8Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}


// Set a slice after parsing a string
func (slc *sliceUint8Value) Set(s string) error {
  var T = reflect.TypeOf(sliceUint8Value{}).Elem()
  var debugging = false
  if debugging {
    fmt.Printf("%v %T\n", T,T)
  }
  var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
    if text = strings.TrimSpace(text); len(text)>0 {
      // start uint8
      n, _ = strconv.ParseUint(text, 0, T.Bits())
      *slc = append(*slc, (uint8)(n.(uint64)))
      // end uint8
    }
	}
	return nil
}

// Get get a slice interface from the value
func (slc *sliceUint8Value) Get() interface{} { return ([]uint8)(*slc) }

// String join a string from slice
func (slc *sliceUint8Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// sliceUint8Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) sliceUint8Var(p *sliceUint8Value, name string, value sliceUint8Value, usage string) {
	f.Var(newsliceUint8Value(value, p), name, usage)
}

// sliceUint8Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func sliceUint8Var(p *sliceUint8Value, name string, value sliceUint8Value, usage string) {
	CommandLine.Var(newsliceUint8Value(value, p), name, usage)
}

// sliceUint8 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) sliceUint8(name string, value sliceUint8Value, usage string) *sliceUint8Value {
	p := new(sliceUint8Value)
	f.sliceUint8Var(p, name, value, usage)
	return p
}

// sliceUint8 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func sliceUint8(name string, value sliceUint8Value, usage string) *sliceUint8Value {
	return CommandLine.sliceUint8(name, value, usage)
}

// sliceUint16Value []uint16
type sliceUint16Value []uint16

func newsliceUint16Value(val sliceUint16Value, p *sliceUint16Value) *sliceUint16Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}


// Set a slice after parsing a string
func (slc *sliceUint16Value) Set(s string) error {
  var T = reflect.TypeOf(sliceUint16Value{}).Elem()
  var debugging = false
  if debugging {
    fmt.Printf("%v %T\n", T,T)
  }
  var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
    if text = strings.TrimSpace(text); len(text)>0 {
      // start uint16
      n, _ = strconv.ParseUint(text, 0, T.Bits())
      *slc = append(*slc, (uint16)(n.(uint64)))
      // end uint16
    }
	}
	return nil
}

// Get get a slice interface from the value
func (slc *sliceUint16Value) Get() interface{} { return ([]uint16)(*slc) }

// String join a string from slice
func (slc *sliceUint16Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// sliceUint16Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) sliceUint16Var(p *sliceUint16Value, name string, value sliceUint16Value, usage string) {
	f.Var(newsliceUint16Value(value, p), name, usage)
}

// sliceUint16Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func sliceUint16Var(p *sliceUint16Value, name string, value sliceUint16Value, usage string) {
	CommandLine.Var(newsliceUint16Value(value, p), name, usage)
}

// sliceUint16 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) sliceUint16(name string, value sliceUint16Value, usage string) *sliceUint16Value {
	p := new(sliceUint16Value)
	f.sliceUint16Var(p, name, value, usage)
	return p
}

// sliceUint16 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func sliceUint16(name string, value sliceUint16Value, usage string) *sliceUint16Value {
	return CommandLine.sliceUint16(name, value, usage)
}

// sliceUint32Value []uint32
type sliceUint32Value []uint32

func newsliceUint32Value(val sliceUint32Value, p *sliceUint32Value) *sliceUint32Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}


// Set a slice after parsing a string
func (slc *sliceUint32Value) Set(s string) error {
  var T = reflect.TypeOf(sliceUint32Value{}).Elem()
  var debugging = false
  if debugging {
    fmt.Printf("%v %T\n", T,T)
  }
  var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
    if text = strings.TrimSpace(text); len(text)>0 {
      // start uint32
      n, _ = strconv.ParseUint(text, 0, T.Bits())
      *slc = append(*slc, (uint32)(n.(uint64)))
      // end uint32
    }
	}
	return nil
}

// Get get a slice interface from the value
func (slc *sliceUint32Value) Get() interface{} { return ([]uint32)(*slc) }

// String join a string from slice
func (slc *sliceUint32Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// sliceUint32Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) sliceUint32Var(p *sliceUint32Value, name string, value sliceUint32Value, usage string) {
	f.Var(newsliceUint32Value(value, p), name, usage)
}

// sliceUint32Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func sliceUint32Var(p *sliceUint32Value, name string, value sliceUint32Value, usage string) {
	CommandLine.Var(newsliceUint32Value(value, p), name, usage)
}

// sliceUint32 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) sliceUint32(name string, value sliceUint32Value, usage string) *sliceUint32Value {
	p := new(sliceUint32Value)
	f.sliceUint32Var(p, name, value, usage)
	return p
}

// sliceUint32 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func sliceUint32(name string, value sliceUint32Value, usage string) *sliceUint32Value {
	return CommandLine.sliceUint32(name, value, usage)
}

// sliceUint64Value []uint64
type sliceUint64Value []uint64

func newsliceUint64Value(val sliceUint64Value, p *sliceUint64Value) *sliceUint64Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}


// Set a slice after parsing a string
func (slc *sliceUint64Value) Set(s string) error {
  var T = reflect.TypeOf(sliceUint64Value{}).Elem()
  var debugging = false
  if debugging {
    fmt.Printf("%v %T\n", T,T)
  }
  var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
    if text = strings.TrimSpace(text); len(text)>0 {
      // start uint64
      n, _ = strconv.ParseUint(text, 0, T.Bits())
      *slc = append(*slc, (uint64)(n.(uint64)))
      // end uint64
    }
	}
	return nil
}

// Get get a slice interface from the value
func (slc *sliceUint64Value) Get() interface{} { return ([]uint64)(*slc) }

// String join a string from slice
func (slc *sliceUint64Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// sliceUint64Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) sliceUint64Var(p *sliceUint64Value, name string, value sliceUint64Value, usage string) {
	f.Var(newsliceUint64Value(value, p), name, usage)
}

// sliceUint64Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func sliceUint64Var(p *sliceUint64Value, name string, value sliceUint64Value, usage string) {
	CommandLine.Var(newsliceUint64Value(value, p), name, usage)
}

// sliceUint64 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) sliceUint64(name string, value sliceUint64Value, usage string) *sliceUint64Value {
	p := new(sliceUint64Value)
	f.sliceUint64Var(p, name, value, usage)
	return p
}

// sliceUint64 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func sliceUint64(name string, value sliceUint64Value, usage string) *sliceUint64Value {
	return CommandLine.sliceUint64(name, value, usage)
}

// sliceFloat64Value []float64
type sliceFloat64Value []float64

func newsliceFloat64Value(val sliceFloat64Value, p *sliceFloat64Value) *sliceFloat64Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}


// Set a slice after parsing a string
func (slc *sliceFloat64Value) Set(s string) error {
  var T = reflect.TypeOf(sliceFloat64Value{}).Elem()
  var debugging = false
  if debugging {
    fmt.Printf("%v %T\n", T,T)
  }
  var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
    if text = strings.TrimSpace(text); len(text)>0 {
      // start float64
      n, _ = strconv.ParseFloat(text, T.Bits())
      *slc = append(*slc, (float64)(n.(float64)))
      // end float64
    }
	}
	return nil
}

// Get get a slice interface from the value
func (slc *sliceFloat64Value) Get() interface{} { return ([]float64)(*slc) }

// String join a string from slice
func (slc *sliceFloat64Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// sliceFloat64Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) sliceFloat64Var(p *sliceFloat64Value, name string, value sliceFloat64Value, usage string) {
	f.Var(newsliceFloat64Value(value, p), name, usage)
}

// sliceFloat64Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func sliceFloat64Var(p *sliceFloat64Value, name string, value sliceFloat64Value, usage string) {
	CommandLine.Var(newsliceFloat64Value(value, p), name, usage)
}

// sliceFloat64 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) sliceFloat64(name string, value sliceFloat64Value, usage string) *sliceFloat64Value {
	p := new(sliceFloat64Value)
	f.sliceFloat64Var(p, name, value, usage)
	return p
}

// sliceFloat64 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func sliceFloat64(name string, value sliceFloat64Value, usage string) *sliceFloat64Value {
	return CommandLine.sliceFloat64(name, value, usage)
}

// sliceFloat32Value []float32
type sliceFloat32Value []float32

func newsliceFloat32Value(val sliceFloat32Value, p *sliceFloat32Value) *sliceFloat32Value {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}


// Set a slice after parsing a string
func (slc *sliceFloat32Value) Set(s string) error {
  var T = reflect.TypeOf(sliceFloat32Value{}).Elem()
  var debugging = false
  if debugging {
    fmt.Printf("%v %T\n", T,T)
  }
  var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
    if text = strings.TrimSpace(text); len(text)>0 {
      // start float32
      n, _ = strconv.ParseFloat(text, T.Bits())
      *slc = append(*slc, (float32)(n.(float64)))
      // end float32
    }
	}
	return nil
}

// Get get a slice interface from the value
func (slc *sliceFloat32Value) Get() interface{} { return ([]float32)(*slc) }

// String join a string from slice
func (slc *sliceFloat32Value) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// sliceFloat32Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) sliceFloat32Var(p *sliceFloat32Value, name string, value sliceFloat32Value, usage string) {
	f.Var(newsliceFloat32Value(value, p), name, usage)
}

// sliceFloat32Var defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func sliceFloat32Var(p *sliceFloat32Value, name string, value sliceFloat32Value, usage string) {
	CommandLine.Var(newsliceFloat32Value(value, p), name, usage)
}

// sliceFloat32 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) sliceFloat32(name string, value sliceFloat32Value, usage string) *sliceFloat32Value {
	p := new(sliceFloat32Value)
	f.sliceFloat32Var(p, name, value, usage)
	return p
}

// sliceFloat32 defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func sliceFloat32(name string, value sliceFloat32Value, usage string) *sliceFloat32Value {
	return CommandLine.sliceFloat32(name, value, usage)
}

// sliceBoolValue []bool
type sliceBoolValue []bool

func newsliceBoolValue(val sliceBoolValue, p *sliceBoolValue) *sliceBoolValue {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}


// Set a slice after parsing a string
func (slc *sliceBoolValue) Set(s string) error {
  var T = reflect.TypeOf(sliceBoolValue{}).Elem()
  var debugging = false
  if debugging {
    fmt.Printf("%v %T\n", T,T)
  }
  var n interface{}
	var l = strings.Split(s, ",")

	for _, text := range l {
    if text = strings.TrimSpace(text); len(text)>0 {
      // start bool
      n, _ = strconv.ParseBool(text)
      *slc = append(*slc, (bool)(n.(bool)))
      // end bool
    }
	}
	return nil
}

// Get get a slice interface from the value
func (slc *sliceBoolValue) Get() interface{} { return ([]bool)(*slc) }

// String join a string from slice
func (slc *sliceBoolValue) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// sliceBoolVar defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) sliceBoolVar(p *sliceBoolValue, name string, value sliceBoolValue, usage string) {
	f.Var(newsliceBoolValue(value, p), name, usage)
}

// sliceBoolVar defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func sliceBoolVar(p *sliceBoolValue, name string, value sliceBoolValue, usage string) {
	CommandLine.Var(newsliceBoolValue(value, p), name, usage)
}

// sliceBool defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) sliceBool(name string, value sliceBoolValue, usage string) *sliceBoolValue {
	p := new(sliceBoolValue)
	f.sliceBoolVar(p, name, value, usage)
	return p
}

// sliceBool defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func sliceBool(name string, value sliceBoolValue, usage string) *sliceBoolValue {
	return CommandLine.sliceBool(name, value, usage)
}

// sliceStringValue []string
type sliceStringValue []string

func newsliceStringValue(val sliceStringValue, p *sliceStringValue) *sliceStringValue {
	for i := 0; i < len(val); i++ {
		*p = append(*p, val[i])
	}
	return p
}


// Set a slice after parsing a string
func (slc *sliceStringValue) Set(s string) error {
  var T = reflect.TypeOf(sliceStringValue{}).Elem()
  var debugging = false
  if debugging {
    fmt.Printf("%v %T\n", T,T)
  }
  
	var l = strings.Split(s, ",")

	for _, text := range l {
    if text = strings.TrimSpace(text); len(text)>0 {
      // start string
      *slc = append(*slc, text)
      // end string
    }
	}
	return nil
}

// Get get a slice interface from the value
func (slc *sliceStringValue) Get() interface{} { return ([]string)(*slc) }

// String join a string from slice
func (slc *sliceStringValue) String() string {
  t := []string{}
  for _, v := range *slc {
    t = append(t, fmt.Sprintf("%v", v))
  }
  return strings.Join(t, ",")
}

// sliceStringVar defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func (f *FlagSet) sliceStringVar(p *sliceStringValue, name string, value sliceStringValue, usage string) {
	f.Var(newsliceStringValue(value, p), name, usage)
}

// sliceStringVar defines an slice flag with specified name,
// default value, and usage string.  The argument p points to an slice
// variable in which to store the value of the flag.
func sliceStringVar(p *sliceStringValue, name string, value sliceStringValue, usage string) {
	CommandLine.Var(newsliceStringValue(value, p), name, usage)
}

// sliceString defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func (f *FlagSet) sliceString(name string, value sliceStringValue, usage string) *sliceStringValue {
	p := new(sliceStringValue)
	f.sliceStringVar(p, name, value, usage)
	return p
}

// sliceString defines an slice flag with specified name, default value, and
// usage string.  The return value is the address of an slice variable
// that stores the value of the flag.
func sliceString(name string, value sliceStringValue, usage string) *sliceStringValue {
	return CommandLine.sliceString(name, value, usage)
}

