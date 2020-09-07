package flag

////////////////////////////////////////////////////////////////////////
//
////////////////////////////////////////////////////////////////////////
import (
	"fmt"
	"strings"
	"time"
)

// mapDurationDurationValue []mapDurationDurationValue
type mapDurationDurationValue map[time.Duration]time.Duration

func newmapDurationDurationValue(val mapDurationDurationValue,
	p *mapDurationDurationValue) *mapDurationDurationValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapDurationDurationValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapDurationDurationValue) Get() interface{} { return map[time.Duration]time.Duration(*slc) }

// String join a string from map
func (slc *mapDurationDurationValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapDurationDurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapDurationDurationVar(p *mapDurationDurationValue, name string, value mapDurationDurationValue, usage string, required, isset bool) {
	f.Var(newmapDurationDurationValue(value, p), name, usage, required, isset)
}

// mapDurationDurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapDurationDurationVar(p *mapDurationDurationValue, name string, value mapDurationDurationValue, usage string, required, isset bool) {
	CommandLine.Var(newmapDurationDurationValue(value, p), name, usage, required, isset)
}

// mapDurationDurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapDurationDuration(name string, value mapDurationDurationValue, usage string, required, isset bool) *mapDurationDurationValue {
	p := new(mapDurationDurationValue)
	f.mapDurationDurationVar(p, name, value, usage, required, isset)
	return p
}

// mapDurationDurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapDurationDuration(name string, value mapDurationDurationValue, usage string, required, isset bool) *mapDurationDurationValue {
	return CommandLine.mapDurationDuration(name, value, usage, required, isset)
}

// mapDurationIntValue []mapDurationIntValue
type mapDurationIntValue map[time.Duration]int

func newmapDurationIntValue(val mapDurationIntValue,
	p *mapDurationIntValue) *mapDurationIntValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapDurationIntValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapDurationIntValue) Get() interface{} { return map[time.Duration]int(*slc) }

// String join a string from map
func (slc *mapDurationIntValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapDurationIntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapDurationIntVar(p *mapDurationIntValue, name string, value mapDurationIntValue, usage string, required, isset bool) {
	f.Var(newmapDurationIntValue(value, p), name, usage, required, isset)
}

// mapDurationIntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapDurationIntVar(p *mapDurationIntValue, name string, value mapDurationIntValue, usage string, required, isset bool) {
	CommandLine.Var(newmapDurationIntValue(value, p), name, usage, required, isset)
}

// mapDurationIntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapDurationInt(name string, value mapDurationIntValue, usage string, required, isset bool) *mapDurationIntValue {
	p := new(mapDurationIntValue)
	f.mapDurationIntVar(p, name, value, usage, required, isset)
	return p
}

// mapDurationIntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapDurationInt(name string, value mapDurationIntValue, usage string, required, isset bool) *mapDurationIntValue {
	return CommandLine.mapDurationInt(name, value, usage, required, isset)
}

// mapDurationInt8Value []mapDurationInt8Value
type mapDurationInt8Value map[time.Duration]int8

func newmapDurationInt8Value(val mapDurationInt8Value,
	p *mapDurationInt8Value) *mapDurationInt8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapDurationInt8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapDurationInt8Value) Get() interface{} { return map[time.Duration]int8(*slc) }

// String join a string from map
func (slc *mapDurationInt8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapDurationInt8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapDurationInt8Var(p *mapDurationInt8Value, name string, value mapDurationInt8Value, usage string, required, isset bool) {
	f.Var(newmapDurationInt8Value(value, p), name, usage, required, isset)
}

// mapDurationInt8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapDurationInt8Var(p *mapDurationInt8Value, name string, value mapDurationInt8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapDurationInt8Value(value, p), name, usage, required, isset)
}

// mapDurationInt8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapDurationInt8(name string, value mapDurationInt8Value, usage string, required, isset bool) *mapDurationInt8Value {
	p := new(mapDurationInt8Value)
	f.mapDurationInt8Var(p, name, value, usage, required, isset)
	return p
}

// mapDurationInt8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapDurationInt8(name string, value mapDurationInt8Value, usage string, required, isset bool) *mapDurationInt8Value {
	return CommandLine.mapDurationInt8(name, value, usage, required, isset)
}

// mapDurationInt16Value []mapDurationInt16Value
type mapDurationInt16Value map[time.Duration]int16

func newmapDurationInt16Value(val mapDurationInt16Value,
	p *mapDurationInt16Value) *mapDurationInt16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapDurationInt16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapDurationInt16Value) Get() interface{} { return map[time.Duration]int16(*slc) }

// String join a string from map
func (slc *mapDurationInt16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapDurationInt16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapDurationInt16Var(p *mapDurationInt16Value, name string, value mapDurationInt16Value, usage string, required, isset bool) {
	f.Var(newmapDurationInt16Value(value, p), name, usage, required, isset)
}

// mapDurationInt16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapDurationInt16Var(p *mapDurationInt16Value, name string, value mapDurationInt16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapDurationInt16Value(value, p), name, usage, required, isset)
}

// mapDurationInt16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapDurationInt16(name string, value mapDurationInt16Value, usage string, required, isset bool) *mapDurationInt16Value {
	p := new(mapDurationInt16Value)
	f.mapDurationInt16Var(p, name, value, usage, required, isset)
	return p
}

// mapDurationInt16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapDurationInt16(name string, value mapDurationInt16Value, usage string, required, isset bool) *mapDurationInt16Value {
	return CommandLine.mapDurationInt16(name, value, usage, required, isset)
}

// mapDurationInt32Value []mapDurationInt32Value
type mapDurationInt32Value map[time.Duration]int32

func newmapDurationInt32Value(val mapDurationInt32Value,
	p *mapDurationInt32Value) *mapDurationInt32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapDurationInt32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapDurationInt32Value) Get() interface{} { return map[time.Duration]int32(*slc) }

// String join a string from map
func (slc *mapDurationInt32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapDurationInt32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapDurationInt32Var(p *mapDurationInt32Value, name string, value mapDurationInt32Value, usage string, required, isset bool) {
	f.Var(newmapDurationInt32Value(value, p), name, usage, required, isset)
}

// mapDurationInt32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapDurationInt32Var(p *mapDurationInt32Value, name string, value mapDurationInt32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapDurationInt32Value(value, p), name, usage, required, isset)
}

// mapDurationInt32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapDurationInt32(name string, value mapDurationInt32Value, usage string, required, isset bool) *mapDurationInt32Value {
	p := new(mapDurationInt32Value)
	f.mapDurationInt32Var(p, name, value, usage, required, isset)
	return p
}

// mapDurationInt32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapDurationInt32(name string, value mapDurationInt32Value, usage string, required, isset bool) *mapDurationInt32Value {
	return CommandLine.mapDurationInt32(name, value, usage, required, isset)
}

// mapDurationInt64Value []mapDurationInt64Value
type mapDurationInt64Value map[time.Duration]int64

func newmapDurationInt64Value(val mapDurationInt64Value,
	p *mapDurationInt64Value) *mapDurationInt64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapDurationInt64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapDurationInt64Value) Get() interface{} { return map[time.Duration]int64(*slc) }

// String join a string from map
func (slc *mapDurationInt64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapDurationInt64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapDurationInt64Var(p *mapDurationInt64Value, name string, value mapDurationInt64Value, usage string, required, isset bool) {
	f.Var(newmapDurationInt64Value(value, p), name, usage, required, isset)
}

// mapDurationInt64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapDurationInt64Var(p *mapDurationInt64Value, name string, value mapDurationInt64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapDurationInt64Value(value, p), name, usage, required, isset)
}

// mapDurationInt64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapDurationInt64(name string, value mapDurationInt64Value, usage string, required, isset bool) *mapDurationInt64Value {
	p := new(mapDurationInt64Value)
	f.mapDurationInt64Var(p, name, value, usage, required, isset)
	return p
}

// mapDurationInt64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapDurationInt64(name string, value mapDurationInt64Value, usage string, required, isset bool) *mapDurationInt64Value {
	return CommandLine.mapDurationInt64(name, value, usage, required, isset)
}

// mapDurationUintValue []mapDurationUintValue
type mapDurationUintValue map[time.Duration]uint

func newmapDurationUintValue(val mapDurationUintValue,
	p *mapDurationUintValue) *mapDurationUintValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapDurationUintValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapDurationUintValue) Get() interface{} { return map[time.Duration]uint(*slc) }

// String join a string from map
func (slc *mapDurationUintValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapDurationUintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapDurationUintVar(p *mapDurationUintValue, name string, value mapDurationUintValue, usage string, required, isset bool) {
	f.Var(newmapDurationUintValue(value, p), name, usage, required, isset)
}

// mapDurationUintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapDurationUintVar(p *mapDurationUintValue, name string, value mapDurationUintValue, usage string, required, isset bool) {
	CommandLine.Var(newmapDurationUintValue(value, p), name, usage, required, isset)
}

// mapDurationUintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapDurationUint(name string, value mapDurationUintValue, usage string, required, isset bool) *mapDurationUintValue {
	p := new(mapDurationUintValue)
	f.mapDurationUintVar(p, name, value, usage, required, isset)
	return p
}

// mapDurationUintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapDurationUint(name string, value mapDurationUintValue, usage string, required, isset bool) *mapDurationUintValue {
	return CommandLine.mapDurationUint(name, value, usage, required, isset)
}

// mapDurationUint8Value []mapDurationUint8Value
type mapDurationUint8Value map[time.Duration]uint8

func newmapDurationUint8Value(val mapDurationUint8Value,
	p *mapDurationUint8Value) *mapDurationUint8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapDurationUint8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapDurationUint8Value) Get() interface{} { return map[time.Duration]uint8(*slc) }

// String join a string from map
func (slc *mapDurationUint8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapDurationUint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapDurationUint8Var(p *mapDurationUint8Value, name string, value mapDurationUint8Value, usage string, required, isset bool) {
	f.Var(newmapDurationUint8Value(value, p), name, usage, required, isset)
}

// mapDurationUint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapDurationUint8Var(p *mapDurationUint8Value, name string, value mapDurationUint8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapDurationUint8Value(value, p), name, usage, required, isset)
}

// mapDurationUint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapDurationUint8(name string, value mapDurationUint8Value, usage string, required, isset bool) *mapDurationUint8Value {
	p := new(mapDurationUint8Value)
	f.mapDurationUint8Var(p, name, value, usage, required, isset)
	return p
}

// mapDurationUint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapDurationUint8(name string, value mapDurationUint8Value, usage string, required, isset bool) *mapDurationUint8Value {
	return CommandLine.mapDurationUint8(name, value, usage, required, isset)
}

// mapDurationUint16Value []mapDurationUint16Value
type mapDurationUint16Value map[time.Duration]uint16

func newmapDurationUint16Value(val mapDurationUint16Value,
	p *mapDurationUint16Value) *mapDurationUint16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapDurationUint16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapDurationUint16Value) Get() interface{} { return map[time.Duration]uint16(*slc) }

// String join a string from map
func (slc *mapDurationUint16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapDurationUint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapDurationUint16Var(p *mapDurationUint16Value, name string, value mapDurationUint16Value, usage string, required, isset bool) {
	f.Var(newmapDurationUint16Value(value, p), name, usage, required, isset)
}

// mapDurationUint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapDurationUint16Var(p *mapDurationUint16Value, name string, value mapDurationUint16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapDurationUint16Value(value, p), name, usage, required, isset)
}

// mapDurationUint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapDurationUint16(name string, value mapDurationUint16Value, usage string, required, isset bool) *mapDurationUint16Value {
	p := new(mapDurationUint16Value)
	f.mapDurationUint16Var(p, name, value, usage, required, isset)
	return p
}

// mapDurationUint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapDurationUint16(name string, value mapDurationUint16Value, usage string, required, isset bool) *mapDurationUint16Value {
	return CommandLine.mapDurationUint16(name, value, usage, required, isset)
}

// mapDurationUint32Value []mapDurationUint32Value
type mapDurationUint32Value map[time.Duration]uint32

func newmapDurationUint32Value(val mapDurationUint32Value,
	p *mapDurationUint32Value) *mapDurationUint32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapDurationUint32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapDurationUint32Value) Get() interface{} { return map[time.Duration]uint32(*slc) }

// String join a string from map
func (slc *mapDurationUint32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapDurationUint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapDurationUint32Var(p *mapDurationUint32Value, name string, value mapDurationUint32Value, usage string, required, isset bool) {
	f.Var(newmapDurationUint32Value(value, p), name, usage, required, isset)
}

// mapDurationUint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapDurationUint32Var(p *mapDurationUint32Value, name string, value mapDurationUint32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapDurationUint32Value(value, p), name, usage, required, isset)
}

// mapDurationUint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapDurationUint32(name string, value mapDurationUint32Value, usage string, required, isset bool) *mapDurationUint32Value {
	p := new(mapDurationUint32Value)
	f.mapDurationUint32Var(p, name, value, usage, required, isset)
	return p
}

// mapDurationUint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapDurationUint32(name string, value mapDurationUint32Value, usage string, required, isset bool) *mapDurationUint32Value {
	return CommandLine.mapDurationUint32(name, value, usage, required, isset)
}

// mapDurationUint64Value []mapDurationUint64Value
type mapDurationUint64Value map[time.Duration]uint64

func newmapDurationUint64Value(val mapDurationUint64Value,
	p *mapDurationUint64Value) *mapDurationUint64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapDurationUint64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapDurationUint64Value) Get() interface{} { return map[time.Duration]uint64(*slc) }

// String join a string from map
func (slc *mapDurationUint64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapDurationUint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapDurationUint64Var(p *mapDurationUint64Value, name string, value mapDurationUint64Value, usage string, required, isset bool) {
	f.Var(newmapDurationUint64Value(value, p), name, usage, required, isset)
}

// mapDurationUint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapDurationUint64Var(p *mapDurationUint64Value, name string, value mapDurationUint64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapDurationUint64Value(value, p), name, usage, required, isset)
}

// mapDurationUint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapDurationUint64(name string, value mapDurationUint64Value, usage string, required, isset bool) *mapDurationUint64Value {
	p := new(mapDurationUint64Value)
	f.mapDurationUint64Var(p, name, value, usage, required, isset)
	return p
}

// mapDurationUint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapDurationUint64(name string, value mapDurationUint64Value, usage string, required, isset bool) *mapDurationUint64Value {
	return CommandLine.mapDurationUint64(name, value, usage, required, isset)
}

// mapDurationFloat64Value []mapDurationFloat64Value
type mapDurationFloat64Value map[time.Duration]float64

func newmapDurationFloat64Value(val mapDurationFloat64Value,
	p *mapDurationFloat64Value) *mapDurationFloat64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapDurationFloat64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapDurationFloat64Value) Get() interface{} { return map[time.Duration]float64(*slc) }

// String join a string from map
func (slc *mapDurationFloat64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapDurationFloat64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapDurationFloat64Var(p *mapDurationFloat64Value, name string, value mapDurationFloat64Value, usage string, required, isset bool) {
	f.Var(newmapDurationFloat64Value(value, p), name, usage, required, isset)
}

// mapDurationFloat64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapDurationFloat64Var(p *mapDurationFloat64Value, name string, value mapDurationFloat64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapDurationFloat64Value(value, p), name, usage, required, isset)
}

// mapDurationFloat64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapDurationFloat64(name string, value mapDurationFloat64Value, usage string, required, isset bool) *mapDurationFloat64Value {
	p := new(mapDurationFloat64Value)
	f.mapDurationFloat64Var(p, name, value, usage, required, isset)
	return p
}

// mapDurationFloat64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapDurationFloat64(name string, value mapDurationFloat64Value, usage string, required, isset bool) *mapDurationFloat64Value {
	return CommandLine.mapDurationFloat64(name, value, usage, required, isset)
}

// mapDurationFloat32Value []mapDurationFloat32Value
type mapDurationFloat32Value map[time.Duration]float32

func newmapDurationFloat32Value(val mapDurationFloat32Value,
	p *mapDurationFloat32Value) *mapDurationFloat32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapDurationFloat32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapDurationFloat32Value) Get() interface{} { return map[time.Duration]float32(*slc) }

// String join a string from map
func (slc *mapDurationFloat32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapDurationFloat32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapDurationFloat32Var(p *mapDurationFloat32Value, name string, value mapDurationFloat32Value, usage string, required, isset bool) {
	f.Var(newmapDurationFloat32Value(value, p), name, usage, required, isset)
}

// mapDurationFloat32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapDurationFloat32Var(p *mapDurationFloat32Value, name string, value mapDurationFloat32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapDurationFloat32Value(value, p), name, usage, required, isset)
}

// mapDurationFloat32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapDurationFloat32(name string, value mapDurationFloat32Value, usage string, required, isset bool) *mapDurationFloat32Value {
	p := new(mapDurationFloat32Value)
	f.mapDurationFloat32Var(p, name, value, usage, required, isset)
	return p
}

// mapDurationFloat32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapDurationFloat32(name string, value mapDurationFloat32Value, usage string, required, isset bool) *mapDurationFloat32Value {
	return CommandLine.mapDurationFloat32(name, value, usage, required, isset)
}

// mapDurationBoolValue []mapDurationBoolValue
type mapDurationBoolValue map[time.Duration]bool

func newmapDurationBoolValue(val mapDurationBoolValue,
	p *mapDurationBoolValue) *mapDurationBoolValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapDurationBoolValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapDurationBoolValue) Get() interface{} { return map[time.Duration]bool(*slc) }

// String join a string from map
func (slc *mapDurationBoolValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapDurationBoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapDurationBoolVar(p *mapDurationBoolValue, name string, value mapDurationBoolValue, usage string, required, isset bool) {
	f.Var(newmapDurationBoolValue(value, p), name, usage, required, isset)
}

// mapDurationBoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapDurationBoolVar(p *mapDurationBoolValue, name string, value mapDurationBoolValue, usage string, required, isset bool) {
	CommandLine.Var(newmapDurationBoolValue(value, p), name, usage, required, isset)
}

// mapDurationBoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapDurationBool(name string, value mapDurationBoolValue, usage string, required, isset bool) *mapDurationBoolValue {
	p := new(mapDurationBoolValue)
	f.mapDurationBoolVar(p, name, value, usage, required, isset)
	return p
}

// mapDurationBoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapDurationBool(name string, value mapDurationBoolValue, usage string, required, isset bool) *mapDurationBoolValue {
	return CommandLine.mapDurationBool(name, value, usage, required, isset)
}

// mapDurationStringValue []mapDurationStringValue
type mapDurationStringValue map[time.Duration]string

func newmapDurationStringValue(val mapDurationStringValue,
	p *mapDurationStringValue) *mapDurationStringValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapDurationStringValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapDurationStringValue) Get() interface{} { return map[time.Duration]string(*slc) }

// String join a string from map
func (slc *mapDurationStringValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapDurationStringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapDurationStringVar(p *mapDurationStringValue, name string, value mapDurationStringValue, usage string, required, isset bool) {
	f.Var(newmapDurationStringValue(value, p), name, usage, required, isset)
}

// mapDurationStringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapDurationStringVar(p *mapDurationStringValue, name string, value mapDurationStringValue, usage string, required, isset bool) {
	CommandLine.Var(newmapDurationStringValue(value, p), name, usage, required, isset)
}

// mapDurationStringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapDurationString(name string, value mapDurationStringValue, usage string, required, isset bool) *mapDurationStringValue {
	p := new(mapDurationStringValue)
	f.mapDurationStringVar(p, name, value, usage, required, isset)
	return p
}

// mapDurationStringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapDurationString(name string, value mapDurationStringValue, usage string, required, isset bool) *mapDurationStringValue {
	return CommandLine.mapDurationString(name, value, usage, required, isset)
}

// mapIntDurationValue []mapIntDurationValue
type mapIntDurationValue map[int]time.Duration

func newmapIntDurationValue(val mapIntDurationValue,
	p *mapIntDurationValue) *mapIntDurationValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapIntDurationValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapIntDurationValue) Get() interface{} { return map[int]time.Duration(*slc) }

// String join a string from map
func (slc *mapIntDurationValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapIntDurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapIntDurationVar(p *mapIntDurationValue, name string, value mapIntDurationValue, usage string, required, isset bool) {
	f.Var(newmapIntDurationValue(value, p), name, usage, required, isset)
}

// mapIntDurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapIntDurationVar(p *mapIntDurationValue, name string, value mapIntDurationValue, usage string, required, isset bool) {
	CommandLine.Var(newmapIntDurationValue(value, p), name, usage, required, isset)
}

// mapIntDurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapIntDuration(name string, value mapIntDurationValue, usage string, required, isset bool) *mapIntDurationValue {
	p := new(mapIntDurationValue)
	f.mapIntDurationVar(p, name, value, usage, required, isset)
	return p
}

// mapIntDurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapIntDuration(name string, value mapIntDurationValue, usage string, required, isset bool) *mapIntDurationValue {
	return CommandLine.mapIntDuration(name, value, usage, required, isset)
}

// mapIntIntValue []mapIntIntValue
type mapIntIntValue map[int]int

func newmapIntIntValue(val mapIntIntValue,
	p *mapIntIntValue) *mapIntIntValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapIntIntValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapIntIntValue) Get() interface{} { return map[int]int(*slc) }

// String join a string from map
func (slc *mapIntIntValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapIntIntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapIntIntVar(p *mapIntIntValue, name string, value mapIntIntValue, usage string, required, isset bool) {
	f.Var(newmapIntIntValue(value, p), name, usage, required, isset)
}

// mapIntIntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapIntIntVar(p *mapIntIntValue, name string, value mapIntIntValue, usage string, required, isset bool) {
	CommandLine.Var(newmapIntIntValue(value, p), name, usage, required, isset)
}

// mapIntIntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapIntInt(name string, value mapIntIntValue, usage string, required, isset bool) *mapIntIntValue {
	p := new(mapIntIntValue)
	f.mapIntIntVar(p, name, value, usage, required, isset)
	return p
}

// mapIntIntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapIntInt(name string, value mapIntIntValue, usage string, required, isset bool) *mapIntIntValue {
	return CommandLine.mapIntInt(name, value, usage, required, isset)
}

// mapIntInt8Value []mapIntInt8Value
type mapIntInt8Value map[int]int8

func newmapIntInt8Value(val mapIntInt8Value,
	p *mapIntInt8Value) *mapIntInt8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapIntInt8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapIntInt8Value) Get() interface{} { return map[int]int8(*slc) }

// String join a string from map
func (slc *mapIntInt8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapIntInt8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapIntInt8Var(p *mapIntInt8Value, name string, value mapIntInt8Value, usage string, required, isset bool) {
	f.Var(newmapIntInt8Value(value, p), name, usage, required, isset)
}

// mapIntInt8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapIntInt8Var(p *mapIntInt8Value, name string, value mapIntInt8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapIntInt8Value(value, p), name, usage, required, isset)
}

// mapIntInt8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapIntInt8(name string, value mapIntInt8Value, usage string, required, isset bool) *mapIntInt8Value {
	p := new(mapIntInt8Value)
	f.mapIntInt8Var(p, name, value, usage, required, isset)
	return p
}

// mapIntInt8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapIntInt8(name string, value mapIntInt8Value, usage string, required, isset bool) *mapIntInt8Value {
	return CommandLine.mapIntInt8(name, value, usage, required, isset)
}

// mapIntInt16Value []mapIntInt16Value
type mapIntInt16Value map[int]int16

func newmapIntInt16Value(val mapIntInt16Value,
	p *mapIntInt16Value) *mapIntInt16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapIntInt16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapIntInt16Value) Get() interface{} { return map[int]int16(*slc) }

// String join a string from map
func (slc *mapIntInt16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapIntInt16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapIntInt16Var(p *mapIntInt16Value, name string, value mapIntInt16Value, usage string, required, isset bool) {
	f.Var(newmapIntInt16Value(value, p), name, usage, required, isset)
}

// mapIntInt16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapIntInt16Var(p *mapIntInt16Value, name string, value mapIntInt16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapIntInt16Value(value, p), name, usage, required, isset)
}

// mapIntInt16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapIntInt16(name string, value mapIntInt16Value, usage string, required, isset bool) *mapIntInt16Value {
	p := new(mapIntInt16Value)
	f.mapIntInt16Var(p, name, value, usage, required, isset)
	return p
}

// mapIntInt16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapIntInt16(name string, value mapIntInt16Value, usage string, required, isset bool) *mapIntInt16Value {
	return CommandLine.mapIntInt16(name, value, usage, required, isset)
}

// mapIntInt32Value []mapIntInt32Value
type mapIntInt32Value map[int]int32

func newmapIntInt32Value(val mapIntInt32Value,
	p *mapIntInt32Value) *mapIntInt32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapIntInt32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapIntInt32Value) Get() interface{} { return map[int]int32(*slc) }

// String join a string from map
func (slc *mapIntInt32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapIntInt32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapIntInt32Var(p *mapIntInt32Value, name string, value mapIntInt32Value, usage string, required, isset bool) {
	f.Var(newmapIntInt32Value(value, p), name, usage, required, isset)
}

// mapIntInt32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapIntInt32Var(p *mapIntInt32Value, name string, value mapIntInt32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapIntInt32Value(value, p), name, usage, required, isset)
}

// mapIntInt32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapIntInt32(name string, value mapIntInt32Value, usage string, required, isset bool) *mapIntInt32Value {
	p := new(mapIntInt32Value)
	f.mapIntInt32Var(p, name, value, usage, required, isset)
	return p
}

// mapIntInt32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapIntInt32(name string, value mapIntInt32Value, usage string, required, isset bool) *mapIntInt32Value {
	return CommandLine.mapIntInt32(name, value, usage, required, isset)
}

// mapIntInt64Value []mapIntInt64Value
type mapIntInt64Value map[int]int64

func newmapIntInt64Value(val mapIntInt64Value,
	p *mapIntInt64Value) *mapIntInt64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapIntInt64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapIntInt64Value) Get() interface{} { return map[int]int64(*slc) }

// String join a string from map
func (slc *mapIntInt64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapIntInt64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapIntInt64Var(p *mapIntInt64Value, name string, value mapIntInt64Value, usage string, required, isset bool) {
	f.Var(newmapIntInt64Value(value, p), name, usage, required, isset)
}

// mapIntInt64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapIntInt64Var(p *mapIntInt64Value, name string, value mapIntInt64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapIntInt64Value(value, p), name, usage, required, isset)
}

// mapIntInt64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapIntInt64(name string, value mapIntInt64Value, usage string, required, isset bool) *mapIntInt64Value {
	p := new(mapIntInt64Value)
	f.mapIntInt64Var(p, name, value, usage, required, isset)
	return p
}

// mapIntInt64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapIntInt64(name string, value mapIntInt64Value, usage string, required, isset bool) *mapIntInt64Value {
	return CommandLine.mapIntInt64(name, value, usage, required, isset)
}

// mapIntUintValue []mapIntUintValue
type mapIntUintValue map[int]uint

func newmapIntUintValue(val mapIntUintValue,
	p *mapIntUintValue) *mapIntUintValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapIntUintValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapIntUintValue) Get() interface{} { return map[int]uint(*slc) }

// String join a string from map
func (slc *mapIntUintValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapIntUintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapIntUintVar(p *mapIntUintValue, name string, value mapIntUintValue, usage string, required, isset bool) {
	f.Var(newmapIntUintValue(value, p), name, usage, required, isset)
}

// mapIntUintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapIntUintVar(p *mapIntUintValue, name string, value mapIntUintValue, usage string, required, isset bool) {
	CommandLine.Var(newmapIntUintValue(value, p), name, usage, required, isset)
}

// mapIntUintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapIntUint(name string, value mapIntUintValue, usage string, required, isset bool) *mapIntUintValue {
	p := new(mapIntUintValue)
	f.mapIntUintVar(p, name, value, usage, required, isset)
	return p
}

// mapIntUintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapIntUint(name string, value mapIntUintValue, usage string, required, isset bool) *mapIntUintValue {
	return CommandLine.mapIntUint(name, value, usage, required, isset)
}

// mapIntUint8Value []mapIntUint8Value
type mapIntUint8Value map[int]uint8

func newmapIntUint8Value(val mapIntUint8Value,
	p *mapIntUint8Value) *mapIntUint8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapIntUint8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapIntUint8Value) Get() interface{} { return map[int]uint8(*slc) }

// String join a string from map
func (slc *mapIntUint8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapIntUint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapIntUint8Var(p *mapIntUint8Value, name string, value mapIntUint8Value, usage string, required, isset bool) {
	f.Var(newmapIntUint8Value(value, p), name, usage, required, isset)
}

// mapIntUint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapIntUint8Var(p *mapIntUint8Value, name string, value mapIntUint8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapIntUint8Value(value, p), name, usage, required, isset)
}

// mapIntUint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapIntUint8(name string, value mapIntUint8Value, usage string, required, isset bool) *mapIntUint8Value {
	p := new(mapIntUint8Value)
	f.mapIntUint8Var(p, name, value, usage, required, isset)
	return p
}

// mapIntUint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapIntUint8(name string, value mapIntUint8Value, usage string, required, isset bool) *mapIntUint8Value {
	return CommandLine.mapIntUint8(name, value, usage, required, isset)
}

// mapIntUint16Value []mapIntUint16Value
type mapIntUint16Value map[int]uint16

func newmapIntUint16Value(val mapIntUint16Value,
	p *mapIntUint16Value) *mapIntUint16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapIntUint16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapIntUint16Value) Get() interface{} { return map[int]uint16(*slc) }

// String join a string from map
func (slc *mapIntUint16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapIntUint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapIntUint16Var(p *mapIntUint16Value, name string, value mapIntUint16Value, usage string, required, isset bool) {
	f.Var(newmapIntUint16Value(value, p), name, usage, required, isset)
}

// mapIntUint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapIntUint16Var(p *mapIntUint16Value, name string, value mapIntUint16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapIntUint16Value(value, p), name, usage, required, isset)
}

// mapIntUint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapIntUint16(name string, value mapIntUint16Value, usage string, required, isset bool) *mapIntUint16Value {
	p := new(mapIntUint16Value)
	f.mapIntUint16Var(p, name, value, usage, required, isset)
	return p
}

// mapIntUint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapIntUint16(name string, value mapIntUint16Value, usage string, required, isset bool) *mapIntUint16Value {
	return CommandLine.mapIntUint16(name, value, usage, required, isset)
}

// mapIntUint32Value []mapIntUint32Value
type mapIntUint32Value map[int]uint32

func newmapIntUint32Value(val mapIntUint32Value,
	p *mapIntUint32Value) *mapIntUint32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapIntUint32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapIntUint32Value) Get() interface{} { return map[int]uint32(*slc) }

// String join a string from map
func (slc *mapIntUint32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapIntUint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapIntUint32Var(p *mapIntUint32Value, name string, value mapIntUint32Value, usage string, required, isset bool) {
	f.Var(newmapIntUint32Value(value, p), name, usage, required, isset)
}

// mapIntUint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapIntUint32Var(p *mapIntUint32Value, name string, value mapIntUint32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapIntUint32Value(value, p), name, usage, required, isset)
}

// mapIntUint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapIntUint32(name string, value mapIntUint32Value, usage string, required, isset bool) *mapIntUint32Value {
	p := new(mapIntUint32Value)
	f.mapIntUint32Var(p, name, value, usage, required, isset)
	return p
}

// mapIntUint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapIntUint32(name string, value mapIntUint32Value, usage string, required, isset bool) *mapIntUint32Value {
	return CommandLine.mapIntUint32(name, value, usage, required, isset)
}

// mapIntUint64Value []mapIntUint64Value
type mapIntUint64Value map[int]uint64

func newmapIntUint64Value(val mapIntUint64Value,
	p *mapIntUint64Value) *mapIntUint64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapIntUint64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapIntUint64Value) Get() interface{} { return map[int]uint64(*slc) }

// String join a string from map
func (slc *mapIntUint64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapIntUint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapIntUint64Var(p *mapIntUint64Value, name string, value mapIntUint64Value, usage string, required, isset bool) {
	f.Var(newmapIntUint64Value(value, p), name, usage, required, isset)
}

// mapIntUint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapIntUint64Var(p *mapIntUint64Value, name string, value mapIntUint64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapIntUint64Value(value, p), name, usage, required, isset)
}

// mapIntUint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapIntUint64(name string, value mapIntUint64Value, usage string, required, isset bool) *mapIntUint64Value {
	p := new(mapIntUint64Value)
	f.mapIntUint64Var(p, name, value, usage, required, isset)
	return p
}

// mapIntUint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapIntUint64(name string, value mapIntUint64Value, usage string, required, isset bool) *mapIntUint64Value {
	return CommandLine.mapIntUint64(name, value, usage, required, isset)
}

// mapIntFloat64Value []mapIntFloat64Value
type mapIntFloat64Value map[int]float64

func newmapIntFloat64Value(val mapIntFloat64Value,
	p *mapIntFloat64Value) *mapIntFloat64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapIntFloat64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapIntFloat64Value) Get() interface{} { return map[int]float64(*slc) }

// String join a string from map
func (slc *mapIntFloat64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapIntFloat64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapIntFloat64Var(p *mapIntFloat64Value, name string, value mapIntFloat64Value, usage string, required, isset bool) {
	f.Var(newmapIntFloat64Value(value, p), name, usage, required, isset)
}

// mapIntFloat64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapIntFloat64Var(p *mapIntFloat64Value, name string, value mapIntFloat64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapIntFloat64Value(value, p), name, usage, required, isset)
}

// mapIntFloat64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapIntFloat64(name string, value mapIntFloat64Value, usage string, required, isset bool) *mapIntFloat64Value {
	p := new(mapIntFloat64Value)
	f.mapIntFloat64Var(p, name, value, usage, required, isset)
	return p
}

// mapIntFloat64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapIntFloat64(name string, value mapIntFloat64Value, usage string, required, isset bool) *mapIntFloat64Value {
	return CommandLine.mapIntFloat64(name, value, usage, required, isset)
}

// mapIntFloat32Value []mapIntFloat32Value
type mapIntFloat32Value map[int]float32

func newmapIntFloat32Value(val mapIntFloat32Value,
	p *mapIntFloat32Value) *mapIntFloat32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapIntFloat32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapIntFloat32Value) Get() interface{} { return map[int]float32(*slc) }

// String join a string from map
func (slc *mapIntFloat32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapIntFloat32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapIntFloat32Var(p *mapIntFloat32Value, name string, value mapIntFloat32Value, usage string, required, isset bool) {
	f.Var(newmapIntFloat32Value(value, p), name, usage, required, isset)
}

// mapIntFloat32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapIntFloat32Var(p *mapIntFloat32Value, name string, value mapIntFloat32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapIntFloat32Value(value, p), name, usage, required, isset)
}

// mapIntFloat32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapIntFloat32(name string, value mapIntFloat32Value, usage string, required, isset bool) *mapIntFloat32Value {
	p := new(mapIntFloat32Value)
	f.mapIntFloat32Var(p, name, value, usage, required, isset)
	return p
}

// mapIntFloat32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapIntFloat32(name string, value mapIntFloat32Value, usage string, required, isset bool) *mapIntFloat32Value {
	return CommandLine.mapIntFloat32(name, value, usage, required, isset)
}

// mapIntBoolValue []mapIntBoolValue
type mapIntBoolValue map[int]bool

func newmapIntBoolValue(val mapIntBoolValue,
	p *mapIntBoolValue) *mapIntBoolValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapIntBoolValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapIntBoolValue) Get() interface{} { return map[int]bool(*slc) }

// String join a string from map
func (slc *mapIntBoolValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapIntBoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapIntBoolVar(p *mapIntBoolValue, name string, value mapIntBoolValue, usage string, required, isset bool) {
	f.Var(newmapIntBoolValue(value, p), name, usage, required, isset)
}

// mapIntBoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapIntBoolVar(p *mapIntBoolValue, name string, value mapIntBoolValue, usage string, required, isset bool) {
	CommandLine.Var(newmapIntBoolValue(value, p), name, usage, required, isset)
}

// mapIntBoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapIntBool(name string, value mapIntBoolValue, usage string, required, isset bool) *mapIntBoolValue {
	p := new(mapIntBoolValue)
	f.mapIntBoolVar(p, name, value, usage, required, isset)
	return p
}

// mapIntBoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapIntBool(name string, value mapIntBoolValue, usage string, required, isset bool) *mapIntBoolValue {
	return CommandLine.mapIntBool(name, value, usage, required, isset)
}

// mapIntStringValue []mapIntStringValue
type mapIntStringValue map[int]string

func newmapIntStringValue(val mapIntStringValue,
	p *mapIntStringValue) *mapIntStringValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapIntStringValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapIntStringValue) Get() interface{} { return map[int]string(*slc) }

// String join a string from map
func (slc *mapIntStringValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapIntStringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapIntStringVar(p *mapIntStringValue, name string, value mapIntStringValue, usage string, required, isset bool) {
	f.Var(newmapIntStringValue(value, p), name, usage, required, isset)
}

// mapIntStringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapIntStringVar(p *mapIntStringValue, name string, value mapIntStringValue, usage string, required, isset bool) {
	CommandLine.Var(newmapIntStringValue(value, p), name, usage, required, isset)
}

// mapIntStringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapIntString(name string, value mapIntStringValue, usage string, required, isset bool) *mapIntStringValue {
	p := new(mapIntStringValue)
	f.mapIntStringVar(p, name, value, usage, required, isset)
	return p
}

// mapIntStringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapIntString(name string, value mapIntStringValue, usage string, required, isset bool) *mapIntStringValue {
	return CommandLine.mapIntString(name, value, usage, required, isset)
}

// mapInt8DurationValue []mapInt8DurationValue
type mapInt8DurationValue map[int8]time.Duration

func newmapInt8DurationValue(val mapInt8DurationValue,
	p *mapInt8DurationValue) *mapInt8DurationValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt8DurationValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt8DurationValue) Get() interface{} { return map[int8]time.Duration(*slc) }

// String join a string from map
func (slc *mapInt8DurationValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt8DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt8DurationVar(p *mapInt8DurationValue, name string, value mapInt8DurationValue, usage string, required, isset bool) {
	f.Var(newmapInt8DurationValue(value, p), name, usage, required, isset)
}

// mapInt8DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt8DurationVar(p *mapInt8DurationValue, name string, value mapInt8DurationValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt8DurationValue(value, p), name, usage, required, isset)
}

// mapInt8DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt8Duration(name string, value mapInt8DurationValue, usage string, required, isset bool) *mapInt8DurationValue {
	p := new(mapInt8DurationValue)
	f.mapInt8DurationVar(p, name, value, usage, required, isset)
	return p
}

// mapInt8DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt8Duration(name string, value mapInt8DurationValue, usage string, required, isset bool) *mapInt8DurationValue {
	return CommandLine.mapInt8Duration(name, value, usage, required, isset)
}

// mapInt8IntValue []mapInt8IntValue
type mapInt8IntValue map[int8]int

func newmapInt8IntValue(val mapInt8IntValue,
	p *mapInt8IntValue) *mapInt8IntValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt8IntValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt8IntValue) Get() interface{} { return map[int8]int(*slc) }

// String join a string from map
func (slc *mapInt8IntValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt8IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt8IntVar(p *mapInt8IntValue, name string, value mapInt8IntValue, usage string, required, isset bool) {
	f.Var(newmapInt8IntValue(value, p), name, usage, required, isset)
}

// mapInt8IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt8IntVar(p *mapInt8IntValue, name string, value mapInt8IntValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt8IntValue(value, p), name, usage, required, isset)
}

// mapInt8IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt8Int(name string, value mapInt8IntValue, usage string, required, isset bool) *mapInt8IntValue {
	p := new(mapInt8IntValue)
	f.mapInt8IntVar(p, name, value, usage, required, isset)
	return p
}

// mapInt8IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt8Int(name string, value mapInt8IntValue, usage string, required, isset bool) *mapInt8IntValue {
	return CommandLine.mapInt8Int(name, value, usage, required, isset)
}

// mapInt8Int8Value []mapInt8Int8Value
type mapInt8Int8Value map[int8]int8

func newmapInt8Int8Value(val mapInt8Int8Value,
	p *mapInt8Int8Value) *mapInt8Int8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt8Int8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt8Int8Value) Get() interface{} { return map[int8]int8(*slc) }

// String join a string from map
func (slc *mapInt8Int8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt8Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt8Int8Var(p *mapInt8Int8Value, name string, value mapInt8Int8Value, usage string, required, isset bool) {
	f.Var(newmapInt8Int8Value(value, p), name, usage, required, isset)
}

// mapInt8Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt8Int8Var(p *mapInt8Int8Value, name string, value mapInt8Int8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt8Int8Value(value, p), name, usage, required, isset)
}

// mapInt8Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt8Int8(name string, value mapInt8Int8Value, usage string, required, isset bool) *mapInt8Int8Value {
	p := new(mapInt8Int8Value)
	f.mapInt8Int8Var(p, name, value, usage, required, isset)
	return p
}

// mapInt8Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt8Int8(name string, value mapInt8Int8Value, usage string, required, isset bool) *mapInt8Int8Value {
	return CommandLine.mapInt8Int8(name, value, usage, required, isset)
}

// mapInt8Int16Value []mapInt8Int16Value
type mapInt8Int16Value map[int8]int16

func newmapInt8Int16Value(val mapInt8Int16Value,
	p *mapInt8Int16Value) *mapInt8Int16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt8Int16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt8Int16Value) Get() interface{} { return map[int8]int16(*slc) }

// String join a string from map
func (slc *mapInt8Int16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt8Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt8Int16Var(p *mapInt8Int16Value, name string, value mapInt8Int16Value, usage string, required, isset bool) {
	f.Var(newmapInt8Int16Value(value, p), name, usage, required, isset)
}

// mapInt8Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt8Int16Var(p *mapInt8Int16Value, name string, value mapInt8Int16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt8Int16Value(value, p), name, usage, required, isset)
}

// mapInt8Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt8Int16(name string, value mapInt8Int16Value, usage string, required, isset bool) *mapInt8Int16Value {
	p := new(mapInt8Int16Value)
	f.mapInt8Int16Var(p, name, value, usage, required, isset)
	return p
}

// mapInt8Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt8Int16(name string, value mapInt8Int16Value, usage string, required, isset bool) *mapInt8Int16Value {
	return CommandLine.mapInt8Int16(name, value, usage, required, isset)
}

// mapInt8Int32Value []mapInt8Int32Value
type mapInt8Int32Value map[int8]int32

func newmapInt8Int32Value(val mapInt8Int32Value,
	p *mapInt8Int32Value) *mapInt8Int32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt8Int32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt8Int32Value) Get() interface{} { return map[int8]int32(*slc) }

// String join a string from map
func (slc *mapInt8Int32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt8Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt8Int32Var(p *mapInt8Int32Value, name string, value mapInt8Int32Value, usage string, required, isset bool) {
	f.Var(newmapInt8Int32Value(value, p), name, usage, required, isset)
}

// mapInt8Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt8Int32Var(p *mapInt8Int32Value, name string, value mapInt8Int32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt8Int32Value(value, p), name, usage, required, isset)
}

// mapInt8Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt8Int32(name string, value mapInt8Int32Value, usage string, required, isset bool) *mapInt8Int32Value {
	p := new(mapInt8Int32Value)
	f.mapInt8Int32Var(p, name, value, usage, required, isset)
	return p
}

// mapInt8Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt8Int32(name string, value mapInt8Int32Value, usage string, required, isset bool) *mapInt8Int32Value {
	return CommandLine.mapInt8Int32(name, value, usage, required, isset)
}

// mapInt8Int64Value []mapInt8Int64Value
type mapInt8Int64Value map[int8]int64

func newmapInt8Int64Value(val mapInt8Int64Value,
	p *mapInt8Int64Value) *mapInt8Int64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt8Int64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt8Int64Value) Get() interface{} { return map[int8]int64(*slc) }

// String join a string from map
func (slc *mapInt8Int64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt8Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt8Int64Var(p *mapInt8Int64Value, name string, value mapInt8Int64Value, usage string, required, isset bool) {
	f.Var(newmapInt8Int64Value(value, p), name, usage, required, isset)
}

// mapInt8Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt8Int64Var(p *mapInt8Int64Value, name string, value mapInt8Int64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt8Int64Value(value, p), name, usage, required, isset)
}

// mapInt8Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt8Int64(name string, value mapInt8Int64Value, usage string, required, isset bool) *mapInt8Int64Value {
	p := new(mapInt8Int64Value)
	f.mapInt8Int64Var(p, name, value, usage, required, isset)
	return p
}

// mapInt8Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt8Int64(name string, value mapInt8Int64Value, usage string, required, isset bool) *mapInt8Int64Value {
	return CommandLine.mapInt8Int64(name, value, usage, required, isset)
}

// mapInt8UintValue []mapInt8UintValue
type mapInt8UintValue map[int8]uint

func newmapInt8UintValue(val mapInt8UintValue,
	p *mapInt8UintValue) *mapInt8UintValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt8UintValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt8UintValue) Get() interface{} { return map[int8]uint(*slc) }

// String join a string from map
func (slc *mapInt8UintValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt8UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt8UintVar(p *mapInt8UintValue, name string, value mapInt8UintValue, usage string, required, isset bool) {
	f.Var(newmapInt8UintValue(value, p), name, usage, required, isset)
}

// mapInt8UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt8UintVar(p *mapInt8UintValue, name string, value mapInt8UintValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt8UintValue(value, p), name, usage, required, isset)
}

// mapInt8UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt8Uint(name string, value mapInt8UintValue, usage string, required, isset bool) *mapInt8UintValue {
	p := new(mapInt8UintValue)
	f.mapInt8UintVar(p, name, value, usage, required, isset)
	return p
}

// mapInt8UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt8Uint(name string, value mapInt8UintValue, usage string, required, isset bool) *mapInt8UintValue {
	return CommandLine.mapInt8Uint(name, value, usage, required, isset)
}

// mapInt8Uint8Value []mapInt8Uint8Value
type mapInt8Uint8Value map[int8]uint8

func newmapInt8Uint8Value(val mapInt8Uint8Value,
	p *mapInt8Uint8Value) *mapInt8Uint8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt8Uint8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt8Uint8Value) Get() interface{} { return map[int8]uint8(*slc) }

// String join a string from map
func (slc *mapInt8Uint8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt8Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt8Uint8Var(p *mapInt8Uint8Value, name string, value mapInt8Uint8Value, usage string, required, isset bool) {
	f.Var(newmapInt8Uint8Value(value, p), name, usage, required, isset)
}

// mapInt8Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt8Uint8Var(p *mapInt8Uint8Value, name string, value mapInt8Uint8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt8Uint8Value(value, p), name, usage, required, isset)
}

// mapInt8Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt8Uint8(name string, value mapInt8Uint8Value, usage string, required, isset bool) *mapInt8Uint8Value {
	p := new(mapInt8Uint8Value)
	f.mapInt8Uint8Var(p, name, value, usage, required, isset)
	return p
}

// mapInt8Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt8Uint8(name string, value mapInt8Uint8Value, usage string, required, isset bool) *mapInt8Uint8Value {
	return CommandLine.mapInt8Uint8(name, value, usage, required, isset)
}

// mapInt8Uint16Value []mapInt8Uint16Value
type mapInt8Uint16Value map[int8]uint16

func newmapInt8Uint16Value(val mapInt8Uint16Value,
	p *mapInt8Uint16Value) *mapInt8Uint16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt8Uint16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt8Uint16Value) Get() interface{} { return map[int8]uint16(*slc) }

// String join a string from map
func (slc *mapInt8Uint16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt8Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt8Uint16Var(p *mapInt8Uint16Value, name string, value mapInt8Uint16Value, usage string, required, isset bool) {
	f.Var(newmapInt8Uint16Value(value, p), name, usage, required, isset)
}

// mapInt8Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt8Uint16Var(p *mapInt8Uint16Value, name string, value mapInt8Uint16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt8Uint16Value(value, p), name, usage, required, isset)
}

// mapInt8Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt8Uint16(name string, value mapInt8Uint16Value, usage string, required, isset bool) *mapInt8Uint16Value {
	p := new(mapInt8Uint16Value)
	f.mapInt8Uint16Var(p, name, value, usage, required, isset)
	return p
}

// mapInt8Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt8Uint16(name string, value mapInt8Uint16Value, usage string, required, isset bool) *mapInt8Uint16Value {
	return CommandLine.mapInt8Uint16(name, value, usage, required, isset)
}

// mapInt8Uint32Value []mapInt8Uint32Value
type mapInt8Uint32Value map[int8]uint32

func newmapInt8Uint32Value(val mapInt8Uint32Value,
	p *mapInt8Uint32Value) *mapInt8Uint32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt8Uint32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt8Uint32Value) Get() interface{} { return map[int8]uint32(*slc) }

// String join a string from map
func (slc *mapInt8Uint32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt8Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt8Uint32Var(p *mapInt8Uint32Value, name string, value mapInt8Uint32Value, usage string, required, isset bool) {
	f.Var(newmapInt8Uint32Value(value, p), name, usage, required, isset)
}

// mapInt8Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt8Uint32Var(p *mapInt8Uint32Value, name string, value mapInt8Uint32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt8Uint32Value(value, p), name, usage, required, isset)
}

// mapInt8Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt8Uint32(name string, value mapInt8Uint32Value, usage string, required, isset bool) *mapInt8Uint32Value {
	p := new(mapInt8Uint32Value)
	f.mapInt8Uint32Var(p, name, value, usage, required, isset)
	return p
}

// mapInt8Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt8Uint32(name string, value mapInt8Uint32Value, usage string, required, isset bool) *mapInt8Uint32Value {
	return CommandLine.mapInt8Uint32(name, value, usage, required, isset)
}

// mapInt8Uint64Value []mapInt8Uint64Value
type mapInt8Uint64Value map[int8]uint64

func newmapInt8Uint64Value(val mapInt8Uint64Value,
	p *mapInt8Uint64Value) *mapInt8Uint64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt8Uint64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt8Uint64Value) Get() interface{} { return map[int8]uint64(*slc) }

// String join a string from map
func (slc *mapInt8Uint64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt8Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt8Uint64Var(p *mapInt8Uint64Value, name string, value mapInt8Uint64Value, usage string, required, isset bool) {
	f.Var(newmapInt8Uint64Value(value, p), name, usage, required, isset)
}

// mapInt8Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt8Uint64Var(p *mapInt8Uint64Value, name string, value mapInt8Uint64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt8Uint64Value(value, p), name, usage, required, isset)
}

// mapInt8Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt8Uint64(name string, value mapInt8Uint64Value, usage string, required, isset bool) *mapInt8Uint64Value {
	p := new(mapInt8Uint64Value)
	f.mapInt8Uint64Var(p, name, value, usage, required, isset)
	return p
}

// mapInt8Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt8Uint64(name string, value mapInt8Uint64Value, usage string, required, isset bool) *mapInt8Uint64Value {
	return CommandLine.mapInt8Uint64(name, value, usage, required, isset)
}

// mapInt8Float64Value []mapInt8Float64Value
type mapInt8Float64Value map[int8]float64

func newmapInt8Float64Value(val mapInt8Float64Value,
	p *mapInt8Float64Value) *mapInt8Float64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt8Float64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt8Float64Value) Get() interface{} { return map[int8]float64(*slc) }

// String join a string from map
func (slc *mapInt8Float64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt8Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt8Float64Var(p *mapInt8Float64Value, name string, value mapInt8Float64Value, usage string, required, isset bool) {
	f.Var(newmapInt8Float64Value(value, p), name, usage, required, isset)
}

// mapInt8Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt8Float64Var(p *mapInt8Float64Value, name string, value mapInt8Float64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt8Float64Value(value, p), name, usage, required, isset)
}

// mapInt8Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt8Float64(name string, value mapInt8Float64Value, usage string, required, isset bool) *mapInt8Float64Value {
	p := new(mapInt8Float64Value)
	f.mapInt8Float64Var(p, name, value, usage, required, isset)
	return p
}

// mapInt8Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt8Float64(name string, value mapInt8Float64Value, usage string, required, isset bool) *mapInt8Float64Value {
	return CommandLine.mapInt8Float64(name, value, usage, required, isset)
}

// mapInt8Float32Value []mapInt8Float32Value
type mapInt8Float32Value map[int8]float32

func newmapInt8Float32Value(val mapInt8Float32Value,
	p *mapInt8Float32Value) *mapInt8Float32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt8Float32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt8Float32Value) Get() interface{} { return map[int8]float32(*slc) }

// String join a string from map
func (slc *mapInt8Float32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt8Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt8Float32Var(p *mapInt8Float32Value, name string, value mapInt8Float32Value, usage string, required, isset bool) {
	f.Var(newmapInt8Float32Value(value, p), name, usage, required, isset)
}

// mapInt8Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt8Float32Var(p *mapInt8Float32Value, name string, value mapInt8Float32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt8Float32Value(value, p), name, usage, required, isset)
}

// mapInt8Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt8Float32(name string, value mapInt8Float32Value, usage string, required, isset bool) *mapInt8Float32Value {
	p := new(mapInt8Float32Value)
	f.mapInt8Float32Var(p, name, value, usage, required, isset)
	return p
}

// mapInt8Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt8Float32(name string, value mapInt8Float32Value, usage string, required, isset bool) *mapInt8Float32Value {
	return CommandLine.mapInt8Float32(name, value, usage, required, isset)
}

// mapInt8BoolValue []mapInt8BoolValue
type mapInt8BoolValue map[int8]bool

func newmapInt8BoolValue(val mapInt8BoolValue,
	p *mapInt8BoolValue) *mapInt8BoolValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt8BoolValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt8BoolValue) Get() interface{} { return map[int8]bool(*slc) }

// String join a string from map
func (slc *mapInt8BoolValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt8BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt8BoolVar(p *mapInt8BoolValue, name string, value mapInt8BoolValue, usage string, required, isset bool) {
	f.Var(newmapInt8BoolValue(value, p), name, usage, required, isset)
}

// mapInt8BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt8BoolVar(p *mapInt8BoolValue, name string, value mapInt8BoolValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt8BoolValue(value, p), name, usage, required, isset)
}

// mapInt8BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt8Bool(name string, value mapInt8BoolValue, usage string, required, isset bool) *mapInt8BoolValue {
	p := new(mapInt8BoolValue)
	f.mapInt8BoolVar(p, name, value, usage, required, isset)
	return p
}

// mapInt8BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt8Bool(name string, value mapInt8BoolValue, usage string, required, isset bool) *mapInt8BoolValue {
	return CommandLine.mapInt8Bool(name, value, usage, required, isset)
}

// mapInt8StringValue []mapInt8StringValue
type mapInt8StringValue map[int8]string

func newmapInt8StringValue(val mapInt8StringValue,
	p *mapInt8StringValue) *mapInt8StringValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt8StringValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt8StringValue) Get() interface{} { return map[int8]string(*slc) }

// String join a string from map
func (slc *mapInt8StringValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt8StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt8StringVar(p *mapInt8StringValue, name string, value mapInt8StringValue, usage string, required, isset bool) {
	f.Var(newmapInt8StringValue(value, p), name, usage, required, isset)
}

// mapInt8StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt8StringVar(p *mapInt8StringValue, name string, value mapInt8StringValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt8StringValue(value, p), name, usage, required, isset)
}

// mapInt8StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt8String(name string, value mapInt8StringValue, usage string, required, isset bool) *mapInt8StringValue {
	p := new(mapInt8StringValue)
	f.mapInt8StringVar(p, name, value, usage, required, isset)
	return p
}

// mapInt8StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt8String(name string, value mapInt8StringValue, usage string, required, isset bool) *mapInt8StringValue {
	return CommandLine.mapInt8String(name, value, usage, required, isset)
}

// mapInt16DurationValue []mapInt16DurationValue
type mapInt16DurationValue map[int16]time.Duration

func newmapInt16DurationValue(val mapInt16DurationValue,
	p *mapInt16DurationValue) *mapInt16DurationValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt16DurationValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt16DurationValue) Get() interface{} { return map[int16]time.Duration(*slc) }

// String join a string from map
func (slc *mapInt16DurationValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt16DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt16DurationVar(p *mapInt16DurationValue, name string, value mapInt16DurationValue, usage string, required, isset bool) {
	f.Var(newmapInt16DurationValue(value, p), name, usage, required, isset)
}

// mapInt16DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt16DurationVar(p *mapInt16DurationValue, name string, value mapInt16DurationValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt16DurationValue(value, p), name, usage, required, isset)
}

// mapInt16DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt16Duration(name string, value mapInt16DurationValue, usage string, required, isset bool) *mapInt16DurationValue {
	p := new(mapInt16DurationValue)
	f.mapInt16DurationVar(p, name, value, usage, required, isset)
	return p
}

// mapInt16DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt16Duration(name string, value mapInt16DurationValue, usage string, required, isset bool) *mapInt16DurationValue {
	return CommandLine.mapInt16Duration(name, value, usage, required, isset)
}

// mapInt16IntValue []mapInt16IntValue
type mapInt16IntValue map[int16]int

func newmapInt16IntValue(val mapInt16IntValue,
	p *mapInt16IntValue) *mapInt16IntValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt16IntValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt16IntValue) Get() interface{} { return map[int16]int(*slc) }

// String join a string from map
func (slc *mapInt16IntValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt16IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt16IntVar(p *mapInt16IntValue, name string, value mapInt16IntValue, usage string, required, isset bool) {
	f.Var(newmapInt16IntValue(value, p), name, usage, required, isset)
}

// mapInt16IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt16IntVar(p *mapInt16IntValue, name string, value mapInt16IntValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt16IntValue(value, p), name, usage, required, isset)
}

// mapInt16IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt16Int(name string, value mapInt16IntValue, usage string, required, isset bool) *mapInt16IntValue {
	p := new(mapInt16IntValue)
	f.mapInt16IntVar(p, name, value, usage, required, isset)
	return p
}

// mapInt16IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt16Int(name string, value mapInt16IntValue, usage string, required, isset bool) *mapInt16IntValue {
	return CommandLine.mapInt16Int(name, value, usage, required, isset)
}

// mapInt16Int8Value []mapInt16Int8Value
type mapInt16Int8Value map[int16]int8

func newmapInt16Int8Value(val mapInt16Int8Value,
	p *mapInt16Int8Value) *mapInt16Int8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt16Int8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt16Int8Value) Get() interface{} { return map[int16]int8(*slc) }

// String join a string from map
func (slc *mapInt16Int8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt16Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt16Int8Var(p *mapInt16Int8Value, name string, value mapInt16Int8Value, usage string, required, isset bool) {
	f.Var(newmapInt16Int8Value(value, p), name, usage, required, isset)
}

// mapInt16Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt16Int8Var(p *mapInt16Int8Value, name string, value mapInt16Int8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt16Int8Value(value, p), name, usage, required, isset)
}

// mapInt16Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt16Int8(name string, value mapInt16Int8Value, usage string, required, isset bool) *mapInt16Int8Value {
	p := new(mapInt16Int8Value)
	f.mapInt16Int8Var(p, name, value, usage, required, isset)
	return p
}

// mapInt16Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt16Int8(name string, value mapInt16Int8Value, usage string, required, isset bool) *mapInt16Int8Value {
	return CommandLine.mapInt16Int8(name, value, usage, required, isset)
}

// mapInt16Int16Value []mapInt16Int16Value
type mapInt16Int16Value map[int16]int16

func newmapInt16Int16Value(val mapInt16Int16Value,
	p *mapInt16Int16Value) *mapInt16Int16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt16Int16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt16Int16Value) Get() interface{} { return map[int16]int16(*slc) }

// String join a string from map
func (slc *mapInt16Int16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt16Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt16Int16Var(p *mapInt16Int16Value, name string, value mapInt16Int16Value, usage string, required, isset bool) {
	f.Var(newmapInt16Int16Value(value, p), name, usage, required, isset)
}

// mapInt16Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt16Int16Var(p *mapInt16Int16Value, name string, value mapInt16Int16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt16Int16Value(value, p), name, usage, required, isset)
}

// mapInt16Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt16Int16(name string, value mapInt16Int16Value, usage string, required, isset bool) *mapInt16Int16Value {
	p := new(mapInt16Int16Value)
	f.mapInt16Int16Var(p, name, value, usage, required, isset)
	return p
}

// mapInt16Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt16Int16(name string, value mapInt16Int16Value, usage string, required, isset bool) *mapInt16Int16Value {
	return CommandLine.mapInt16Int16(name, value, usage, required, isset)
}

// mapInt16Int32Value []mapInt16Int32Value
type mapInt16Int32Value map[int16]int32

func newmapInt16Int32Value(val mapInt16Int32Value,
	p *mapInt16Int32Value) *mapInt16Int32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt16Int32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt16Int32Value) Get() interface{} { return map[int16]int32(*slc) }

// String join a string from map
func (slc *mapInt16Int32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt16Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt16Int32Var(p *mapInt16Int32Value, name string, value mapInt16Int32Value, usage string, required, isset bool) {
	f.Var(newmapInt16Int32Value(value, p), name, usage, required, isset)
}

// mapInt16Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt16Int32Var(p *mapInt16Int32Value, name string, value mapInt16Int32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt16Int32Value(value, p), name, usage, required, isset)
}

// mapInt16Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt16Int32(name string, value mapInt16Int32Value, usage string, required, isset bool) *mapInt16Int32Value {
	p := new(mapInt16Int32Value)
	f.mapInt16Int32Var(p, name, value, usage, required, isset)
	return p
}

// mapInt16Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt16Int32(name string, value mapInt16Int32Value, usage string, required, isset bool) *mapInt16Int32Value {
	return CommandLine.mapInt16Int32(name, value, usage, required, isset)
}

// mapInt16Int64Value []mapInt16Int64Value
type mapInt16Int64Value map[int16]int64

func newmapInt16Int64Value(val mapInt16Int64Value,
	p *mapInt16Int64Value) *mapInt16Int64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt16Int64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt16Int64Value) Get() interface{} { return map[int16]int64(*slc) }

// String join a string from map
func (slc *mapInt16Int64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt16Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt16Int64Var(p *mapInt16Int64Value, name string, value mapInt16Int64Value, usage string, required, isset bool) {
	f.Var(newmapInt16Int64Value(value, p), name, usage, required, isset)
}

// mapInt16Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt16Int64Var(p *mapInt16Int64Value, name string, value mapInt16Int64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt16Int64Value(value, p), name, usage, required, isset)
}

// mapInt16Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt16Int64(name string, value mapInt16Int64Value, usage string, required, isset bool) *mapInt16Int64Value {
	p := new(mapInt16Int64Value)
	f.mapInt16Int64Var(p, name, value, usage, required, isset)
	return p
}

// mapInt16Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt16Int64(name string, value mapInt16Int64Value, usage string, required, isset bool) *mapInt16Int64Value {
	return CommandLine.mapInt16Int64(name, value, usage, required, isset)
}

// mapInt16UintValue []mapInt16UintValue
type mapInt16UintValue map[int16]uint

func newmapInt16UintValue(val mapInt16UintValue,
	p *mapInt16UintValue) *mapInt16UintValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt16UintValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt16UintValue) Get() interface{} { return map[int16]uint(*slc) }

// String join a string from map
func (slc *mapInt16UintValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt16UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt16UintVar(p *mapInt16UintValue, name string, value mapInt16UintValue, usage string, required, isset bool) {
	f.Var(newmapInt16UintValue(value, p), name, usage, required, isset)
}

// mapInt16UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt16UintVar(p *mapInt16UintValue, name string, value mapInt16UintValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt16UintValue(value, p), name, usage, required, isset)
}

// mapInt16UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt16Uint(name string, value mapInt16UintValue, usage string, required, isset bool) *mapInt16UintValue {
	p := new(mapInt16UintValue)
	f.mapInt16UintVar(p, name, value, usage, required, isset)
	return p
}

// mapInt16UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt16Uint(name string, value mapInt16UintValue, usage string, required, isset bool) *mapInt16UintValue {
	return CommandLine.mapInt16Uint(name, value, usage, required, isset)
}

// mapInt16Uint8Value []mapInt16Uint8Value
type mapInt16Uint8Value map[int16]uint8

func newmapInt16Uint8Value(val mapInt16Uint8Value,
	p *mapInt16Uint8Value) *mapInt16Uint8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt16Uint8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt16Uint8Value) Get() interface{} { return map[int16]uint8(*slc) }

// String join a string from map
func (slc *mapInt16Uint8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt16Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt16Uint8Var(p *mapInt16Uint8Value, name string, value mapInt16Uint8Value, usage string, required, isset bool) {
	f.Var(newmapInt16Uint8Value(value, p), name, usage, required, isset)
}

// mapInt16Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt16Uint8Var(p *mapInt16Uint8Value, name string, value mapInt16Uint8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt16Uint8Value(value, p), name, usage, required, isset)
}

// mapInt16Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt16Uint8(name string, value mapInt16Uint8Value, usage string, required, isset bool) *mapInt16Uint8Value {
	p := new(mapInt16Uint8Value)
	f.mapInt16Uint8Var(p, name, value, usage, required, isset)
	return p
}

// mapInt16Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt16Uint8(name string, value mapInt16Uint8Value, usage string, required, isset bool) *mapInt16Uint8Value {
	return CommandLine.mapInt16Uint8(name, value, usage, required, isset)
}

// mapInt16Uint16Value []mapInt16Uint16Value
type mapInt16Uint16Value map[int16]uint16

func newmapInt16Uint16Value(val mapInt16Uint16Value,
	p *mapInt16Uint16Value) *mapInt16Uint16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt16Uint16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt16Uint16Value) Get() interface{} { return map[int16]uint16(*slc) }

// String join a string from map
func (slc *mapInt16Uint16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt16Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt16Uint16Var(p *mapInt16Uint16Value, name string, value mapInt16Uint16Value, usage string, required, isset bool) {
	f.Var(newmapInt16Uint16Value(value, p), name, usage, required, isset)
}

// mapInt16Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt16Uint16Var(p *mapInt16Uint16Value, name string, value mapInt16Uint16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt16Uint16Value(value, p), name, usage, required, isset)
}

// mapInt16Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt16Uint16(name string, value mapInt16Uint16Value, usage string, required, isset bool) *mapInt16Uint16Value {
	p := new(mapInt16Uint16Value)
	f.mapInt16Uint16Var(p, name, value, usage, required, isset)
	return p
}

// mapInt16Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt16Uint16(name string, value mapInt16Uint16Value, usage string, required, isset bool) *mapInt16Uint16Value {
	return CommandLine.mapInt16Uint16(name, value, usage, required, isset)
}

// mapInt16Uint32Value []mapInt16Uint32Value
type mapInt16Uint32Value map[int16]uint32

func newmapInt16Uint32Value(val mapInt16Uint32Value,
	p *mapInt16Uint32Value) *mapInt16Uint32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt16Uint32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt16Uint32Value) Get() interface{} { return map[int16]uint32(*slc) }

// String join a string from map
func (slc *mapInt16Uint32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt16Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt16Uint32Var(p *mapInt16Uint32Value, name string, value mapInt16Uint32Value, usage string, required, isset bool) {
	f.Var(newmapInt16Uint32Value(value, p), name, usage, required, isset)
}

// mapInt16Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt16Uint32Var(p *mapInt16Uint32Value, name string, value mapInt16Uint32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt16Uint32Value(value, p), name, usage, required, isset)
}

// mapInt16Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt16Uint32(name string, value mapInt16Uint32Value, usage string, required, isset bool) *mapInt16Uint32Value {
	p := new(mapInt16Uint32Value)
	f.mapInt16Uint32Var(p, name, value, usage, required, isset)
	return p
}

// mapInt16Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt16Uint32(name string, value mapInt16Uint32Value, usage string, required, isset bool) *mapInt16Uint32Value {
	return CommandLine.mapInt16Uint32(name, value, usage, required, isset)
}

// mapInt16Uint64Value []mapInt16Uint64Value
type mapInt16Uint64Value map[int16]uint64

func newmapInt16Uint64Value(val mapInt16Uint64Value,
	p *mapInt16Uint64Value) *mapInt16Uint64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt16Uint64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt16Uint64Value) Get() interface{} { return map[int16]uint64(*slc) }

// String join a string from map
func (slc *mapInt16Uint64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt16Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt16Uint64Var(p *mapInt16Uint64Value, name string, value mapInt16Uint64Value, usage string, required, isset bool) {
	f.Var(newmapInt16Uint64Value(value, p), name, usage, required, isset)
}

// mapInt16Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt16Uint64Var(p *mapInt16Uint64Value, name string, value mapInt16Uint64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt16Uint64Value(value, p), name, usage, required, isset)
}

// mapInt16Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt16Uint64(name string, value mapInt16Uint64Value, usage string, required, isset bool) *mapInt16Uint64Value {
	p := new(mapInt16Uint64Value)
	f.mapInt16Uint64Var(p, name, value, usage, required, isset)
	return p
}

// mapInt16Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt16Uint64(name string, value mapInt16Uint64Value, usage string, required, isset bool) *mapInt16Uint64Value {
	return CommandLine.mapInt16Uint64(name, value, usage, required, isset)
}

// mapInt16Float64Value []mapInt16Float64Value
type mapInt16Float64Value map[int16]float64

func newmapInt16Float64Value(val mapInt16Float64Value,
	p *mapInt16Float64Value) *mapInt16Float64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt16Float64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt16Float64Value) Get() interface{} { return map[int16]float64(*slc) }

// String join a string from map
func (slc *mapInt16Float64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt16Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt16Float64Var(p *mapInt16Float64Value, name string, value mapInt16Float64Value, usage string, required, isset bool) {
	f.Var(newmapInt16Float64Value(value, p), name, usage, required, isset)
}

// mapInt16Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt16Float64Var(p *mapInt16Float64Value, name string, value mapInt16Float64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt16Float64Value(value, p), name, usage, required, isset)
}

// mapInt16Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt16Float64(name string, value mapInt16Float64Value, usage string, required, isset bool) *mapInt16Float64Value {
	p := new(mapInt16Float64Value)
	f.mapInt16Float64Var(p, name, value, usage, required, isset)
	return p
}

// mapInt16Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt16Float64(name string, value mapInt16Float64Value, usage string, required, isset bool) *mapInt16Float64Value {
	return CommandLine.mapInt16Float64(name, value, usage, required, isset)
}

// mapInt16Float32Value []mapInt16Float32Value
type mapInt16Float32Value map[int16]float32

func newmapInt16Float32Value(val mapInt16Float32Value,
	p *mapInt16Float32Value) *mapInt16Float32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt16Float32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt16Float32Value) Get() interface{} { return map[int16]float32(*slc) }

// String join a string from map
func (slc *mapInt16Float32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt16Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt16Float32Var(p *mapInt16Float32Value, name string, value mapInt16Float32Value, usage string, required, isset bool) {
	f.Var(newmapInt16Float32Value(value, p), name, usage, required, isset)
}

// mapInt16Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt16Float32Var(p *mapInt16Float32Value, name string, value mapInt16Float32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt16Float32Value(value, p), name, usage, required, isset)
}

// mapInt16Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt16Float32(name string, value mapInt16Float32Value, usage string, required, isset bool) *mapInt16Float32Value {
	p := new(mapInt16Float32Value)
	f.mapInt16Float32Var(p, name, value, usage, required, isset)
	return p
}

// mapInt16Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt16Float32(name string, value mapInt16Float32Value, usage string, required, isset bool) *mapInt16Float32Value {
	return CommandLine.mapInt16Float32(name, value, usage, required, isset)
}

// mapInt16BoolValue []mapInt16BoolValue
type mapInt16BoolValue map[int16]bool

func newmapInt16BoolValue(val mapInt16BoolValue,
	p *mapInt16BoolValue) *mapInt16BoolValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt16BoolValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt16BoolValue) Get() interface{} { return map[int16]bool(*slc) }

// String join a string from map
func (slc *mapInt16BoolValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt16BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt16BoolVar(p *mapInt16BoolValue, name string, value mapInt16BoolValue, usage string, required, isset bool) {
	f.Var(newmapInt16BoolValue(value, p), name, usage, required, isset)
}

// mapInt16BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt16BoolVar(p *mapInt16BoolValue, name string, value mapInt16BoolValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt16BoolValue(value, p), name, usage, required, isset)
}

// mapInt16BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt16Bool(name string, value mapInt16BoolValue, usage string, required, isset bool) *mapInt16BoolValue {
	p := new(mapInt16BoolValue)
	f.mapInt16BoolVar(p, name, value, usage, required, isset)
	return p
}

// mapInt16BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt16Bool(name string, value mapInt16BoolValue, usage string, required, isset bool) *mapInt16BoolValue {
	return CommandLine.mapInt16Bool(name, value, usage, required, isset)
}

// mapInt16StringValue []mapInt16StringValue
type mapInt16StringValue map[int16]string

func newmapInt16StringValue(val mapInt16StringValue,
	p *mapInt16StringValue) *mapInt16StringValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt16StringValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt16StringValue) Get() interface{} { return map[int16]string(*slc) }

// String join a string from map
func (slc *mapInt16StringValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt16StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt16StringVar(p *mapInt16StringValue, name string, value mapInt16StringValue, usage string, required, isset bool) {
	f.Var(newmapInt16StringValue(value, p), name, usage, required, isset)
}

// mapInt16StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt16StringVar(p *mapInt16StringValue, name string, value mapInt16StringValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt16StringValue(value, p), name, usage, required, isset)
}

// mapInt16StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt16String(name string, value mapInt16StringValue, usage string, required, isset bool) *mapInt16StringValue {
	p := new(mapInt16StringValue)
	f.mapInt16StringVar(p, name, value, usage, required, isset)
	return p
}

// mapInt16StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt16String(name string, value mapInt16StringValue, usage string, required, isset bool) *mapInt16StringValue {
	return CommandLine.mapInt16String(name, value, usage, required, isset)
}

// mapInt32DurationValue []mapInt32DurationValue
type mapInt32DurationValue map[int32]time.Duration

func newmapInt32DurationValue(val mapInt32DurationValue,
	p *mapInt32DurationValue) *mapInt32DurationValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt32DurationValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt32DurationValue) Get() interface{} { return map[int32]time.Duration(*slc) }

// String join a string from map
func (slc *mapInt32DurationValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt32DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt32DurationVar(p *mapInt32DurationValue, name string, value mapInt32DurationValue, usage string, required, isset bool) {
	f.Var(newmapInt32DurationValue(value, p), name, usage, required, isset)
}

// mapInt32DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt32DurationVar(p *mapInt32DurationValue, name string, value mapInt32DurationValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt32DurationValue(value, p), name, usage, required, isset)
}

// mapInt32DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt32Duration(name string, value mapInt32DurationValue, usage string, required, isset bool) *mapInt32DurationValue {
	p := new(mapInt32DurationValue)
	f.mapInt32DurationVar(p, name, value, usage, required, isset)
	return p
}

// mapInt32DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt32Duration(name string, value mapInt32DurationValue, usage string, required, isset bool) *mapInt32DurationValue {
	return CommandLine.mapInt32Duration(name, value, usage, required, isset)
}

// mapInt32IntValue []mapInt32IntValue
type mapInt32IntValue map[int32]int

func newmapInt32IntValue(val mapInt32IntValue,
	p *mapInt32IntValue) *mapInt32IntValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt32IntValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt32IntValue) Get() interface{} { return map[int32]int(*slc) }

// String join a string from map
func (slc *mapInt32IntValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt32IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt32IntVar(p *mapInt32IntValue, name string, value mapInt32IntValue, usage string, required, isset bool) {
	f.Var(newmapInt32IntValue(value, p), name, usage, required, isset)
}

// mapInt32IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt32IntVar(p *mapInt32IntValue, name string, value mapInt32IntValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt32IntValue(value, p), name, usage, required, isset)
}

// mapInt32IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt32Int(name string, value mapInt32IntValue, usage string, required, isset bool) *mapInt32IntValue {
	p := new(mapInt32IntValue)
	f.mapInt32IntVar(p, name, value, usage, required, isset)
	return p
}

// mapInt32IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt32Int(name string, value mapInt32IntValue, usage string, required, isset bool) *mapInt32IntValue {
	return CommandLine.mapInt32Int(name, value, usage, required, isset)
}

// mapInt32Int8Value []mapInt32Int8Value
type mapInt32Int8Value map[int32]int8

func newmapInt32Int8Value(val mapInt32Int8Value,
	p *mapInt32Int8Value) *mapInt32Int8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt32Int8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt32Int8Value) Get() interface{} { return map[int32]int8(*slc) }

// String join a string from map
func (slc *mapInt32Int8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt32Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt32Int8Var(p *mapInt32Int8Value, name string, value mapInt32Int8Value, usage string, required, isset bool) {
	f.Var(newmapInt32Int8Value(value, p), name, usage, required, isset)
}

// mapInt32Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt32Int8Var(p *mapInt32Int8Value, name string, value mapInt32Int8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt32Int8Value(value, p), name, usage, required, isset)
}

// mapInt32Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt32Int8(name string, value mapInt32Int8Value, usage string, required, isset bool) *mapInt32Int8Value {
	p := new(mapInt32Int8Value)
	f.mapInt32Int8Var(p, name, value, usage, required, isset)
	return p
}

// mapInt32Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt32Int8(name string, value mapInt32Int8Value, usage string, required, isset bool) *mapInt32Int8Value {
	return CommandLine.mapInt32Int8(name, value, usage, required, isset)
}

// mapInt32Int16Value []mapInt32Int16Value
type mapInt32Int16Value map[int32]int16

func newmapInt32Int16Value(val mapInt32Int16Value,
	p *mapInt32Int16Value) *mapInt32Int16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt32Int16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt32Int16Value) Get() interface{} { return map[int32]int16(*slc) }

// String join a string from map
func (slc *mapInt32Int16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt32Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt32Int16Var(p *mapInt32Int16Value, name string, value mapInt32Int16Value, usage string, required, isset bool) {
	f.Var(newmapInt32Int16Value(value, p), name, usage, required, isset)
}

// mapInt32Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt32Int16Var(p *mapInt32Int16Value, name string, value mapInt32Int16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt32Int16Value(value, p), name, usage, required, isset)
}

// mapInt32Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt32Int16(name string, value mapInt32Int16Value, usage string, required, isset bool) *mapInt32Int16Value {
	p := new(mapInt32Int16Value)
	f.mapInt32Int16Var(p, name, value, usage, required, isset)
	return p
}

// mapInt32Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt32Int16(name string, value mapInt32Int16Value, usage string, required, isset bool) *mapInt32Int16Value {
	return CommandLine.mapInt32Int16(name, value, usage, required, isset)
}

// mapInt32Int32Value []mapInt32Int32Value
type mapInt32Int32Value map[int32]int32

func newmapInt32Int32Value(val mapInt32Int32Value,
	p *mapInt32Int32Value) *mapInt32Int32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt32Int32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt32Int32Value) Get() interface{} { return map[int32]int32(*slc) }

// String join a string from map
func (slc *mapInt32Int32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt32Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt32Int32Var(p *mapInt32Int32Value, name string, value mapInt32Int32Value, usage string, required, isset bool) {
	f.Var(newmapInt32Int32Value(value, p), name, usage, required, isset)
}

// mapInt32Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt32Int32Var(p *mapInt32Int32Value, name string, value mapInt32Int32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt32Int32Value(value, p), name, usage, required, isset)
}

// mapInt32Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt32Int32(name string, value mapInt32Int32Value, usage string, required, isset bool) *mapInt32Int32Value {
	p := new(mapInt32Int32Value)
	f.mapInt32Int32Var(p, name, value, usage, required, isset)
	return p
}

// mapInt32Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt32Int32(name string, value mapInt32Int32Value, usage string, required, isset bool) *mapInt32Int32Value {
	return CommandLine.mapInt32Int32(name, value, usage, required, isset)
}

// mapInt32Int64Value []mapInt32Int64Value
type mapInt32Int64Value map[int32]int64

func newmapInt32Int64Value(val mapInt32Int64Value,
	p *mapInt32Int64Value) *mapInt32Int64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt32Int64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt32Int64Value) Get() interface{} { return map[int32]int64(*slc) }

// String join a string from map
func (slc *mapInt32Int64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt32Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt32Int64Var(p *mapInt32Int64Value, name string, value mapInt32Int64Value, usage string, required, isset bool) {
	f.Var(newmapInt32Int64Value(value, p), name, usage, required, isset)
}

// mapInt32Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt32Int64Var(p *mapInt32Int64Value, name string, value mapInt32Int64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt32Int64Value(value, p), name, usage, required, isset)
}

// mapInt32Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt32Int64(name string, value mapInt32Int64Value, usage string, required, isset bool) *mapInt32Int64Value {
	p := new(mapInt32Int64Value)
	f.mapInt32Int64Var(p, name, value, usage, required, isset)
	return p
}

// mapInt32Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt32Int64(name string, value mapInt32Int64Value, usage string, required, isset bool) *mapInt32Int64Value {
	return CommandLine.mapInt32Int64(name, value, usage, required, isset)
}

// mapInt32UintValue []mapInt32UintValue
type mapInt32UintValue map[int32]uint

func newmapInt32UintValue(val mapInt32UintValue,
	p *mapInt32UintValue) *mapInt32UintValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt32UintValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt32UintValue) Get() interface{} { return map[int32]uint(*slc) }

// String join a string from map
func (slc *mapInt32UintValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt32UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt32UintVar(p *mapInt32UintValue, name string, value mapInt32UintValue, usage string, required, isset bool) {
	f.Var(newmapInt32UintValue(value, p), name, usage, required, isset)
}

// mapInt32UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt32UintVar(p *mapInt32UintValue, name string, value mapInt32UintValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt32UintValue(value, p), name, usage, required, isset)
}

// mapInt32UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt32Uint(name string, value mapInt32UintValue, usage string, required, isset bool) *mapInt32UintValue {
	p := new(mapInt32UintValue)
	f.mapInt32UintVar(p, name, value, usage, required, isset)
	return p
}

// mapInt32UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt32Uint(name string, value mapInt32UintValue, usage string, required, isset bool) *mapInt32UintValue {
	return CommandLine.mapInt32Uint(name, value, usage, required, isset)
}

// mapInt32Uint8Value []mapInt32Uint8Value
type mapInt32Uint8Value map[int32]uint8

func newmapInt32Uint8Value(val mapInt32Uint8Value,
	p *mapInt32Uint8Value) *mapInt32Uint8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt32Uint8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt32Uint8Value) Get() interface{} { return map[int32]uint8(*slc) }

// String join a string from map
func (slc *mapInt32Uint8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt32Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt32Uint8Var(p *mapInt32Uint8Value, name string, value mapInt32Uint8Value, usage string, required, isset bool) {
	f.Var(newmapInt32Uint8Value(value, p), name, usage, required, isset)
}

// mapInt32Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt32Uint8Var(p *mapInt32Uint8Value, name string, value mapInt32Uint8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt32Uint8Value(value, p), name, usage, required, isset)
}

// mapInt32Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt32Uint8(name string, value mapInt32Uint8Value, usage string, required, isset bool) *mapInt32Uint8Value {
	p := new(mapInt32Uint8Value)
	f.mapInt32Uint8Var(p, name, value, usage, required, isset)
	return p
}

// mapInt32Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt32Uint8(name string, value mapInt32Uint8Value, usage string, required, isset bool) *mapInt32Uint8Value {
	return CommandLine.mapInt32Uint8(name, value, usage, required, isset)
}

// mapInt32Uint16Value []mapInt32Uint16Value
type mapInt32Uint16Value map[int32]uint16

func newmapInt32Uint16Value(val mapInt32Uint16Value,
	p *mapInt32Uint16Value) *mapInt32Uint16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt32Uint16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt32Uint16Value) Get() interface{} { return map[int32]uint16(*slc) }

// String join a string from map
func (slc *mapInt32Uint16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt32Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt32Uint16Var(p *mapInt32Uint16Value, name string, value mapInt32Uint16Value, usage string, required, isset bool) {
	f.Var(newmapInt32Uint16Value(value, p), name, usage, required, isset)
}

// mapInt32Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt32Uint16Var(p *mapInt32Uint16Value, name string, value mapInt32Uint16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt32Uint16Value(value, p), name, usage, required, isset)
}

// mapInt32Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt32Uint16(name string, value mapInt32Uint16Value, usage string, required, isset bool) *mapInt32Uint16Value {
	p := new(mapInt32Uint16Value)
	f.mapInt32Uint16Var(p, name, value, usage, required, isset)
	return p
}

// mapInt32Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt32Uint16(name string, value mapInt32Uint16Value, usage string, required, isset bool) *mapInt32Uint16Value {
	return CommandLine.mapInt32Uint16(name, value, usage, required, isset)
}

// mapInt32Uint32Value []mapInt32Uint32Value
type mapInt32Uint32Value map[int32]uint32

func newmapInt32Uint32Value(val mapInt32Uint32Value,
	p *mapInt32Uint32Value) *mapInt32Uint32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt32Uint32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt32Uint32Value) Get() interface{} { return map[int32]uint32(*slc) }

// String join a string from map
func (slc *mapInt32Uint32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt32Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt32Uint32Var(p *mapInt32Uint32Value, name string, value mapInt32Uint32Value, usage string, required, isset bool) {
	f.Var(newmapInt32Uint32Value(value, p), name, usage, required, isset)
}

// mapInt32Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt32Uint32Var(p *mapInt32Uint32Value, name string, value mapInt32Uint32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt32Uint32Value(value, p), name, usage, required, isset)
}

// mapInt32Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt32Uint32(name string, value mapInt32Uint32Value, usage string, required, isset bool) *mapInt32Uint32Value {
	p := new(mapInt32Uint32Value)
	f.mapInt32Uint32Var(p, name, value, usage, required, isset)
	return p
}

// mapInt32Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt32Uint32(name string, value mapInt32Uint32Value, usage string, required, isset bool) *mapInt32Uint32Value {
	return CommandLine.mapInt32Uint32(name, value, usage, required, isset)
}

// mapInt32Uint64Value []mapInt32Uint64Value
type mapInt32Uint64Value map[int32]uint64

func newmapInt32Uint64Value(val mapInt32Uint64Value,
	p *mapInt32Uint64Value) *mapInt32Uint64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt32Uint64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt32Uint64Value) Get() interface{} { return map[int32]uint64(*slc) }

// String join a string from map
func (slc *mapInt32Uint64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt32Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt32Uint64Var(p *mapInt32Uint64Value, name string, value mapInt32Uint64Value, usage string, required, isset bool) {
	f.Var(newmapInt32Uint64Value(value, p), name, usage, required, isset)
}

// mapInt32Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt32Uint64Var(p *mapInt32Uint64Value, name string, value mapInt32Uint64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt32Uint64Value(value, p), name, usage, required, isset)
}

// mapInt32Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt32Uint64(name string, value mapInt32Uint64Value, usage string, required, isset bool) *mapInt32Uint64Value {
	p := new(mapInt32Uint64Value)
	f.mapInt32Uint64Var(p, name, value, usage, required, isset)
	return p
}

// mapInt32Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt32Uint64(name string, value mapInt32Uint64Value, usage string, required, isset bool) *mapInt32Uint64Value {
	return CommandLine.mapInt32Uint64(name, value, usage, required, isset)
}

// mapInt32Float64Value []mapInt32Float64Value
type mapInt32Float64Value map[int32]float64

func newmapInt32Float64Value(val mapInt32Float64Value,
	p *mapInt32Float64Value) *mapInt32Float64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt32Float64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt32Float64Value) Get() interface{} { return map[int32]float64(*slc) }

// String join a string from map
func (slc *mapInt32Float64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt32Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt32Float64Var(p *mapInt32Float64Value, name string, value mapInt32Float64Value, usage string, required, isset bool) {
	f.Var(newmapInt32Float64Value(value, p), name, usage, required, isset)
}

// mapInt32Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt32Float64Var(p *mapInt32Float64Value, name string, value mapInt32Float64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt32Float64Value(value, p), name, usage, required, isset)
}

// mapInt32Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt32Float64(name string, value mapInt32Float64Value, usage string, required, isset bool) *mapInt32Float64Value {
	p := new(mapInt32Float64Value)
	f.mapInt32Float64Var(p, name, value, usage, required, isset)
	return p
}

// mapInt32Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt32Float64(name string, value mapInt32Float64Value, usage string, required, isset bool) *mapInt32Float64Value {
	return CommandLine.mapInt32Float64(name, value, usage, required, isset)
}

// mapInt32Float32Value []mapInt32Float32Value
type mapInt32Float32Value map[int32]float32

func newmapInt32Float32Value(val mapInt32Float32Value,
	p *mapInt32Float32Value) *mapInt32Float32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt32Float32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt32Float32Value) Get() interface{} { return map[int32]float32(*slc) }

// String join a string from map
func (slc *mapInt32Float32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt32Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt32Float32Var(p *mapInt32Float32Value, name string, value mapInt32Float32Value, usage string, required, isset bool) {
	f.Var(newmapInt32Float32Value(value, p), name, usage, required, isset)
}

// mapInt32Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt32Float32Var(p *mapInt32Float32Value, name string, value mapInt32Float32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt32Float32Value(value, p), name, usage, required, isset)
}

// mapInt32Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt32Float32(name string, value mapInt32Float32Value, usage string, required, isset bool) *mapInt32Float32Value {
	p := new(mapInt32Float32Value)
	f.mapInt32Float32Var(p, name, value, usage, required, isset)
	return p
}

// mapInt32Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt32Float32(name string, value mapInt32Float32Value, usage string, required, isset bool) *mapInt32Float32Value {
	return CommandLine.mapInt32Float32(name, value, usage, required, isset)
}

// mapInt32BoolValue []mapInt32BoolValue
type mapInt32BoolValue map[int32]bool

func newmapInt32BoolValue(val mapInt32BoolValue,
	p *mapInt32BoolValue) *mapInt32BoolValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt32BoolValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt32BoolValue) Get() interface{} { return map[int32]bool(*slc) }

// String join a string from map
func (slc *mapInt32BoolValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt32BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt32BoolVar(p *mapInt32BoolValue, name string, value mapInt32BoolValue, usage string, required, isset bool) {
	f.Var(newmapInt32BoolValue(value, p), name, usage, required, isset)
}

// mapInt32BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt32BoolVar(p *mapInt32BoolValue, name string, value mapInt32BoolValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt32BoolValue(value, p), name, usage, required, isset)
}

// mapInt32BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt32Bool(name string, value mapInt32BoolValue, usage string, required, isset bool) *mapInt32BoolValue {
	p := new(mapInt32BoolValue)
	f.mapInt32BoolVar(p, name, value, usage, required, isset)
	return p
}

// mapInt32BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt32Bool(name string, value mapInt32BoolValue, usage string, required, isset bool) *mapInt32BoolValue {
	return CommandLine.mapInt32Bool(name, value, usage, required, isset)
}

// mapInt32StringValue []mapInt32StringValue
type mapInt32StringValue map[int32]string

func newmapInt32StringValue(val mapInt32StringValue,
	p *mapInt32StringValue) *mapInt32StringValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt32StringValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt32StringValue) Get() interface{} { return map[int32]string(*slc) }

// String join a string from map
func (slc *mapInt32StringValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt32StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt32StringVar(p *mapInt32StringValue, name string, value mapInt32StringValue, usage string, required, isset bool) {
	f.Var(newmapInt32StringValue(value, p), name, usage, required, isset)
}

// mapInt32StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt32StringVar(p *mapInt32StringValue, name string, value mapInt32StringValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt32StringValue(value, p), name, usage, required, isset)
}

// mapInt32StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt32String(name string, value mapInt32StringValue, usage string, required, isset bool) *mapInt32StringValue {
	p := new(mapInt32StringValue)
	f.mapInt32StringVar(p, name, value, usage, required, isset)
	return p
}

// mapInt32StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt32String(name string, value mapInt32StringValue, usage string, required, isset bool) *mapInt32StringValue {
	return CommandLine.mapInt32String(name, value, usage, required, isset)
}

// mapInt64DurationValue []mapInt64DurationValue
type mapInt64DurationValue map[int64]time.Duration

func newmapInt64DurationValue(val mapInt64DurationValue,
	p *mapInt64DurationValue) *mapInt64DurationValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt64DurationValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt64DurationValue) Get() interface{} { return map[int64]time.Duration(*slc) }

// String join a string from map
func (slc *mapInt64DurationValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt64DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt64DurationVar(p *mapInt64DurationValue, name string, value mapInt64DurationValue, usage string, required, isset bool) {
	f.Var(newmapInt64DurationValue(value, p), name, usage, required, isset)
}

// mapInt64DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt64DurationVar(p *mapInt64DurationValue, name string, value mapInt64DurationValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt64DurationValue(value, p), name, usage, required, isset)
}

// mapInt64DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt64Duration(name string, value mapInt64DurationValue, usage string, required, isset bool) *mapInt64DurationValue {
	p := new(mapInt64DurationValue)
	f.mapInt64DurationVar(p, name, value, usage, required, isset)
	return p
}

// mapInt64DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt64Duration(name string, value mapInt64DurationValue, usage string, required, isset bool) *mapInt64DurationValue {
	return CommandLine.mapInt64Duration(name, value, usage, required, isset)
}

// mapInt64IntValue []mapInt64IntValue
type mapInt64IntValue map[int64]int

func newmapInt64IntValue(val mapInt64IntValue,
	p *mapInt64IntValue) *mapInt64IntValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt64IntValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt64IntValue) Get() interface{} { return map[int64]int(*slc) }

// String join a string from map
func (slc *mapInt64IntValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt64IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt64IntVar(p *mapInt64IntValue, name string, value mapInt64IntValue, usage string, required, isset bool) {
	f.Var(newmapInt64IntValue(value, p), name, usage, required, isset)
}

// mapInt64IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt64IntVar(p *mapInt64IntValue, name string, value mapInt64IntValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt64IntValue(value, p), name, usage, required, isset)
}

// mapInt64IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt64Int(name string, value mapInt64IntValue, usage string, required, isset bool) *mapInt64IntValue {
	p := new(mapInt64IntValue)
	f.mapInt64IntVar(p, name, value, usage, required, isset)
	return p
}

// mapInt64IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt64Int(name string, value mapInt64IntValue, usage string, required, isset bool) *mapInt64IntValue {
	return CommandLine.mapInt64Int(name, value, usage, required, isset)
}

// mapInt64Int8Value []mapInt64Int8Value
type mapInt64Int8Value map[int64]int8

func newmapInt64Int8Value(val mapInt64Int8Value,
	p *mapInt64Int8Value) *mapInt64Int8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt64Int8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt64Int8Value) Get() interface{} { return map[int64]int8(*slc) }

// String join a string from map
func (slc *mapInt64Int8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt64Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt64Int8Var(p *mapInt64Int8Value, name string, value mapInt64Int8Value, usage string, required, isset bool) {
	f.Var(newmapInt64Int8Value(value, p), name, usage, required, isset)
}

// mapInt64Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt64Int8Var(p *mapInt64Int8Value, name string, value mapInt64Int8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt64Int8Value(value, p), name, usage, required, isset)
}

// mapInt64Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt64Int8(name string, value mapInt64Int8Value, usage string, required, isset bool) *mapInt64Int8Value {
	p := new(mapInt64Int8Value)
	f.mapInt64Int8Var(p, name, value, usage, required, isset)
	return p
}

// mapInt64Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt64Int8(name string, value mapInt64Int8Value, usage string, required, isset bool) *mapInt64Int8Value {
	return CommandLine.mapInt64Int8(name, value, usage, required, isset)
}

// mapInt64Int16Value []mapInt64Int16Value
type mapInt64Int16Value map[int64]int16

func newmapInt64Int16Value(val mapInt64Int16Value,
	p *mapInt64Int16Value) *mapInt64Int16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt64Int16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt64Int16Value) Get() interface{} { return map[int64]int16(*slc) }

// String join a string from map
func (slc *mapInt64Int16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt64Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt64Int16Var(p *mapInt64Int16Value, name string, value mapInt64Int16Value, usage string, required, isset bool) {
	f.Var(newmapInt64Int16Value(value, p), name, usage, required, isset)
}

// mapInt64Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt64Int16Var(p *mapInt64Int16Value, name string, value mapInt64Int16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt64Int16Value(value, p), name, usage, required, isset)
}

// mapInt64Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt64Int16(name string, value mapInt64Int16Value, usage string, required, isset bool) *mapInt64Int16Value {
	p := new(mapInt64Int16Value)
	f.mapInt64Int16Var(p, name, value, usage, required, isset)
	return p
}

// mapInt64Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt64Int16(name string, value mapInt64Int16Value, usage string, required, isset bool) *mapInt64Int16Value {
	return CommandLine.mapInt64Int16(name, value, usage, required, isset)
}

// mapInt64Int32Value []mapInt64Int32Value
type mapInt64Int32Value map[int64]int32

func newmapInt64Int32Value(val mapInt64Int32Value,
	p *mapInt64Int32Value) *mapInt64Int32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt64Int32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt64Int32Value) Get() interface{} { return map[int64]int32(*slc) }

// String join a string from map
func (slc *mapInt64Int32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt64Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt64Int32Var(p *mapInt64Int32Value, name string, value mapInt64Int32Value, usage string, required, isset bool) {
	f.Var(newmapInt64Int32Value(value, p), name, usage, required, isset)
}

// mapInt64Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt64Int32Var(p *mapInt64Int32Value, name string, value mapInt64Int32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt64Int32Value(value, p), name, usage, required, isset)
}

// mapInt64Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt64Int32(name string, value mapInt64Int32Value, usage string, required, isset bool) *mapInt64Int32Value {
	p := new(mapInt64Int32Value)
	f.mapInt64Int32Var(p, name, value, usage, required, isset)
	return p
}

// mapInt64Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt64Int32(name string, value mapInt64Int32Value, usage string, required, isset bool) *mapInt64Int32Value {
	return CommandLine.mapInt64Int32(name, value, usage, required, isset)
}

// mapInt64Int64Value []mapInt64Int64Value
type mapInt64Int64Value map[int64]int64

func newmapInt64Int64Value(val mapInt64Int64Value,
	p *mapInt64Int64Value) *mapInt64Int64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt64Int64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt64Int64Value) Get() interface{} { return map[int64]int64(*slc) }

// String join a string from map
func (slc *mapInt64Int64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt64Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt64Int64Var(p *mapInt64Int64Value, name string, value mapInt64Int64Value, usage string, required, isset bool) {
	f.Var(newmapInt64Int64Value(value, p), name, usage, required, isset)
}

// mapInt64Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt64Int64Var(p *mapInt64Int64Value, name string, value mapInt64Int64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt64Int64Value(value, p), name, usage, required, isset)
}

// mapInt64Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt64Int64(name string, value mapInt64Int64Value, usage string, required, isset bool) *mapInt64Int64Value {
	p := new(mapInt64Int64Value)
	f.mapInt64Int64Var(p, name, value, usage, required, isset)
	return p
}

// mapInt64Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt64Int64(name string, value mapInt64Int64Value, usage string, required, isset bool) *mapInt64Int64Value {
	return CommandLine.mapInt64Int64(name, value, usage, required, isset)
}

// mapInt64UintValue []mapInt64UintValue
type mapInt64UintValue map[int64]uint

func newmapInt64UintValue(val mapInt64UintValue,
	p *mapInt64UintValue) *mapInt64UintValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt64UintValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt64UintValue) Get() interface{} { return map[int64]uint(*slc) }

// String join a string from map
func (slc *mapInt64UintValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt64UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt64UintVar(p *mapInt64UintValue, name string, value mapInt64UintValue, usage string, required, isset bool) {
	f.Var(newmapInt64UintValue(value, p), name, usage, required, isset)
}

// mapInt64UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt64UintVar(p *mapInt64UintValue, name string, value mapInt64UintValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt64UintValue(value, p), name, usage, required, isset)
}

// mapInt64UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt64Uint(name string, value mapInt64UintValue, usage string, required, isset bool) *mapInt64UintValue {
	p := new(mapInt64UintValue)
	f.mapInt64UintVar(p, name, value, usage, required, isset)
	return p
}

// mapInt64UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt64Uint(name string, value mapInt64UintValue, usage string, required, isset bool) *mapInt64UintValue {
	return CommandLine.mapInt64Uint(name, value, usage, required, isset)
}

// mapInt64Uint8Value []mapInt64Uint8Value
type mapInt64Uint8Value map[int64]uint8

func newmapInt64Uint8Value(val mapInt64Uint8Value,
	p *mapInt64Uint8Value) *mapInt64Uint8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt64Uint8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt64Uint8Value) Get() interface{} { return map[int64]uint8(*slc) }

// String join a string from map
func (slc *mapInt64Uint8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt64Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt64Uint8Var(p *mapInt64Uint8Value, name string, value mapInt64Uint8Value, usage string, required, isset bool) {
	f.Var(newmapInt64Uint8Value(value, p), name, usage, required, isset)
}

// mapInt64Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt64Uint8Var(p *mapInt64Uint8Value, name string, value mapInt64Uint8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt64Uint8Value(value, p), name, usage, required, isset)
}

// mapInt64Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt64Uint8(name string, value mapInt64Uint8Value, usage string, required, isset bool) *mapInt64Uint8Value {
	p := new(mapInt64Uint8Value)
	f.mapInt64Uint8Var(p, name, value, usage, required, isset)
	return p
}

// mapInt64Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt64Uint8(name string, value mapInt64Uint8Value, usage string, required, isset bool) *mapInt64Uint8Value {
	return CommandLine.mapInt64Uint8(name, value, usage, required, isset)
}

// mapInt64Uint16Value []mapInt64Uint16Value
type mapInt64Uint16Value map[int64]uint16

func newmapInt64Uint16Value(val mapInt64Uint16Value,
	p *mapInt64Uint16Value) *mapInt64Uint16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt64Uint16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt64Uint16Value) Get() interface{} { return map[int64]uint16(*slc) }

// String join a string from map
func (slc *mapInt64Uint16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt64Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt64Uint16Var(p *mapInt64Uint16Value, name string, value mapInt64Uint16Value, usage string, required, isset bool) {
	f.Var(newmapInt64Uint16Value(value, p), name, usage, required, isset)
}

// mapInt64Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt64Uint16Var(p *mapInt64Uint16Value, name string, value mapInt64Uint16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt64Uint16Value(value, p), name, usage, required, isset)
}

// mapInt64Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt64Uint16(name string, value mapInt64Uint16Value, usage string, required, isset bool) *mapInt64Uint16Value {
	p := new(mapInt64Uint16Value)
	f.mapInt64Uint16Var(p, name, value, usage, required, isset)
	return p
}

// mapInt64Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt64Uint16(name string, value mapInt64Uint16Value, usage string, required, isset bool) *mapInt64Uint16Value {
	return CommandLine.mapInt64Uint16(name, value, usage, required, isset)
}

// mapInt64Uint32Value []mapInt64Uint32Value
type mapInt64Uint32Value map[int64]uint32

func newmapInt64Uint32Value(val mapInt64Uint32Value,
	p *mapInt64Uint32Value) *mapInt64Uint32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt64Uint32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt64Uint32Value) Get() interface{} { return map[int64]uint32(*slc) }

// String join a string from map
func (slc *mapInt64Uint32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt64Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt64Uint32Var(p *mapInt64Uint32Value, name string, value mapInt64Uint32Value, usage string, required, isset bool) {
	f.Var(newmapInt64Uint32Value(value, p), name, usage, required, isset)
}

// mapInt64Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt64Uint32Var(p *mapInt64Uint32Value, name string, value mapInt64Uint32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt64Uint32Value(value, p), name, usage, required, isset)
}

// mapInt64Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt64Uint32(name string, value mapInt64Uint32Value, usage string, required, isset bool) *mapInt64Uint32Value {
	p := new(mapInt64Uint32Value)
	f.mapInt64Uint32Var(p, name, value, usage, required, isset)
	return p
}

// mapInt64Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt64Uint32(name string, value mapInt64Uint32Value, usage string, required, isset bool) *mapInt64Uint32Value {
	return CommandLine.mapInt64Uint32(name, value, usage, required, isset)
}

// mapInt64Uint64Value []mapInt64Uint64Value
type mapInt64Uint64Value map[int64]uint64

func newmapInt64Uint64Value(val mapInt64Uint64Value,
	p *mapInt64Uint64Value) *mapInt64Uint64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt64Uint64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt64Uint64Value) Get() interface{} { return map[int64]uint64(*slc) }

// String join a string from map
func (slc *mapInt64Uint64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt64Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt64Uint64Var(p *mapInt64Uint64Value, name string, value mapInt64Uint64Value, usage string, required, isset bool) {
	f.Var(newmapInt64Uint64Value(value, p), name, usage, required, isset)
}

// mapInt64Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt64Uint64Var(p *mapInt64Uint64Value, name string, value mapInt64Uint64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt64Uint64Value(value, p), name, usage, required, isset)
}

// mapInt64Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt64Uint64(name string, value mapInt64Uint64Value, usage string, required, isset bool) *mapInt64Uint64Value {
	p := new(mapInt64Uint64Value)
	f.mapInt64Uint64Var(p, name, value, usage, required, isset)
	return p
}

// mapInt64Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt64Uint64(name string, value mapInt64Uint64Value, usage string, required, isset bool) *mapInt64Uint64Value {
	return CommandLine.mapInt64Uint64(name, value, usage, required, isset)
}

// mapInt64Float64Value []mapInt64Float64Value
type mapInt64Float64Value map[int64]float64

func newmapInt64Float64Value(val mapInt64Float64Value,
	p *mapInt64Float64Value) *mapInt64Float64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt64Float64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt64Float64Value) Get() interface{} { return map[int64]float64(*slc) }

// String join a string from map
func (slc *mapInt64Float64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt64Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt64Float64Var(p *mapInt64Float64Value, name string, value mapInt64Float64Value, usage string, required, isset bool) {
	f.Var(newmapInt64Float64Value(value, p), name, usage, required, isset)
}

// mapInt64Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt64Float64Var(p *mapInt64Float64Value, name string, value mapInt64Float64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt64Float64Value(value, p), name, usage, required, isset)
}

// mapInt64Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt64Float64(name string, value mapInt64Float64Value, usage string, required, isset bool) *mapInt64Float64Value {
	p := new(mapInt64Float64Value)
	f.mapInt64Float64Var(p, name, value, usage, required, isset)
	return p
}

// mapInt64Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt64Float64(name string, value mapInt64Float64Value, usage string, required, isset bool) *mapInt64Float64Value {
	return CommandLine.mapInt64Float64(name, value, usage, required, isset)
}

// mapInt64Float32Value []mapInt64Float32Value
type mapInt64Float32Value map[int64]float32

func newmapInt64Float32Value(val mapInt64Float32Value,
	p *mapInt64Float32Value) *mapInt64Float32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt64Float32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt64Float32Value) Get() interface{} { return map[int64]float32(*slc) }

// String join a string from map
func (slc *mapInt64Float32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt64Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt64Float32Var(p *mapInt64Float32Value, name string, value mapInt64Float32Value, usage string, required, isset bool) {
	f.Var(newmapInt64Float32Value(value, p), name, usage, required, isset)
}

// mapInt64Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt64Float32Var(p *mapInt64Float32Value, name string, value mapInt64Float32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapInt64Float32Value(value, p), name, usage, required, isset)
}

// mapInt64Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt64Float32(name string, value mapInt64Float32Value, usage string, required, isset bool) *mapInt64Float32Value {
	p := new(mapInt64Float32Value)
	f.mapInt64Float32Var(p, name, value, usage, required, isset)
	return p
}

// mapInt64Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt64Float32(name string, value mapInt64Float32Value, usage string, required, isset bool) *mapInt64Float32Value {
	return CommandLine.mapInt64Float32(name, value, usage, required, isset)
}

// mapInt64BoolValue []mapInt64BoolValue
type mapInt64BoolValue map[int64]bool

func newmapInt64BoolValue(val mapInt64BoolValue,
	p *mapInt64BoolValue) *mapInt64BoolValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt64BoolValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt64BoolValue) Get() interface{} { return map[int64]bool(*slc) }

// String join a string from map
func (slc *mapInt64BoolValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt64BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt64BoolVar(p *mapInt64BoolValue, name string, value mapInt64BoolValue, usage string, required, isset bool) {
	f.Var(newmapInt64BoolValue(value, p), name, usage, required, isset)
}

// mapInt64BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt64BoolVar(p *mapInt64BoolValue, name string, value mapInt64BoolValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt64BoolValue(value, p), name, usage, required, isset)
}

// mapInt64BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt64Bool(name string, value mapInt64BoolValue, usage string, required, isset bool) *mapInt64BoolValue {
	p := new(mapInt64BoolValue)
	f.mapInt64BoolVar(p, name, value, usage, required, isset)
	return p
}

// mapInt64BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt64Bool(name string, value mapInt64BoolValue, usage string, required, isset bool) *mapInt64BoolValue {
	return CommandLine.mapInt64Bool(name, value, usage, required, isset)
}

// mapInt64StringValue []mapInt64StringValue
type mapInt64StringValue map[int64]string

func newmapInt64StringValue(val mapInt64StringValue,
	p *mapInt64StringValue) *mapInt64StringValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapInt64StringValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapInt64StringValue) Get() interface{} { return map[int64]string(*slc) }

// String join a string from map
func (slc *mapInt64StringValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapInt64StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapInt64StringVar(p *mapInt64StringValue, name string, value mapInt64StringValue, usage string, required, isset bool) {
	f.Var(newmapInt64StringValue(value, p), name, usage, required, isset)
}

// mapInt64StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapInt64StringVar(p *mapInt64StringValue, name string, value mapInt64StringValue, usage string, required, isset bool) {
	CommandLine.Var(newmapInt64StringValue(value, p), name, usage, required, isset)
}

// mapInt64StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapInt64String(name string, value mapInt64StringValue, usage string, required, isset bool) *mapInt64StringValue {
	p := new(mapInt64StringValue)
	f.mapInt64StringVar(p, name, value, usage, required, isset)
	return p
}

// mapInt64StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapInt64String(name string, value mapInt64StringValue, usage string, required, isset bool) *mapInt64StringValue {
	return CommandLine.mapInt64String(name, value, usage, required, isset)
}

// mapUintDurationValue []mapUintDurationValue
type mapUintDurationValue map[uint]time.Duration

func newmapUintDurationValue(val mapUintDurationValue,
	p *mapUintDurationValue) *mapUintDurationValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUintDurationValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUintDurationValue) Get() interface{} { return map[uint]time.Duration(*slc) }

// String join a string from map
func (slc *mapUintDurationValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUintDurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUintDurationVar(p *mapUintDurationValue, name string, value mapUintDurationValue, usage string, required, isset bool) {
	f.Var(newmapUintDurationValue(value, p), name, usage, required, isset)
}

// mapUintDurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUintDurationVar(p *mapUintDurationValue, name string, value mapUintDurationValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUintDurationValue(value, p), name, usage, required, isset)
}

// mapUintDurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUintDuration(name string, value mapUintDurationValue, usage string, required, isset bool) *mapUintDurationValue {
	p := new(mapUintDurationValue)
	f.mapUintDurationVar(p, name, value, usage, required, isset)
	return p
}

// mapUintDurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUintDuration(name string, value mapUintDurationValue, usage string, required, isset bool) *mapUintDurationValue {
	return CommandLine.mapUintDuration(name, value, usage, required, isset)
}

// mapUintIntValue []mapUintIntValue
type mapUintIntValue map[uint]int

func newmapUintIntValue(val mapUintIntValue,
	p *mapUintIntValue) *mapUintIntValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUintIntValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUintIntValue) Get() interface{} { return map[uint]int(*slc) }

// String join a string from map
func (slc *mapUintIntValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUintIntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUintIntVar(p *mapUintIntValue, name string, value mapUintIntValue, usage string, required, isset bool) {
	f.Var(newmapUintIntValue(value, p), name, usage, required, isset)
}

// mapUintIntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUintIntVar(p *mapUintIntValue, name string, value mapUintIntValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUintIntValue(value, p), name, usage, required, isset)
}

// mapUintIntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUintInt(name string, value mapUintIntValue, usage string, required, isset bool) *mapUintIntValue {
	p := new(mapUintIntValue)
	f.mapUintIntVar(p, name, value, usage, required, isset)
	return p
}

// mapUintIntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUintInt(name string, value mapUintIntValue, usage string, required, isset bool) *mapUintIntValue {
	return CommandLine.mapUintInt(name, value, usage, required, isset)
}

// mapUintInt8Value []mapUintInt8Value
type mapUintInt8Value map[uint]int8

func newmapUintInt8Value(val mapUintInt8Value,
	p *mapUintInt8Value) *mapUintInt8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUintInt8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUintInt8Value) Get() interface{} { return map[uint]int8(*slc) }

// String join a string from map
func (slc *mapUintInt8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUintInt8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUintInt8Var(p *mapUintInt8Value, name string, value mapUintInt8Value, usage string, required, isset bool) {
	f.Var(newmapUintInt8Value(value, p), name, usage, required, isset)
}

// mapUintInt8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUintInt8Var(p *mapUintInt8Value, name string, value mapUintInt8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUintInt8Value(value, p), name, usage, required, isset)
}

// mapUintInt8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUintInt8(name string, value mapUintInt8Value, usage string, required, isset bool) *mapUintInt8Value {
	p := new(mapUintInt8Value)
	f.mapUintInt8Var(p, name, value, usage, required, isset)
	return p
}

// mapUintInt8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUintInt8(name string, value mapUintInt8Value, usage string, required, isset bool) *mapUintInt8Value {
	return CommandLine.mapUintInt8(name, value, usage, required, isset)
}

// mapUintInt16Value []mapUintInt16Value
type mapUintInt16Value map[uint]int16

func newmapUintInt16Value(val mapUintInt16Value,
	p *mapUintInt16Value) *mapUintInt16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUintInt16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUintInt16Value) Get() interface{} { return map[uint]int16(*slc) }

// String join a string from map
func (slc *mapUintInt16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUintInt16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUintInt16Var(p *mapUintInt16Value, name string, value mapUintInt16Value, usage string, required, isset bool) {
	f.Var(newmapUintInt16Value(value, p), name, usage, required, isset)
}

// mapUintInt16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUintInt16Var(p *mapUintInt16Value, name string, value mapUintInt16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUintInt16Value(value, p), name, usage, required, isset)
}

// mapUintInt16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUintInt16(name string, value mapUintInt16Value, usage string, required, isset bool) *mapUintInt16Value {
	p := new(mapUintInt16Value)
	f.mapUintInt16Var(p, name, value, usage, required, isset)
	return p
}

// mapUintInt16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUintInt16(name string, value mapUintInt16Value, usage string, required, isset bool) *mapUintInt16Value {
	return CommandLine.mapUintInt16(name, value, usage, required, isset)
}

// mapUintInt32Value []mapUintInt32Value
type mapUintInt32Value map[uint]int32

func newmapUintInt32Value(val mapUintInt32Value,
	p *mapUintInt32Value) *mapUintInt32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUintInt32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUintInt32Value) Get() interface{} { return map[uint]int32(*slc) }

// String join a string from map
func (slc *mapUintInt32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUintInt32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUintInt32Var(p *mapUintInt32Value, name string, value mapUintInt32Value, usage string, required, isset bool) {
	f.Var(newmapUintInt32Value(value, p), name, usage, required, isset)
}

// mapUintInt32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUintInt32Var(p *mapUintInt32Value, name string, value mapUintInt32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUintInt32Value(value, p), name, usage, required, isset)
}

// mapUintInt32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUintInt32(name string, value mapUintInt32Value, usage string, required, isset bool) *mapUintInt32Value {
	p := new(mapUintInt32Value)
	f.mapUintInt32Var(p, name, value, usage, required, isset)
	return p
}

// mapUintInt32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUintInt32(name string, value mapUintInt32Value, usage string, required, isset bool) *mapUintInt32Value {
	return CommandLine.mapUintInt32(name, value, usage, required, isset)
}

// mapUintInt64Value []mapUintInt64Value
type mapUintInt64Value map[uint]int64

func newmapUintInt64Value(val mapUintInt64Value,
	p *mapUintInt64Value) *mapUintInt64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUintInt64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUintInt64Value) Get() interface{} { return map[uint]int64(*slc) }

// String join a string from map
func (slc *mapUintInt64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUintInt64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUintInt64Var(p *mapUintInt64Value, name string, value mapUintInt64Value, usage string, required, isset bool) {
	f.Var(newmapUintInt64Value(value, p), name, usage, required, isset)
}

// mapUintInt64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUintInt64Var(p *mapUintInt64Value, name string, value mapUintInt64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUintInt64Value(value, p), name, usage, required, isset)
}

// mapUintInt64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUintInt64(name string, value mapUintInt64Value, usage string, required, isset bool) *mapUintInt64Value {
	p := new(mapUintInt64Value)
	f.mapUintInt64Var(p, name, value, usage, required, isset)
	return p
}

// mapUintInt64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUintInt64(name string, value mapUintInt64Value, usage string, required, isset bool) *mapUintInt64Value {
	return CommandLine.mapUintInt64(name, value, usage, required, isset)
}

// mapUintUintValue []mapUintUintValue
type mapUintUintValue map[uint]uint

func newmapUintUintValue(val mapUintUintValue,
	p *mapUintUintValue) *mapUintUintValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUintUintValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUintUintValue) Get() interface{} { return map[uint]uint(*slc) }

// String join a string from map
func (slc *mapUintUintValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUintUintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUintUintVar(p *mapUintUintValue, name string, value mapUintUintValue, usage string, required, isset bool) {
	f.Var(newmapUintUintValue(value, p), name, usage, required, isset)
}

// mapUintUintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUintUintVar(p *mapUintUintValue, name string, value mapUintUintValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUintUintValue(value, p), name, usage, required, isset)
}

// mapUintUintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUintUint(name string, value mapUintUintValue, usage string, required, isset bool) *mapUintUintValue {
	p := new(mapUintUintValue)
	f.mapUintUintVar(p, name, value, usage, required, isset)
	return p
}

// mapUintUintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUintUint(name string, value mapUintUintValue, usage string, required, isset bool) *mapUintUintValue {
	return CommandLine.mapUintUint(name, value, usage, required, isset)
}

// mapUintUint8Value []mapUintUint8Value
type mapUintUint8Value map[uint]uint8

func newmapUintUint8Value(val mapUintUint8Value,
	p *mapUintUint8Value) *mapUintUint8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUintUint8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUintUint8Value) Get() interface{} { return map[uint]uint8(*slc) }

// String join a string from map
func (slc *mapUintUint8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUintUint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUintUint8Var(p *mapUintUint8Value, name string, value mapUintUint8Value, usage string, required, isset bool) {
	f.Var(newmapUintUint8Value(value, p), name, usage, required, isset)
}

// mapUintUint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUintUint8Var(p *mapUintUint8Value, name string, value mapUintUint8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUintUint8Value(value, p), name, usage, required, isset)
}

// mapUintUint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUintUint8(name string, value mapUintUint8Value, usage string, required, isset bool) *mapUintUint8Value {
	p := new(mapUintUint8Value)
	f.mapUintUint8Var(p, name, value, usage, required, isset)
	return p
}

// mapUintUint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUintUint8(name string, value mapUintUint8Value, usage string, required, isset bool) *mapUintUint8Value {
	return CommandLine.mapUintUint8(name, value, usage, required, isset)
}

// mapUintUint16Value []mapUintUint16Value
type mapUintUint16Value map[uint]uint16

func newmapUintUint16Value(val mapUintUint16Value,
	p *mapUintUint16Value) *mapUintUint16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUintUint16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUintUint16Value) Get() interface{} { return map[uint]uint16(*slc) }

// String join a string from map
func (slc *mapUintUint16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUintUint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUintUint16Var(p *mapUintUint16Value, name string, value mapUintUint16Value, usage string, required, isset bool) {
	f.Var(newmapUintUint16Value(value, p), name, usage, required, isset)
}

// mapUintUint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUintUint16Var(p *mapUintUint16Value, name string, value mapUintUint16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUintUint16Value(value, p), name, usage, required, isset)
}

// mapUintUint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUintUint16(name string, value mapUintUint16Value, usage string, required, isset bool) *mapUintUint16Value {
	p := new(mapUintUint16Value)
	f.mapUintUint16Var(p, name, value, usage, required, isset)
	return p
}

// mapUintUint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUintUint16(name string, value mapUintUint16Value, usage string, required, isset bool) *mapUintUint16Value {
	return CommandLine.mapUintUint16(name, value, usage, required, isset)
}

// mapUintUint32Value []mapUintUint32Value
type mapUintUint32Value map[uint]uint32

func newmapUintUint32Value(val mapUintUint32Value,
	p *mapUintUint32Value) *mapUintUint32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUintUint32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUintUint32Value) Get() interface{} { return map[uint]uint32(*slc) }

// String join a string from map
func (slc *mapUintUint32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUintUint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUintUint32Var(p *mapUintUint32Value, name string, value mapUintUint32Value, usage string, required, isset bool) {
	f.Var(newmapUintUint32Value(value, p), name, usage, required, isset)
}

// mapUintUint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUintUint32Var(p *mapUintUint32Value, name string, value mapUintUint32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUintUint32Value(value, p), name, usage, required, isset)
}

// mapUintUint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUintUint32(name string, value mapUintUint32Value, usage string, required, isset bool) *mapUintUint32Value {
	p := new(mapUintUint32Value)
	f.mapUintUint32Var(p, name, value, usage, required, isset)
	return p
}

// mapUintUint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUintUint32(name string, value mapUintUint32Value, usage string, required, isset bool) *mapUintUint32Value {
	return CommandLine.mapUintUint32(name, value, usage, required, isset)
}

// mapUintUint64Value []mapUintUint64Value
type mapUintUint64Value map[uint]uint64

func newmapUintUint64Value(val mapUintUint64Value,
	p *mapUintUint64Value) *mapUintUint64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUintUint64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUintUint64Value) Get() interface{} { return map[uint]uint64(*slc) }

// String join a string from map
func (slc *mapUintUint64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUintUint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUintUint64Var(p *mapUintUint64Value, name string, value mapUintUint64Value, usage string, required, isset bool) {
	f.Var(newmapUintUint64Value(value, p), name, usage, required, isset)
}

// mapUintUint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUintUint64Var(p *mapUintUint64Value, name string, value mapUintUint64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUintUint64Value(value, p), name, usage, required, isset)
}

// mapUintUint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUintUint64(name string, value mapUintUint64Value, usage string, required, isset bool) *mapUintUint64Value {
	p := new(mapUintUint64Value)
	f.mapUintUint64Var(p, name, value, usage, required, isset)
	return p
}

// mapUintUint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUintUint64(name string, value mapUintUint64Value, usage string, required, isset bool) *mapUintUint64Value {
	return CommandLine.mapUintUint64(name, value, usage, required, isset)
}

// mapUintFloat64Value []mapUintFloat64Value
type mapUintFloat64Value map[uint]float64

func newmapUintFloat64Value(val mapUintFloat64Value,
	p *mapUintFloat64Value) *mapUintFloat64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUintFloat64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUintFloat64Value) Get() interface{} { return map[uint]float64(*slc) }

// String join a string from map
func (slc *mapUintFloat64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUintFloat64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUintFloat64Var(p *mapUintFloat64Value, name string, value mapUintFloat64Value, usage string, required, isset bool) {
	f.Var(newmapUintFloat64Value(value, p), name, usage, required, isset)
}

// mapUintFloat64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUintFloat64Var(p *mapUintFloat64Value, name string, value mapUintFloat64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUintFloat64Value(value, p), name, usage, required, isset)
}

// mapUintFloat64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUintFloat64(name string, value mapUintFloat64Value, usage string, required, isset bool) *mapUintFloat64Value {
	p := new(mapUintFloat64Value)
	f.mapUintFloat64Var(p, name, value, usage, required, isset)
	return p
}

// mapUintFloat64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUintFloat64(name string, value mapUintFloat64Value, usage string, required, isset bool) *mapUintFloat64Value {
	return CommandLine.mapUintFloat64(name, value, usage, required, isset)
}

// mapUintFloat32Value []mapUintFloat32Value
type mapUintFloat32Value map[uint]float32

func newmapUintFloat32Value(val mapUintFloat32Value,
	p *mapUintFloat32Value) *mapUintFloat32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUintFloat32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUintFloat32Value) Get() interface{} { return map[uint]float32(*slc) }

// String join a string from map
func (slc *mapUintFloat32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUintFloat32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUintFloat32Var(p *mapUintFloat32Value, name string, value mapUintFloat32Value, usage string, required, isset bool) {
	f.Var(newmapUintFloat32Value(value, p), name, usage, required, isset)
}

// mapUintFloat32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUintFloat32Var(p *mapUintFloat32Value, name string, value mapUintFloat32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUintFloat32Value(value, p), name, usage, required, isset)
}

// mapUintFloat32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUintFloat32(name string, value mapUintFloat32Value, usage string, required, isset bool) *mapUintFloat32Value {
	p := new(mapUintFloat32Value)
	f.mapUintFloat32Var(p, name, value, usage, required, isset)
	return p
}

// mapUintFloat32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUintFloat32(name string, value mapUintFloat32Value, usage string, required, isset bool) *mapUintFloat32Value {
	return CommandLine.mapUintFloat32(name, value, usage, required, isset)
}

// mapUintBoolValue []mapUintBoolValue
type mapUintBoolValue map[uint]bool

func newmapUintBoolValue(val mapUintBoolValue,
	p *mapUintBoolValue) *mapUintBoolValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUintBoolValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUintBoolValue) Get() interface{} { return map[uint]bool(*slc) }

// String join a string from map
func (slc *mapUintBoolValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUintBoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUintBoolVar(p *mapUintBoolValue, name string, value mapUintBoolValue, usage string, required, isset bool) {
	f.Var(newmapUintBoolValue(value, p), name, usage, required, isset)
}

// mapUintBoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUintBoolVar(p *mapUintBoolValue, name string, value mapUintBoolValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUintBoolValue(value, p), name, usage, required, isset)
}

// mapUintBoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUintBool(name string, value mapUintBoolValue, usage string, required, isset bool) *mapUintBoolValue {
	p := new(mapUintBoolValue)
	f.mapUintBoolVar(p, name, value, usage, required, isset)
	return p
}

// mapUintBoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUintBool(name string, value mapUintBoolValue, usage string, required, isset bool) *mapUintBoolValue {
	return CommandLine.mapUintBool(name, value, usage, required, isset)
}

// mapUintStringValue []mapUintStringValue
type mapUintStringValue map[uint]string

func newmapUintStringValue(val mapUintStringValue,
	p *mapUintStringValue) *mapUintStringValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUintStringValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUintStringValue) Get() interface{} { return map[uint]string(*slc) }

// String join a string from map
func (slc *mapUintStringValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUintStringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUintStringVar(p *mapUintStringValue, name string, value mapUintStringValue, usage string, required, isset bool) {
	f.Var(newmapUintStringValue(value, p), name, usage, required, isset)
}

// mapUintStringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUintStringVar(p *mapUintStringValue, name string, value mapUintStringValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUintStringValue(value, p), name, usage, required, isset)
}

// mapUintStringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUintString(name string, value mapUintStringValue, usage string, required, isset bool) *mapUintStringValue {
	p := new(mapUintStringValue)
	f.mapUintStringVar(p, name, value, usage, required, isset)
	return p
}

// mapUintStringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUintString(name string, value mapUintStringValue, usage string, required, isset bool) *mapUintStringValue {
	return CommandLine.mapUintString(name, value, usage, required, isset)
}

// mapUint8DurationValue []mapUint8DurationValue
type mapUint8DurationValue map[uint8]time.Duration

func newmapUint8DurationValue(val mapUint8DurationValue,
	p *mapUint8DurationValue) *mapUint8DurationValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint8DurationValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint8DurationValue) Get() interface{} { return map[uint8]time.Duration(*slc) }

// String join a string from map
func (slc *mapUint8DurationValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint8DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint8DurationVar(p *mapUint8DurationValue, name string, value mapUint8DurationValue, usage string, required, isset bool) {
	f.Var(newmapUint8DurationValue(value, p), name, usage, required, isset)
}

// mapUint8DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint8DurationVar(p *mapUint8DurationValue, name string, value mapUint8DurationValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint8DurationValue(value, p), name, usage, required, isset)
}

// mapUint8DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint8Duration(name string, value mapUint8DurationValue, usage string, required, isset bool) *mapUint8DurationValue {
	p := new(mapUint8DurationValue)
	f.mapUint8DurationVar(p, name, value, usage, required, isset)
	return p
}

// mapUint8DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint8Duration(name string, value mapUint8DurationValue, usage string, required, isset bool) *mapUint8DurationValue {
	return CommandLine.mapUint8Duration(name, value, usage, required, isset)
}

// mapUint8IntValue []mapUint8IntValue
type mapUint8IntValue map[uint8]int

func newmapUint8IntValue(val mapUint8IntValue,
	p *mapUint8IntValue) *mapUint8IntValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint8IntValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint8IntValue) Get() interface{} { return map[uint8]int(*slc) }

// String join a string from map
func (slc *mapUint8IntValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint8IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint8IntVar(p *mapUint8IntValue, name string, value mapUint8IntValue, usage string, required, isset bool) {
	f.Var(newmapUint8IntValue(value, p), name, usage, required, isset)
}

// mapUint8IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint8IntVar(p *mapUint8IntValue, name string, value mapUint8IntValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint8IntValue(value, p), name, usage, required, isset)
}

// mapUint8IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint8Int(name string, value mapUint8IntValue, usage string, required, isset bool) *mapUint8IntValue {
	p := new(mapUint8IntValue)
	f.mapUint8IntVar(p, name, value, usage, required, isset)
	return p
}

// mapUint8IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint8Int(name string, value mapUint8IntValue, usage string, required, isset bool) *mapUint8IntValue {
	return CommandLine.mapUint8Int(name, value, usage, required, isset)
}

// mapUint8Int8Value []mapUint8Int8Value
type mapUint8Int8Value map[uint8]int8

func newmapUint8Int8Value(val mapUint8Int8Value,
	p *mapUint8Int8Value) *mapUint8Int8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint8Int8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint8Int8Value) Get() interface{} { return map[uint8]int8(*slc) }

// String join a string from map
func (slc *mapUint8Int8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint8Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint8Int8Var(p *mapUint8Int8Value, name string, value mapUint8Int8Value, usage string, required, isset bool) {
	f.Var(newmapUint8Int8Value(value, p), name, usage, required, isset)
}

// mapUint8Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint8Int8Var(p *mapUint8Int8Value, name string, value mapUint8Int8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint8Int8Value(value, p), name, usage, required, isset)
}

// mapUint8Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint8Int8(name string, value mapUint8Int8Value, usage string, required, isset bool) *mapUint8Int8Value {
	p := new(mapUint8Int8Value)
	f.mapUint8Int8Var(p, name, value, usage, required, isset)
	return p
}

// mapUint8Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint8Int8(name string, value mapUint8Int8Value, usage string, required, isset bool) *mapUint8Int8Value {
	return CommandLine.mapUint8Int8(name, value, usage, required, isset)
}

// mapUint8Int16Value []mapUint8Int16Value
type mapUint8Int16Value map[uint8]int16

func newmapUint8Int16Value(val mapUint8Int16Value,
	p *mapUint8Int16Value) *mapUint8Int16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint8Int16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint8Int16Value) Get() interface{} { return map[uint8]int16(*slc) }

// String join a string from map
func (slc *mapUint8Int16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint8Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint8Int16Var(p *mapUint8Int16Value, name string, value mapUint8Int16Value, usage string, required, isset bool) {
	f.Var(newmapUint8Int16Value(value, p), name, usage, required, isset)
}

// mapUint8Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint8Int16Var(p *mapUint8Int16Value, name string, value mapUint8Int16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint8Int16Value(value, p), name, usage, required, isset)
}

// mapUint8Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint8Int16(name string, value mapUint8Int16Value, usage string, required, isset bool) *mapUint8Int16Value {
	p := new(mapUint8Int16Value)
	f.mapUint8Int16Var(p, name, value, usage, required, isset)
	return p
}

// mapUint8Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint8Int16(name string, value mapUint8Int16Value, usage string, required, isset bool) *mapUint8Int16Value {
	return CommandLine.mapUint8Int16(name, value, usage, required, isset)
}

// mapUint8Int32Value []mapUint8Int32Value
type mapUint8Int32Value map[uint8]int32

func newmapUint8Int32Value(val mapUint8Int32Value,
	p *mapUint8Int32Value) *mapUint8Int32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint8Int32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint8Int32Value) Get() interface{} { return map[uint8]int32(*slc) }

// String join a string from map
func (slc *mapUint8Int32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint8Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint8Int32Var(p *mapUint8Int32Value, name string, value mapUint8Int32Value, usage string, required, isset bool) {
	f.Var(newmapUint8Int32Value(value, p), name, usage, required, isset)
}

// mapUint8Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint8Int32Var(p *mapUint8Int32Value, name string, value mapUint8Int32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint8Int32Value(value, p), name, usage, required, isset)
}

// mapUint8Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint8Int32(name string, value mapUint8Int32Value, usage string, required, isset bool) *mapUint8Int32Value {
	p := new(mapUint8Int32Value)
	f.mapUint8Int32Var(p, name, value, usage, required, isset)
	return p
}

// mapUint8Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint8Int32(name string, value mapUint8Int32Value, usage string, required, isset bool) *mapUint8Int32Value {
	return CommandLine.mapUint8Int32(name, value, usage, required, isset)
}

// mapUint8Int64Value []mapUint8Int64Value
type mapUint8Int64Value map[uint8]int64

func newmapUint8Int64Value(val mapUint8Int64Value,
	p *mapUint8Int64Value) *mapUint8Int64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint8Int64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint8Int64Value) Get() interface{} { return map[uint8]int64(*slc) }

// String join a string from map
func (slc *mapUint8Int64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint8Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint8Int64Var(p *mapUint8Int64Value, name string, value mapUint8Int64Value, usage string, required, isset bool) {
	f.Var(newmapUint8Int64Value(value, p), name, usage, required, isset)
}

// mapUint8Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint8Int64Var(p *mapUint8Int64Value, name string, value mapUint8Int64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint8Int64Value(value, p), name, usage, required, isset)
}

// mapUint8Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint8Int64(name string, value mapUint8Int64Value, usage string, required, isset bool) *mapUint8Int64Value {
	p := new(mapUint8Int64Value)
	f.mapUint8Int64Var(p, name, value, usage, required, isset)
	return p
}

// mapUint8Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint8Int64(name string, value mapUint8Int64Value, usage string, required, isset bool) *mapUint8Int64Value {
	return CommandLine.mapUint8Int64(name, value, usage, required, isset)
}

// mapUint8UintValue []mapUint8UintValue
type mapUint8UintValue map[uint8]uint

func newmapUint8UintValue(val mapUint8UintValue,
	p *mapUint8UintValue) *mapUint8UintValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint8UintValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint8UintValue) Get() interface{} { return map[uint8]uint(*slc) }

// String join a string from map
func (slc *mapUint8UintValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint8UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint8UintVar(p *mapUint8UintValue, name string, value mapUint8UintValue, usage string, required, isset bool) {
	f.Var(newmapUint8UintValue(value, p), name, usage, required, isset)
}

// mapUint8UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint8UintVar(p *mapUint8UintValue, name string, value mapUint8UintValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint8UintValue(value, p), name, usage, required, isset)
}

// mapUint8UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint8Uint(name string, value mapUint8UintValue, usage string, required, isset bool) *mapUint8UintValue {
	p := new(mapUint8UintValue)
	f.mapUint8UintVar(p, name, value, usage, required, isset)
	return p
}

// mapUint8UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint8Uint(name string, value mapUint8UintValue, usage string, required, isset bool) *mapUint8UintValue {
	return CommandLine.mapUint8Uint(name, value, usage, required, isset)
}

// mapUint8Uint8Value []mapUint8Uint8Value
type mapUint8Uint8Value map[uint8]uint8

func newmapUint8Uint8Value(val mapUint8Uint8Value,
	p *mapUint8Uint8Value) *mapUint8Uint8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint8Uint8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint8Uint8Value) Get() interface{} { return map[uint8]uint8(*slc) }

// String join a string from map
func (slc *mapUint8Uint8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint8Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint8Uint8Var(p *mapUint8Uint8Value, name string, value mapUint8Uint8Value, usage string, required, isset bool) {
	f.Var(newmapUint8Uint8Value(value, p), name, usage, required, isset)
}

// mapUint8Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint8Uint8Var(p *mapUint8Uint8Value, name string, value mapUint8Uint8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint8Uint8Value(value, p), name, usage, required, isset)
}

// mapUint8Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint8Uint8(name string, value mapUint8Uint8Value, usage string, required, isset bool) *mapUint8Uint8Value {
	p := new(mapUint8Uint8Value)
	f.mapUint8Uint8Var(p, name, value, usage, required, isset)
	return p
}

// mapUint8Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint8Uint8(name string, value mapUint8Uint8Value, usage string, required, isset bool) *mapUint8Uint8Value {
	return CommandLine.mapUint8Uint8(name, value, usage, required, isset)
}

// mapUint8Uint16Value []mapUint8Uint16Value
type mapUint8Uint16Value map[uint8]uint16

func newmapUint8Uint16Value(val mapUint8Uint16Value,
	p *mapUint8Uint16Value) *mapUint8Uint16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint8Uint16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint8Uint16Value) Get() interface{} { return map[uint8]uint16(*slc) }

// String join a string from map
func (slc *mapUint8Uint16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint8Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint8Uint16Var(p *mapUint8Uint16Value, name string, value mapUint8Uint16Value, usage string, required, isset bool) {
	f.Var(newmapUint8Uint16Value(value, p), name, usage, required, isset)
}

// mapUint8Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint8Uint16Var(p *mapUint8Uint16Value, name string, value mapUint8Uint16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint8Uint16Value(value, p), name, usage, required, isset)
}

// mapUint8Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint8Uint16(name string, value mapUint8Uint16Value, usage string, required, isset bool) *mapUint8Uint16Value {
	p := new(mapUint8Uint16Value)
	f.mapUint8Uint16Var(p, name, value, usage, required, isset)
	return p
}

// mapUint8Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint8Uint16(name string, value mapUint8Uint16Value, usage string, required, isset bool) *mapUint8Uint16Value {
	return CommandLine.mapUint8Uint16(name, value, usage, required, isset)
}

// mapUint8Uint32Value []mapUint8Uint32Value
type mapUint8Uint32Value map[uint8]uint32

func newmapUint8Uint32Value(val mapUint8Uint32Value,
	p *mapUint8Uint32Value) *mapUint8Uint32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint8Uint32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint8Uint32Value) Get() interface{} { return map[uint8]uint32(*slc) }

// String join a string from map
func (slc *mapUint8Uint32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint8Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint8Uint32Var(p *mapUint8Uint32Value, name string, value mapUint8Uint32Value, usage string, required, isset bool) {
	f.Var(newmapUint8Uint32Value(value, p), name, usage, required, isset)
}

// mapUint8Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint8Uint32Var(p *mapUint8Uint32Value, name string, value mapUint8Uint32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint8Uint32Value(value, p), name, usage, required, isset)
}

// mapUint8Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint8Uint32(name string, value mapUint8Uint32Value, usage string, required, isset bool) *mapUint8Uint32Value {
	p := new(mapUint8Uint32Value)
	f.mapUint8Uint32Var(p, name, value, usage, required, isset)
	return p
}

// mapUint8Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint8Uint32(name string, value mapUint8Uint32Value, usage string, required, isset bool) *mapUint8Uint32Value {
	return CommandLine.mapUint8Uint32(name, value, usage, required, isset)
}

// mapUint8Uint64Value []mapUint8Uint64Value
type mapUint8Uint64Value map[uint8]uint64

func newmapUint8Uint64Value(val mapUint8Uint64Value,
	p *mapUint8Uint64Value) *mapUint8Uint64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint8Uint64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint8Uint64Value) Get() interface{} { return map[uint8]uint64(*slc) }

// String join a string from map
func (slc *mapUint8Uint64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint8Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint8Uint64Var(p *mapUint8Uint64Value, name string, value mapUint8Uint64Value, usage string, required, isset bool) {
	f.Var(newmapUint8Uint64Value(value, p), name, usage, required, isset)
}

// mapUint8Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint8Uint64Var(p *mapUint8Uint64Value, name string, value mapUint8Uint64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint8Uint64Value(value, p), name, usage, required, isset)
}

// mapUint8Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint8Uint64(name string, value mapUint8Uint64Value, usage string, required, isset bool) *mapUint8Uint64Value {
	p := new(mapUint8Uint64Value)
	f.mapUint8Uint64Var(p, name, value, usage, required, isset)
	return p
}

// mapUint8Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint8Uint64(name string, value mapUint8Uint64Value, usage string, required, isset bool) *mapUint8Uint64Value {
	return CommandLine.mapUint8Uint64(name, value, usage, required, isset)
}

// mapUint8Float64Value []mapUint8Float64Value
type mapUint8Float64Value map[uint8]float64

func newmapUint8Float64Value(val mapUint8Float64Value,
	p *mapUint8Float64Value) *mapUint8Float64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint8Float64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint8Float64Value) Get() interface{} { return map[uint8]float64(*slc) }

// String join a string from map
func (slc *mapUint8Float64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint8Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint8Float64Var(p *mapUint8Float64Value, name string, value mapUint8Float64Value, usage string, required, isset bool) {
	f.Var(newmapUint8Float64Value(value, p), name, usage, required, isset)
}

// mapUint8Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint8Float64Var(p *mapUint8Float64Value, name string, value mapUint8Float64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint8Float64Value(value, p), name, usage, required, isset)
}

// mapUint8Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint8Float64(name string, value mapUint8Float64Value, usage string, required, isset bool) *mapUint8Float64Value {
	p := new(mapUint8Float64Value)
	f.mapUint8Float64Var(p, name, value, usage, required, isset)
	return p
}

// mapUint8Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint8Float64(name string, value mapUint8Float64Value, usage string, required, isset bool) *mapUint8Float64Value {
	return CommandLine.mapUint8Float64(name, value, usage, required, isset)
}

// mapUint8Float32Value []mapUint8Float32Value
type mapUint8Float32Value map[uint8]float32

func newmapUint8Float32Value(val mapUint8Float32Value,
	p *mapUint8Float32Value) *mapUint8Float32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint8Float32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint8Float32Value) Get() interface{} { return map[uint8]float32(*slc) }

// String join a string from map
func (slc *mapUint8Float32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint8Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint8Float32Var(p *mapUint8Float32Value, name string, value mapUint8Float32Value, usage string, required, isset bool) {
	f.Var(newmapUint8Float32Value(value, p), name, usage, required, isset)
}

// mapUint8Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint8Float32Var(p *mapUint8Float32Value, name string, value mapUint8Float32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint8Float32Value(value, p), name, usage, required, isset)
}

// mapUint8Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint8Float32(name string, value mapUint8Float32Value, usage string, required, isset bool) *mapUint8Float32Value {
	p := new(mapUint8Float32Value)
	f.mapUint8Float32Var(p, name, value, usage, required, isset)
	return p
}

// mapUint8Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint8Float32(name string, value mapUint8Float32Value, usage string, required, isset bool) *mapUint8Float32Value {
	return CommandLine.mapUint8Float32(name, value, usage, required, isset)
}

// mapUint8BoolValue []mapUint8BoolValue
type mapUint8BoolValue map[uint8]bool

func newmapUint8BoolValue(val mapUint8BoolValue,
	p *mapUint8BoolValue) *mapUint8BoolValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint8BoolValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint8BoolValue) Get() interface{} { return map[uint8]bool(*slc) }

// String join a string from map
func (slc *mapUint8BoolValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint8BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint8BoolVar(p *mapUint8BoolValue, name string, value mapUint8BoolValue, usage string, required, isset bool) {
	f.Var(newmapUint8BoolValue(value, p), name, usage, required, isset)
}

// mapUint8BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint8BoolVar(p *mapUint8BoolValue, name string, value mapUint8BoolValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint8BoolValue(value, p), name, usage, required, isset)
}

// mapUint8BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint8Bool(name string, value mapUint8BoolValue, usage string, required, isset bool) *mapUint8BoolValue {
	p := new(mapUint8BoolValue)
	f.mapUint8BoolVar(p, name, value, usage, required, isset)
	return p
}

// mapUint8BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint8Bool(name string, value mapUint8BoolValue, usage string, required, isset bool) *mapUint8BoolValue {
	return CommandLine.mapUint8Bool(name, value, usage, required, isset)
}

// mapUint8StringValue []mapUint8StringValue
type mapUint8StringValue map[uint8]string

func newmapUint8StringValue(val mapUint8StringValue,
	p *mapUint8StringValue) *mapUint8StringValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint8StringValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint8StringValue) Get() interface{} { return map[uint8]string(*slc) }

// String join a string from map
func (slc *mapUint8StringValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint8StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint8StringVar(p *mapUint8StringValue, name string, value mapUint8StringValue, usage string, required, isset bool) {
	f.Var(newmapUint8StringValue(value, p), name, usage, required, isset)
}

// mapUint8StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint8StringVar(p *mapUint8StringValue, name string, value mapUint8StringValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint8StringValue(value, p), name, usage, required, isset)
}

// mapUint8StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint8String(name string, value mapUint8StringValue, usage string, required, isset bool) *mapUint8StringValue {
	p := new(mapUint8StringValue)
	f.mapUint8StringVar(p, name, value, usage, required, isset)
	return p
}

// mapUint8StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint8String(name string, value mapUint8StringValue, usage string, required, isset bool) *mapUint8StringValue {
	return CommandLine.mapUint8String(name, value, usage, required, isset)
}

// mapUint16DurationValue []mapUint16DurationValue
type mapUint16DurationValue map[uint16]time.Duration

func newmapUint16DurationValue(val mapUint16DurationValue,
	p *mapUint16DurationValue) *mapUint16DurationValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint16DurationValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint16DurationValue) Get() interface{} { return map[uint16]time.Duration(*slc) }

// String join a string from map
func (slc *mapUint16DurationValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint16DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint16DurationVar(p *mapUint16DurationValue, name string, value mapUint16DurationValue, usage string, required, isset bool) {
	f.Var(newmapUint16DurationValue(value, p), name, usage, required, isset)
}

// mapUint16DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint16DurationVar(p *mapUint16DurationValue, name string, value mapUint16DurationValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint16DurationValue(value, p), name, usage, required, isset)
}

// mapUint16DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint16Duration(name string, value mapUint16DurationValue, usage string, required, isset bool) *mapUint16DurationValue {
	p := new(mapUint16DurationValue)
	f.mapUint16DurationVar(p, name, value, usage, required, isset)
	return p
}

// mapUint16DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint16Duration(name string, value mapUint16DurationValue, usage string, required, isset bool) *mapUint16DurationValue {
	return CommandLine.mapUint16Duration(name, value, usage, required, isset)
}

// mapUint16IntValue []mapUint16IntValue
type mapUint16IntValue map[uint16]int

func newmapUint16IntValue(val mapUint16IntValue,
	p *mapUint16IntValue) *mapUint16IntValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint16IntValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint16IntValue) Get() interface{} { return map[uint16]int(*slc) }

// String join a string from map
func (slc *mapUint16IntValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint16IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint16IntVar(p *mapUint16IntValue, name string, value mapUint16IntValue, usage string, required, isset bool) {
	f.Var(newmapUint16IntValue(value, p), name, usage, required, isset)
}

// mapUint16IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint16IntVar(p *mapUint16IntValue, name string, value mapUint16IntValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint16IntValue(value, p), name, usage, required, isset)
}

// mapUint16IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint16Int(name string, value mapUint16IntValue, usage string, required, isset bool) *mapUint16IntValue {
	p := new(mapUint16IntValue)
	f.mapUint16IntVar(p, name, value, usage, required, isset)
	return p
}

// mapUint16IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint16Int(name string, value mapUint16IntValue, usage string, required, isset bool) *mapUint16IntValue {
	return CommandLine.mapUint16Int(name, value, usage, required, isset)
}

// mapUint16Int8Value []mapUint16Int8Value
type mapUint16Int8Value map[uint16]int8

func newmapUint16Int8Value(val mapUint16Int8Value,
	p *mapUint16Int8Value) *mapUint16Int8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint16Int8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint16Int8Value) Get() interface{} { return map[uint16]int8(*slc) }

// String join a string from map
func (slc *mapUint16Int8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint16Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint16Int8Var(p *mapUint16Int8Value, name string, value mapUint16Int8Value, usage string, required, isset bool) {
	f.Var(newmapUint16Int8Value(value, p), name, usage, required, isset)
}

// mapUint16Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint16Int8Var(p *mapUint16Int8Value, name string, value mapUint16Int8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint16Int8Value(value, p), name, usage, required, isset)
}

// mapUint16Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint16Int8(name string, value mapUint16Int8Value, usage string, required, isset bool) *mapUint16Int8Value {
	p := new(mapUint16Int8Value)
	f.mapUint16Int8Var(p, name, value, usage, required, isset)
	return p
}

// mapUint16Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint16Int8(name string, value mapUint16Int8Value, usage string, required, isset bool) *mapUint16Int8Value {
	return CommandLine.mapUint16Int8(name, value, usage, required, isset)
}

// mapUint16Int16Value []mapUint16Int16Value
type mapUint16Int16Value map[uint16]int16

func newmapUint16Int16Value(val mapUint16Int16Value,
	p *mapUint16Int16Value) *mapUint16Int16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint16Int16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint16Int16Value) Get() interface{} { return map[uint16]int16(*slc) }

// String join a string from map
func (slc *mapUint16Int16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint16Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint16Int16Var(p *mapUint16Int16Value, name string, value mapUint16Int16Value, usage string, required, isset bool) {
	f.Var(newmapUint16Int16Value(value, p), name, usage, required, isset)
}

// mapUint16Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint16Int16Var(p *mapUint16Int16Value, name string, value mapUint16Int16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint16Int16Value(value, p), name, usage, required, isset)
}

// mapUint16Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint16Int16(name string, value mapUint16Int16Value, usage string, required, isset bool) *mapUint16Int16Value {
	p := new(mapUint16Int16Value)
	f.mapUint16Int16Var(p, name, value, usage, required, isset)
	return p
}

// mapUint16Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint16Int16(name string, value mapUint16Int16Value, usage string, required, isset bool) *mapUint16Int16Value {
	return CommandLine.mapUint16Int16(name, value, usage, required, isset)
}

// mapUint16Int32Value []mapUint16Int32Value
type mapUint16Int32Value map[uint16]int32

func newmapUint16Int32Value(val mapUint16Int32Value,
	p *mapUint16Int32Value) *mapUint16Int32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint16Int32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint16Int32Value) Get() interface{} { return map[uint16]int32(*slc) }

// String join a string from map
func (slc *mapUint16Int32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint16Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint16Int32Var(p *mapUint16Int32Value, name string, value mapUint16Int32Value, usage string, required, isset bool) {
	f.Var(newmapUint16Int32Value(value, p), name, usage, required, isset)
}

// mapUint16Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint16Int32Var(p *mapUint16Int32Value, name string, value mapUint16Int32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint16Int32Value(value, p), name, usage, required, isset)
}

// mapUint16Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint16Int32(name string, value mapUint16Int32Value, usage string, required, isset bool) *mapUint16Int32Value {
	p := new(mapUint16Int32Value)
	f.mapUint16Int32Var(p, name, value, usage, required, isset)
	return p
}

// mapUint16Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint16Int32(name string, value mapUint16Int32Value, usage string, required, isset bool) *mapUint16Int32Value {
	return CommandLine.mapUint16Int32(name, value, usage, required, isset)
}

// mapUint16Int64Value []mapUint16Int64Value
type mapUint16Int64Value map[uint16]int64

func newmapUint16Int64Value(val mapUint16Int64Value,
	p *mapUint16Int64Value) *mapUint16Int64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint16Int64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint16Int64Value) Get() interface{} { return map[uint16]int64(*slc) }

// String join a string from map
func (slc *mapUint16Int64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint16Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint16Int64Var(p *mapUint16Int64Value, name string, value mapUint16Int64Value, usage string, required, isset bool) {
	f.Var(newmapUint16Int64Value(value, p), name, usage, required, isset)
}

// mapUint16Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint16Int64Var(p *mapUint16Int64Value, name string, value mapUint16Int64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint16Int64Value(value, p), name, usage, required, isset)
}

// mapUint16Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint16Int64(name string, value mapUint16Int64Value, usage string, required, isset bool) *mapUint16Int64Value {
	p := new(mapUint16Int64Value)
	f.mapUint16Int64Var(p, name, value, usage, required, isset)
	return p
}

// mapUint16Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint16Int64(name string, value mapUint16Int64Value, usage string, required, isset bool) *mapUint16Int64Value {
	return CommandLine.mapUint16Int64(name, value, usage, required, isset)
}

// mapUint16UintValue []mapUint16UintValue
type mapUint16UintValue map[uint16]uint

func newmapUint16UintValue(val mapUint16UintValue,
	p *mapUint16UintValue) *mapUint16UintValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint16UintValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint16UintValue) Get() interface{} { return map[uint16]uint(*slc) }

// String join a string from map
func (slc *mapUint16UintValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint16UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint16UintVar(p *mapUint16UintValue, name string, value mapUint16UintValue, usage string, required, isset bool) {
	f.Var(newmapUint16UintValue(value, p), name, usage, required, isset)
}

// mapUint16UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint16UintVar(p *mapUint16UintValue, name string, value mapUint16UintValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint16UintValue(value, p), name, usage, required, isset)
}

// mapUint16UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint16Uint(name string, value mapUint16UintValue, usage string, required, isset bool) *mapUint16UintValue {
	p := new(mapUint16UintValue)
	f.mapUint16UintVar(p, name, value, usage, required, isset)
	return p
}

// mapUint16UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint16Uint(name string, value mapUint16UintValue, usage string, required, isset bool) *mapUint16UintValue {
	return CommandLine.mapUint16Uint(name, value, usage, required, isset)
}

// mapUint16Uint8Value []mapUint16Uint8Value
type mapUint16Uint8Value map[uint16]uint8

func newmapUint16Uint8Value(val mapUint16Uint8Value,
	p *mapUint16Uint8Value) *mapUint16Uint8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint16Uint8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint16Uint8Value) Get() interface{} { return map[uint16]uint8(*slc) }

// String join a string from map
func (slc *mapUint16Uint8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint16Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint16Uint8Var(p *mapUint16Uint8Value, name string, value mapUint16Uint8Value, usage string, required, isset bool) {
	f.Var(newmapUint16Uint8Value(value, p), name, usage, required, isset)
}

// mapUint16Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint16Uint8Var(p *mapUint16Uint8Value, name string, value mapUint16Uint8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint16Uint8Value(value, p), name, usage, required, isset)
}

// mapUint16Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint16Uint8(name string, value mapUint16Uint8Value, usage string, required, isset bool) *mapUint16Uint8Value {
	p := new(mapUint16Uint8Value)
	f.mapUint16Uint8Var(p, name, value, usage, required, isset)
	return p
}

// mapUint16Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint16Uint8(name string, value mapUint16Uint8Value, usage string, required, isset bool) *mapUint16Uint8Value {
	return CommandLine.mapUint16Uint8(name, value, usage, required, isset)
}

// mapUint16Uint16Value []mapUint16Uint16Value
type mapUint16Uint16Value map[uint16]uint16

func newmapUint16Uint16Value(val mapUint16Uint16Value,
	p *mapUint16Uint16Value) *mapUint16Uint16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint16Uint16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint16Uint16Value) Get() interface{} { return map[uint16]uint16(*slc) }

// String join a string from map
func (slc *mapUint16Uint16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint16Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint16Uint16Var(p *mapUint16Uint16Value, name string, value mapUint16Uint16Value, usage string, required, isset bool) {
	f.Var(newmapUint16Uint16Value(value, p), name, usage, required, isset)
}

// mapUint16Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint16Uint16Var(p *mapUint16Uint16Value, name string, value mapUint16Uint16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint16Uint16Value(value, p), name, usage, required, isset)
}

// mapUint16Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint16Uint16(name string, value mapUint16Uint16Value, usage string, required, isset bool) *mapUint16Uint16Value {
	p := new(mapUint16Uint16Value)
	f.mapUint16Uint16Var(p, name, value, usage, required, isset)
	return p
}

// mapUint16Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint16Uint16(name string, value mapUint16Uint16Value, usage string, required, isset bool) *mapUint16Uint16Value {
	return CommandLine.mapUint16Uint16(name, value, usage, required, isset)
}

// mapUint16Uint32Value []mapUint16Uint32Value
type mapUint16Uint32Value map[uint16]uint32

func newmapUint16Uint32Value(val mapUint16Uint32Value,
	p *mapUint16Uint32Value) *mapUint16Uint32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint16Uint32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint16Uint32Value) Get() interface{} { return map[uint16]uint32(*slc) }

// String join a string from map
func (slc *mapUint16Uint32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint16Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint16Uint32Var(p *mapUint16Uint32Value, name string, value mapUint16Uint32Value, usage string, required, isset bool) {
	f.Var(newmapUint16Uint32Value(value, p), name, usage, required, isset)
}

// mapUint16Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint16Uint32Var(p *mapUint16Uint32Value, name string, value mapUint16Uint32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint16Uint32Value(value, p), name, usage, required, isset)
}

// mapUint16Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint16Uint32(name string, value mapUint16Uint32Value, usage string, required, isset bool) *mapUint16Uint32Value {
	p := new(mapUint16Uint32Value)
	f.mapUint16Uint32Var(p, name, value, usage, required, isset)
	return p
}

// mapUint16Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint16Uint32(name string, value mapUint16Uint32Value, usage string, required, isset bool) *mapUint16Uint32Value {
	return CommandLine.mapUint16Uint32(name, value, usage, required, isset)
}

// mapUint16Uint64Value []mapUint16Uint64Value
type mapUint16Uint64Value map[uint16]uint64

func newmapUint16Uint64Value(val mapUint16Uint64Value,
	p *mapUint16Uint64Value) *mapUint16Uint64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint16Uint64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint16Uint64Value) Get() interface{} { return map[uint16]uint64(*slc) }

// String join a string from map
func (slc *mapUint16Uint64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint16Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint16Uint64Var(p *mapUint16Uint64Value, name string, value mapUint16Uint64Value, usage string, required, isset bool) {
	f.Var(newmapUint16Uint64Value(value, p), name, usage, required, isset)
}

// mapUint16Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint16Uint64Var(p *mapUint16Uint64Value, name string, value mapUint16Uint64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint16Uint64Value(value, p), name, usage, required, isset)
}

// mapUint16Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint16Uint64(name string, value mapUint16Uint64Value, usage string, required, isset bool) *mapUint16Uint64Value {
	p := new(mapUint16Uint64Value)
	f.mapUint16Uint64Var(p, name, value, usage, required, isset)
	return p
}

// mapUint16Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint16Uint64(name string, value mapUint16Uint64Value, usage string, required, isset bool) *mapUint16Uint64Value {
	return CommandLine.mapUint16Uint64(name, value, usage, required, isset)
}

// mapUint16Float64Value []mapUint16Float64Value
type mapUint16Float64Value map[uint16]float64

func newmapUint16Float64Value(val mapUint16Float64Value,
	p *mapUint16Float64Value) *mapUint16Float64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint16Float64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint16Float64Value) Get() interface{} { return map[uint16]float64(*slc) }

// String join a string from map
func (slc *mapUint16Float64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint16Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint16Float64Var(p *mapUint16Float64Value, name string, value mapUint16Float64Value, usage string, required, isset bool) {
	f.Var(newmapUint16Float64Value(value, p), name, usage, required, isset)
}

// mapUint16Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint16Float64Var(p *mapUint16Float64Value, name string, value mapUint16Float64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint16Float64Value(value, p), name, usage, required, isset)
}

// mapUint16Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint16Float64(name string, value mapUint16Float64Value, usage string, required, isset bool) *mapUint16Float64Value {
	p := new(mapUint16Float64Value)
	f.mapUint16Float64Var(p, name, value, usage, required, isset)
	return p
}

// mapUint16Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint16Float64(name string, value mapUint16Float64Value, usage string, required, isset bool) *mapUint16Float64Value {
	return CommandLine.mapUint16Float64(name, value, usage, required, isset)
}

// mapUint16Float32Value []mapUint16Float32Value
type mapUint16Float32Value map[uint16]float32

func newmapUint16Float32Value(val mapUint16Float32Value,
	p *mapUint16Float32Value) *mapUint16Float32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint16Float32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint16Float32Value) Get() interface{} { return map[uint16]float32(*slc) }

// String join a string from map
func (slc *mapUint16Float32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint16Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint16Float32Var(p *mapUint16Float32Value, name string, value mapUint16Float32Value, usage string, required, isset bool) {
	f.Var(newmapUint16Float32Value(value, p), name, usage, required, isset)
}

// mapUint16Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint16Float32Var(p *mapUint16Float32Value, name string, value mapUint16Float32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint16Float32Value(value, p), name, usage, required, isset)
}

// mapUint16Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint16Float32(name string, value mapUint16Float32Value, usage string, required, isset bool) *mapUint16Float32Value {
	p := new(mapUint16Float32Value)
	f.mapUint16Float32Var(p, name, value, usage, required, isset)
	return p
}

// mapUint16Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint16Float32(name string, value mapUint16Float32Value, usage string, required, isset bool) *mapUint16Float32Value {
	return CommandLine.mapUint16Float32(name, value, usage, required, isset)
}

// mapUint16BoolValue []mapUint16BoolValue
type mapUint16BoolValue map[uint16]bool

func newmapUint16BoolValue(val mapUint16BoolValue,
	p *mapUint16BoolValue) *mapUint16BoolValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint16BoolValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint16BoolValue) Get() interface{} { return map[uint16]bool(*slc) }

// String join a string from map
func (slc *mapUint16BoolValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint16BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint16BoolVar(p *mapUint16BoolValue, name string, value mapUint16BoolValue, usage string, required, isset bool) {
	f.Var(newmapUint16BoolValue(value, p), name, usage, required, isset)
}

// mapUint16BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint16BoolVar(p *mapUint16BoolValue, name string, value mapUint16BoolValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint16BoolValue(value, p), name, usage, required, isset)
}

// mapUint16BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint16Bool(name string, value mapUint16BoolValue, usage string, required, isset bool) *mapUint16BoolValue {
	p := new(mapUint16BoolValue)
	f.mapUint16BoolVar(p, name, value, usage, required, isset)
	return p
}

// mapUint16BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint16Bool(name string, value mapUint16BoolValue, usage string, required, isset bool) *mapUint16BoolValue {
	return CommandLine.mapUint16Bool(name, value, usage, required, isset)
}

// mapUint16StringValue []mapUint16StringValue
type mapUint16StringValue map[uint16]string

func newmapUint16StringValue(val mapUint16StringValue,
	p *mapUint16StringValue) *mapUint16StringValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint16StringValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint16StringValue) Get() interface{} { return map[uint16]string(*slc) }

// String join a string from map
func (slc *mapUint16StringValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint16StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint16StringVar(p *mapUint16StringValue, name string, value mapUint16StringValue, usage string, required, isset bool) {
	f.Var(newmapUint16StringValue(value, p), name, usage, required, isset)
}

// mapUint16StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint16StringVar(p *mapUint16StringValue, name string, value mapUint16StringValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint16StringValue(value, p), name, usage, required, isset)
}

// mapUint16StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint16String(name string, value mapUint16StringValue, usage string, required, isset bool) *mapUint16StringValue {
	p := new(mapUint16StringValue)
	f.mapUint16StringVar(p, name, value, usage, required, isset)
	return p
}

// mapUint16StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint16String(name string, value mapUint16StringValue, usage string, required, isset bool) *mapUint16StringValue {
	return CommandLine.mapUint16String(name, value, usage, required, isset)
}

// mapUint32DurationValue []mapUint32DurationValue
type mapUint32DurationValue map[uint32]time.Duration

func newmapUint32DurationValue(val mapUint32DurationValue,
	p *mapUint32DurationValue) *mapUint32DurationValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint32DurationValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint32DurationValue) Get() interface{} { return map[uint32]time.Duration(*slc) }

// String join a string from map
func (slc *mapUint32DurationValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint32DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint32DurationVar(p *mapUint32DurationValue, name string, value mapUint32DurationValue, usage string, required, isset bool) {
	f.Var(newmapUint32DurationValue(value, p), name, usage, required, isset)
}

// mapUint32DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint32DurationVar(p *mapUint32DurationValue, name string, value mapUint32DurationValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint32DurationValue(value, p), name, usage, required, isset)
}

// mapUint32DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint32Duration(name string, value mapUint32DurationValue, usage string, required, isset bool) *mapUint32DurationValue {
	p := new(mapUint32DurationValue)
	f.mapUint32DurationVar(p, name, value, usage, required, isset)
	return p
}

// mapUint32DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint32Duration(name string, value mapUint32DurationValue, usage string, required, isset bool) *mapUint32DurationValue {
	return CommandLine.mapUint32Duration(name, value, usage, required, isset)
}

// mapUint32IntValue []mapUint32IntValue
type mapUint32IntValue map[uint32]int

func newmapUint32IntValue(val mapUint32IntValue,
	p *mapUint32IntValue) *mapUint32IntValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint32IntValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint32IntValue) Get() interface{} { return map[uint32]int(*slc) }

// String join a string from map
func (slc *mapUint32IntValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint32IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint32IntVar(p *mapUint32IntValue, name string, value mapUint32IntValue, usage string, required, isset bool) {
	f.Var(newmapUint32IntValue(value, p), name, usage, required, isset)
}

// mapUint32IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint32IntVar(p *mapUint32IntValue, name string, value mapUint32IntValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint32IntValue(value, p), name, usage, required, isset)
}

// mapUint32IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint32Int(name string, value mapUint32IntValue, usage string, required, isset bool) *mapUint32IntValue {
	p := new(mapUint32IntValue)
	f.mapUint32IntVar(p, name, value, usage, required, isset)
	return p
}

// mapUint32IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint32Int(name string, value mapUint32IntValue, usage string, required, isset bool) *mapUint32IntValue {
	return CommandLine.mapUint32Int(name, value, usage, required, isset)
}

// mapUint32Int8Value []mapUint32Int8Value
type mapUint32Int8Value map[uint32]int8

func newmapUint32Int8Value(val mapUint32Int8Value,
	p *mapUint32Int8Value) *mapUint32Int8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint32Int8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint32Int8Value) Get() interface{} { return map[uint32]int8(*slc) }

// String join a string from map
func (slc *mapUint32Int8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint32Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint32Int8Var(p *mapUint32Int8Value, name string, value mapUint32Int8Value, usage string, required, isset bool) {
	f.Var(newmapUint32Int8Value(value, p), name, usage, required, isset)
}

// mapUint32Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint32Int8Var(p *mapUint32Int8Value, name string, value mapUint32Int8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint32Int8Value(value, p), name, usage, required, isset)
}

// mapUint32Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint32Int8(name string, value mapUint32Int8Value, usage string, required, isset bool) *mapUint32Int8Value {
	p := new(mapUint32Int8Value)
	f.mapUint32Int8Var(p, name, value, usage, required, isset)
	return p
}

// mapUint32Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint32Int8(name string, value mapUint32Int8Value, usage string, required, isset bool) *mapUint32Int8Value {
	return CommandLine.mapUint32Int8(name, value, usage, required, isset)
}

// mapUint32Int16Value []mapUint32Int16Value
type mapUint32Int16Value map[uint32]int16

func newmapUint32Int16Value(val mapUint32Int16Value,
	p *mapUint32Int16Value) *mapUint32Int16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint32Int16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint32Int16Value) Get() interface{} { return map[uint32]int16(*slc) }

// String join a string from map
func (slc *mapUint32Int16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint32Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint32Int16Var(p *mapUint32Int16Value, name string, value mapUint32Int16Value, usage string, required, isset bool) {
	f.Var(newmapUint32Int16Value(value, p), name, usage, required, isset)
}

// mapUint32Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint32Int16Var(p *mapUint32Int16Value, name string, value mapUint32Int16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint32Int16Value(value, p), name, usage, required, isset)
}

// mapUint32Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint32Int16(name string, value mapUint32Int16Value, usage string, required, isset bool) *mapUint32Int16Value {
	p := new(mapUint32Int16Value)
	f.mapUint32Int16Var(p, name, value, usage, required, isset)
	return p
}

// mapUint32Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint32Int16(name string, value mapUint32Int16Value, usage string, required, isset bool) *mapUint32Int16Value {
	return CommandLine.mapUint32Int16(name, value, usage, required, isset)
}

// mapUint32Int32Value []mapUint32Int32Value
type mapUint32Int32Value map[uint32]int32

func newmapUint32Int32Value(val mapUint32Int32Value,
	p *mapUint32Int32Value) *mapUint32Int32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint32Int32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint32Int32Value) Get() interface{} { return map[uint32]int32(*slc) }

// String join a string from map
func (slc *mapUint32Int32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint32Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint32Int32Var(p *mapUint32Int32Value, name string, value mapUint32Int32Value, usage string, required, isset bool) {
	f.Var(newmapUint32Int32Value(value, p), name, usage, required, isset)
}

// mapUint32Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint32Int32Var(p *mapUint32Int32Value, name string, value mapUint32Int32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint32Int32Value(value, p), name, usage, required, isset)
}

// mapUint32Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint32Int32(name string, value mapUint32Int32Value, usage string, required, isset bool) *mapUint32Int32Value {
	p := new(mapUint32Int32Value)
	f.mapUint32Int32Var(p, name, value, usage, required, isset)
	return p
}

// mapUint32Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint32Int32(name string, value mapUint32Int32Value, usage string, required, isset bool) *mapUint32Int32Value {
	return CommandLine.mapUint32Int32(name, value, usage, required, isset)
}

// mapUint32Int64Value []mapUint32Int64Value
type mapUint32Int64Value map[uint32]int64

func newmapUint32Int64Value(val mapUint32Int64Value,
	p *mapUint32Int64Value) *mapUint32Int64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint32Int64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint32Int64Value) Get() interface{} { return map[uint32]int64(*slc) }

// String join a string from map
func (slc *mapUint32Int64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint32Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint32Int64Var(p *mapUint32Int64Value, name string, value mapUint32Int64Value, usage string, required, isset bool) {
	f.Var(newmapUint32Int64Value(value, p), name, usage, required, isset)
}

// mapUint32Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint32Int64Var(p *mapUint32Int64Value, name string, value mapUint32Int64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint32Int64Value(value, p), name, usage, required, isset)
}

// mapUint32Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint32Int64(name string, value mapUint32Int64Value, usage string, required, isset bool) *mapUint32Int64Value {
	p := new(mapUint32Int64Value)
	f.mapUint32Int64Var(p, name, value, usage, required, isset)
	return p
}

// mapUint32Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint32Int64(name string, value mapUint32Int64Value, usage string, required, isset bool) *mapUint32Int64Value {
	return CommandLine.mapUint32Int64(name, value, usage, required, isset)
}

// mapUint32UintValue []mapUint32UintValue
type mapUint32UintValue map[uint32]uint

func newmapUint32UintValue(val mapUint32UintValue,
	p *mapUint32UintValue) *mapUint32UintValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint32UintValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint32UintValue) Get() interface{} { return map[uint32]uint(*slc) }

// String join a string from map
func (slc *mapUint32UintValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint32UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint32UintVar(p *mapUint32UintValue, name string, value mapUint32UintValue, usage string, required, isset bool) {
	f.Var(newmapUint32UintValue(value, p), name, usage, required, isset)
}

// mapUint32UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint32UintVar(p *mapUint32UintValue, name string, value mapUint32UintValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint32UintValue(value, p), name, usage, required, isset)
}

// mapUint32UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint32Uint(name string, value mapUint32UintValue, usage string, required, isset bool) *mapUint32UintValue {
	p := new(mapUint32UintValue)
	f.mapUint32UintVar(p, name, value, usage, required, isset)
	return p
}

// mapUint32UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint32Uint(name string, value mapUint32UintValue, usage string, required, isset bool) *mapUint32UintValue {
	return CommandLine.mapUint32Uint(name, value, usage, required, isset)
}

// mapUint32Uint8Value []mapUint32Uint8Value
type mapUint32Uint8Value map[uint32]uint8

func newmapUint32Uint8Value(val mapUint32Uint8Value,
	p *mapUint32Uint8Value) *mapUint32Uint8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint32Uint8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint32Uint8Value) Get() interface{} { return map[uint32]uint8(*slc) }

// String join a string from map
func (slc *mapUint32Uint8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint32Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint32Uint8Var(p *mapUint32Uint8Value, name string, value mapUint32Uint8Value, usage string, required, isset bool) {
	f.Var(newmapUint32Uint8Value(value, p), name, usage, required, isset)
}

// mapUint32Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint32Uint8Var(p *mapUint32Uint8Value, name string, value mapUint32Uint8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint32Uint8Value(value, p), name, usage, required, isset)
}

// mapUint32Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint32Uint8(name string, value mapUint32Uint8Value, usage string, required, isset bool) *mapUint32Uint8Value {
	p := new(mapUint32Uint8Value)
	f.mapUint32Uint8Var(p, name, value, usage, required, isset)
	return p
}

// mapUint32Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint32Uint8(name string, value mapUint32Uint8Value, usage string, required, isset bool) *mapUint32Uint8Value {
	return CommandLine.mapUint32Uint8(name, value, usage, required, isset)
}

// mapUint32Uint16Value []mapUint32Uint16Value
type mapUint32Uint16Value map[uint32]uint16

func newmapUint32Uint16Value(val mapUint32Uint16Value,
	p *mapUint32Uint16Value) *mapUint32Uint16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint32Uint16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint32Uint16Value) Get() interface{} { return map[uint32]uint16(*slc) }

// String join a string from map
func (slc *mapUint32Uint16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint32Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint32Uint16Var(p *mapUint32Uint16Value, name string, value mapUint32Uint16Value, usage string, required, isset bool) {
	f.Var(newmapUint32Uint16Value(value, p), name, usage, required, isset)
}

// mapUint32Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint32Uint16Var(p *mapUint32Uint16Value, name string, value mapUint32Uint16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint32Uint16Value(value, p), name, usage, required, isset)
}

// mapUint32Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint32Uint16(name string, value mapUint32Uint16Value, usage string, required, isset bool) *mapUint32Uint16Value {
	p := new(mapUint32Uint16Value)
	f.mapUint32Uint16Var(p, name, value, usage, required, isset)
	return p
}

// mapUint32Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint32Uint16(name string, value mapUint32Uint16Value, usage string, required, isset bool) *mapUint32Uint16Value {
	return CommandLine.mapUint32Uint16(name, value, usage, required, isset)
}

// mapUint32Uint32Value []mapUint32Uint32Value
type mapUint32Uint32Value map[uint32]uint32

func newmapUint32Uint32Value(val mapUint32Uint32Value,
	p *mapUint32Uint32Value) *mapUint32Uint32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint32Uint32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint32Uint32Value) Get() interface{} { return map[uint32]uint32(*slc) }

// String join a string from map
func (slc *mapUint32Uint32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint32Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint32Uint32Var(p *mapUint32Uint32Value, name string, value mapUint32Uint32Value, usage string, required, isset bool) {
	f.Var(newmapUint32Uint32Value(value, p), name, usage, required, isset)
}

// mapUint32Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint32Uint32Var(p *mapUint32Uint32Value, name string, value mapUint32Uint32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint32Uint32Value(value, p), name, usage, required, isset)
}

// mapUint32Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint32Uint32(name string, value mapUint32Uint32Value, usage string, required, isset bool) *mapUint32Uint32Value {
	p := new(mapUint32Uint32Value)
	f.mapUint32Uint32Var(p, name, value, usage, required, isset)
	return p
}

// mapUint32Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint32Uint32(name string, value mapUint32Uint32Value, usage string, required, isset bool) *mapUint32Uint32Value {
	return CommandLine.mapUint32Uint32(name, value, usage, required, isset)
}

// mapUint32Uint64Value []mapUint32Uint64Value
type mapUint32Uint64Value map[uint32]uint64

func newmapUint32Uint64Value(val mapUint32Uint64Value,
	p *mapUint32Uint64Value) *mapUint32Uint64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint32Uint64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint32Uint64Value) Get() interface{} { return map[uint32]uint64(*slc) }

// String join a string from map
func (slc *mapUint32Uint64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint32Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint32Uint64Var(p *mapUint32Uint64Value, name string, value mapUint32Uint64Value, usage string, required, isset bool) {
	f.Var(newmapUint32Uint64Value(value, p), name, usage, required, isset)
}

// mapUint32Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint32Uint64Var(p *mapUint32Uint64Value, name string, value mapUint32Uint64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint32Uint64Value(value, p), name, usage, required, isset)
}

// mapUint32Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint32Uint64(name string, value mapUint32Uint64Value, usage string, required, isset bool) *mapUint32Uint64Value {
	p := new(mapUint32Uint64Value)
	f.mapUint32Uint64Var(p, name, value, usage, required, isset)
	return p
}

// mapUint32Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint32Uint64(name string, value mapUint32Uint64Value, usage string, required, isset bool) *mapUint32Uint64Value {
	return CommandLine.mapUint32Uint64(name, value, usage, required, isset)
}

// mapUint32Float64Value []mapUint32Float64Value
type mapUint32Float64Value map[uint32]float64

func newmapUint32Float64Value(val mapUint32Float64Value,
	p *mapUint32Float64Value) *mapUint32Float64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint32Float64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint32Float64Value) Get() interface{} { return map[uint32]float64(*slc) }

// String join a string from map
func (slc *mapUint32Float64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint32Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint32Float64Var(p *mapUint32Float64Value, name string, value mapUint32Float64Value, usage string, required, isset bool) {
	f.Var(newmapUint32Float64Value(value, p), name, usage, required, isset)
}

// mapUint32Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint32Float64Var(p *mapUint32Float64Value, name string, value mapUint32Float64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint32Float64Value(value, p), name, usage, required, isset)
}

// mapUint32Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint32Float64(name string, value mapUint32Float64Value, usage string, required, isset bool) *mapUint32Float64Value {
	p := new(mapUint32Float64Value)
	f.mapUint32Float64Var(p, name, value, usage, required, isset)
	return p
}

// mapUint32Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint32Float64(name string, value mapUint32Float64Value, usage string, required, isset bool) *mapUint32Float64Value {
	return CommandLine.mapUint32Float64(name, value, usage, required, isset)
}

// mapUint32Float32Value []mapUint32Float32Value
type mapUint32Float32Value map[uint32]float32

func newmapUint32Float32Value(val mapUint32Float32Value,
	p *mapUint32Float32Value) *mapUint32Float32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint32Float32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint32Float32Value) Get() interface{} { return map[uint32]float32(*slc) }

// String join a string from map
func (slc *mapUint32Float32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint32Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint32Float32Var(p *mapUint32Float32Value, name string, value mapUint32Float32Value, usage string, required, isset bool) {
	f.Var(newmapUint32Float32Value(value, p), name, usage, required, isset)
}

// mapUint32Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint32Float32Var(p *mapUint32Float32Value, name string, value mapUint32Float32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint32Float32Value(value, p), name, usage, required, isset)
}

// mapUint32Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint32Float32(name string, value mapUint32Float32Value, usage string, required, isset bool) *mapUint32Float32Value {
	p := new(mapUint32Float32Value)
	f.mapUint32Float32Var(p, name, value, usage, required, isset)
	return p
}

// mapUint32Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint32Float32(name string, value mapUint32Float32Value, usage string, required, isset bool) *mapUint32Float32Value {
	return CommandLine.mapUint32Float32(name, value, usage, required, isset)
}

// mapUint32BoolValue []mapUint32BoolValue
type mapUint32BoolValue map[uint32]bool

func newmapUint32BoolValue(val mapUint32BoolValue,
	p *mapUint32BoolValue) *mapUint32BoolValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint32BoolValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint32BoolValue) Get() interface{} { return map[uint32]bool(*slc) }

// String join a string from map
func (slc *mapUint32BoolValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint32BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint32BoolVar(p *mapUint32BoolValue, name string, value mapUint32BoolValue, usage string, required, isset bool) {
	f.Var(newmapUint32BoolValue(value, p), name, usage, required, isset)
}

// mapUint32BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint32BoolVar(p *mapUint32BoolValue, name string, value mapUint32BoolValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint32BoolValue(value, p), name, usage, required, isset)
}

// mapUint32BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint32Bool(name string, value mapUint32BoolValue, usage string, required, isset bool) *mapUint32BoolValue {
	p := new(mapUint32BoolValue)
	f.mapUint32BoolVar(p, name, value, usage, required, isset)
	return p
}

// mapUint32BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint32Bool(name string, value mapUint32BoolValue, usage string, required, isset bool) *mapUint32BoolValue {
	return CommandLine.mapUint32Bool(name, value, usage, required, isset)
}

// mapUint32StringValue []mapUint32StringValue
type mapUint32StringValue map[uint32]string

func newmapUint32StringValue(val mapUint32StringValue,
	p *mapUint32StringValue) *mapUint32StringValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint32StringValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint32StringValue) Get() interface{} { return map[uint32]string(*slc) }

// String join a string from map
func (slc *mapUint32StringValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint32StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint32StringVar(p *mapUint32StringValue, name string, value mapUint32StringValue, usage string, required, isset bool) {
	f.Var(newmapUint32StringValue(value, p), name, usage, required, isset)
}

// mapUint32StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint32StringVar(p *mapUint32StringValue, name string, value mapUint32StringValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint32StringValue(value, p), name, usage, required, isset)
}

// mapUint32StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint32String(name string, value mapUint32StringValue, usage string, required, isset bool) *mapUint32StringValue {
	p := new(mapUint32StringValue)
	f.mapUint32StringVar(p, name, value, usage, required, isset)
	return p
}

// mapUint32StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint32String(name string, value mapUint32StringValue, usage string, required, isset bool) *mapUint32StringValue {
	return CommandLine.mapUint32String(name, value, usage, required, isset)
}

// mapUint64DurationValue []mapUint64DurationValue
type mapUint64DurationValue map[uint64]time.Duration

func newmapUint64DurationValue(val mapUint64DurationValue,
	p *mapUint64DurationValue) *mapUint64DurationValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint64DurationValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint64DurationValue) Get() interface{} { return map[uint64]time.Duration(*slc) }

// String join a string from map
func (slc *mapUint64DurationValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint64DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint64DurationVar(p *mapUint64DurationValue, name string, value mapUint64DurationValue, usage string, required, isset bool) {
	f.Var(newmapUint64DurationValue(value, p), name, usage, required, isset)
}

// mapUint64DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint64DurationVar(p *mapUint64DurationValue, name string, value mapUint64DurationValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint64DurationValue(value, p), name, usage, required, isset)
}

// mapUint64DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint64Duration(name string, value mapUint64DurationValue, usage string, required, isset bool) *mapUint64DurationValue {
	p := new(mapUint64DurationValue)
	f.mapUint64DurationVar(p, name, value, usage, required, isset)
	return p
}

// mapUint64DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint64Duration(name string, value mapUint64DurationValue, usage string, required, isset bool) *mapUint64DurationValue {
	return CommandLine.mapUint64Duration(name, value, usage, required, isset)
}

// mapUint64IntValue []mapUint64IntValue
type mapUint64IntValue map[uint64]int

func newmapUint64IntValue(val mapUint64IntValue,
	p *mapUint64IntValue) *mapUint64IntValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint64IntValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint64IntValue) Get() interface{} { return map[uint64]int(*slc) }

// String join a string from map
func (slc *mapUint64IntValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint64IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint64IntVar(p *mapUint64IntValue, name string, value mapUint64IntValue, usage string, required, isset bool) {
	f.Var(newmapUint64IntValue(value, p), name, usage, required, isset)
}

// mapUint64IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint64IntVar(p *mapUint64IntValue, name string, value mapUint64IntValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint64IntValue(value, p), name, usage, required, isset)
}

// mapUint64IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint64Int(name string, value mapUint64IntValue, usage string, required, isset bool) *mapUint64IntValue {
	p := new(mapUint64IntValue)
	f.mapUint64IntVar(p, name, value, usage, required, isset)
	return p
}

// mapUint64IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint64Int(name string, value mapUint64IntValue, usage string, required, isset bool) *mapUint64IntValue {
	return CommandLine.mapUint64Int(name, value, usage, required, isset)
}

// mapUint64Int8Value []mapUint64Int8Value
type mapUint64Int8Value map[uint64]int8

func newmapUint64Int8Value(val mapUint64Int8Value,
	p *mapUint64Int8Value) *mapUint64Int8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint64Int8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint64Int8Value) Get() interface{} { return map[uint64]int8(*slc) }

// String join a string from map
func (slc *mapUint64Int8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint64Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint64Int8Var(p *mapUint64Int8Value, name string, value mapUint64Int8Value, usage string, required, isset bool) {
	f.Var(newmapUint64Int8Value(value, p), name, usage, required, isset)
}

// mapUint64Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint64Int8Var(p *mapUint64Int8Value, name string, value mapUint64Int8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint64Int8Value(value, p), name, usage, required, isset)
}

// mapUint64Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint64Int8(name string, value mapUint64Int8Value, usage string, required, isset bool) *mapUint64Int8Value {
	p := new(mapUint64Int8Value)
	f.mapUint64Int8Var(p, name, value, usage, required, isset)
	return p
}

// mapUint64Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint64Int8(name string, value mapUint64Int8Value, usage string, required, isset bool) *mapUint64Int8Value {
	return CommandLine.mapUint64Int8(name, value, usage, required, isset)
}

// mapUint64Int16Value []mapUint64Int16Value
type mapUint64Int16Value map[uint64]int16

func newmapUint64Int16Value(val mapUint64Int16Value,
	p *mapUint64Int16Value) *mapUint64Int16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint64Int16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint64Int16Value) Get() interface{} { return map[uint64]int16(*slc) }

// String join a string from map
func (slc *mapUint64Int16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint64Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint64Int16Var(p *mapUint64Int16Value, name string, value mapUint64Int16Value, usage string, required, isset bool) {
	f.Var(newmapUint64Int16Value(value, p), name, usage, required, isset)
}

// mapUint64Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint64Int16Var(p *mapUint64Int16Value, name string, value mapUint64Int16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint64Int16Value(value, p), name, usage, required, isset)
}

// mapUint64Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint64Int16(name string, value mapUint64Int16Value, usage string, required, isset bool) *mapUint64Int16Value {
	p := new(mapUint64Int16Value)
	f.mapUint64Int16Var(p, name, value, usage, required, isset)
	return p
}

// mapUint64Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint64Int16(name string, value mapUint64Int16Value, usage string, required, isset bool) *mapUint64Int16Value {
	return CommandLine.mapUint64Int16(name, value, usage, required, isset)
}

// mapUint64Int32Value []mapUint64Int32Value
type mapUint64Int32Value map[uint64]int32

func newmapUint64Int32Value(val mapUint64Int32Value,
	p *mapUint64Int32Value) *mapUint64Int32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint64Int32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint64Int32Value) Get() interface{} { return map[uint64]int32(*slc) }

// String join a string from map
func (slc *mapUint64Int32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint64Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint64Int32Var(p *mapUint64Int32Value, name string, value mapUint64Int32Value, usage string, required, isset bool) {
	f.Var(newmapUint64Int32Value(value, p), name, usage, required, isset)
}

// mapUint64Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint64Int32Var(p *mapUint64Int32Value, name string, value mapUint64Int32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint64Int32Value(value, p), name, usage, required, isset)
}

// mapUint64Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint64Int32(name string, value mapUint64Int32Value, usage string, required, isset bool) *mapUint64Int32Value {
	p := new(mapUint64Int32Value)
	f.mapUint64Int32Var(p, name, value, usage, required, isset)
	return p
}

// mapUint64Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint64Int32(name string, value mapUint64Int32Value, usage string, required, isset bool) *mapUint64Int32Value {
	return CommandLine.mapUint64Int32(name, value, usage, required, isset)
}

// mapUint64Int64Value []mapUint64Int64Value
type mapUint64Int64Value map[uint64]int64

func newmapUint64Int64Value(val mapUint64Int64Value,
	p *mapUint64Int64Value) *mapUint64Int64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint64Int64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint64Int64Value) Get() interface{} { return map[uint64]int64(*slc) }

// String join a string from map
func (slc *mapUint64Int64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint64Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint64Int64Var(p *mapUint64Int64Value, name string, value mapUint64Int64Value, usage string, required, isset bool) {
	f.Var(newmapUint64Int64Value(value, p), name, usage, required, isset)
}

// mapUint64Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint64Int64Var(p *mapUint64Int64Value, name string, value mapUint64Int64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint64Int64Value(value, p), name, usage, required, isset)
}

// mapUint64Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint64Int64(name string, value mapUint64Int64Value, usage string, required, isset bool) *mapUint64Int64Value {
	p := new(mapUint64Int64Value)
	f.mapUint64Int64Var(p, name, value, usage, required, isset)
	return p
}

// mapUint64Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint64Int64(name string, value mapUint64Int64Value, usage string, required, isset bool) *mapUint64Int64Value {
	return CommandLine.mapUint64Int64(name, value, usage, required, isset)
}

// mapUint64UintValue []mapUint64UintValue
type mapUint64UintValue map[uint64]uint

func newmapUint64UintValue(val mapUint64UintValue,
	p *mapUint64UintValue) *mapUint64UintValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint64UintValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint64UintValue) Get() interface{} { return map[uint64]uint(*slc) }

// String join a string from map
func (slc *mapUint64UintValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint64UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint64UintVar(p *mapUint64UintValue, name string, value mapUint64UintValue, usage string, required, isset bool) {
	f.Var(newmapUint64UintValue(value, p), name, usage, required, isset)
}

// mapUint64UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint64UintVar(p *mapUint64UintValue, name string, value mapUint64UintValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint64UintValue(value, p), name, usage, required, isset)
}

// mapUint64UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint64Uint(name string, value mapUint64UintValue, usage string, required, isset bool) *mapUint64UintValue {
	p := new(mapUint64UintValue)
	f.mapUint64UintVar(p, name, value, usage, required, isset)
	return p
}

// mapUint64UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint64Uint(name string, value mapUint64UintValue, usage string, required, isset bool) *mapUint64UintValue {
	return CommandLine.mapUint64Uint(name, value, usage, required, isset)
}

// mapUint64Uint8Value []mapUint64Uint8Value
type mapUint64Uint8Value map[uint64]uint8

func newmapUint64Uint8Value(val mapUint64Uint8Value,
	p *mapUint64Uint8Value) *mapUint64Uint8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint64Uint8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint64Uint8Value) Get() interface{} { return map[uint64]uint8(*slc) }

// String join a string from map
func (slc *mapUint64Uint8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint64Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint64Uint8Var(p *mapUint64Uint8Value, name string, value mapUint64Uint8Value, usage string, required, isset bool) {
	f.Var(newmapUint64Uint8Value(value, p), name, usage, required, isset)
}

// mapUint64Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint64Uint8Var(p *mapUint64Uint8Value, name string, value mapUint64Uint8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint64Uint8Value(value, p), name, usage, required, isset)
}

// mapUint64Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint64Uint8(name string, value mapUint64Uint8Value, usage string, required, isset bool) *mapUint64Uint8Value {
	p := new(mapUint64Uint8Value)
	f.mapUint64Uint8Var(p, name, value, usage, required, isset)
	return p
}

// mapUint64Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint64Uint8(name string, value mapUint64Uint8Value, usage string, required, isset bool) *mapUint64Uint8Value {
	return CommandLine.mapUint64Uint8(name, value, usage, required, isset)
}

// mapUint64Uint16Value []mapUint64Uint16Value
type mapUint64Uint16Value map[uint64]uint16

func newmapUint64Uint16Value(val mapUint64Uint16Value,
	p *mapUint64Uint16Value) *mapUint64Uint16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint64Uint16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint64Uint16Value) Get() interface{} { return map[uint64]uint16(*slc) }

// String join a string from map
func (slc *mapUint64Uint16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint64Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint64Uint16Var(p *mapUint64Uint16Value, name string, value mapUint64Uint16Value, usage string, required, isset bool) {
	f.Var(newmapUint64Uint16Value(value, p), name, usage, required, isset)
}

// mapUint64Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint64Uint16Var(p *mapUint64Uint16Value, name string, value mapUint64Uint16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint64Uint16Value(value, p), name, usage, required, isset)
}

// mapUint64Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint64Uint16(name string, value mapUint64Uint16Value, usage string, required, isset bool) *mapUint64Uint16Value {
	p := new(mapUint64Uint16Value)
	f.mapUint64Uint16Var(p, name, value, usage, required, isset)
	return p
}

// mapUint64Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint64Uint16(name string, value mapUint64Uint16Value, usage string, required, isset bool) *mapUint64Uint16Value {
	return CommandLine.mapUint64Uint16(name, value, usage, required, isset)
}

// mapUint64Uint32Value []mapUint64Uint32Value
type mapUint64Uint32Value map[uint64]uint32

func newmapUint64Uint32Value(val mapUint64Uint32Value,
	p *mapUint64Uint32Value) *mapUint64Uint32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint64Uint32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint64Uint32Value) Get() interface{} { return map[uint64]uint32(*slc) }

// String join a string from map
func (slc *mapUint64Uint32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint64Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint64Uint32Var(p *mapUint64Uint32Value, name string, value mapUint64Uint32Value, usage string, required, isset bool) {
	f.Var(newmapUint64Uint32Value(value, p), name, usage, required, isset)
}

// mapUint64Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint64Uint32Var(p *mapUint64Uint32Value, name string, value mapUint64Uint32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint64Uint32Value(value, p), name, usage, required, isset)
}

// mapUint64Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint64Uint32(name string, value mapUint64Uint32Value, usage string, required, isset bool) *mapUint64Uint32Value {
	p := new(mapUint64Uint32Value)
	f.mapUint64Uint32Var(p, name, value, usage, required, isset)
	return p
}

// mapUint64Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint64Uint32(name string, value mapUint64Uint32Value, usage string, required, isset bool) *mapUint64Uint32Value {
	return CommandLine.mapUint64Uint32(name, value, usage, required, isset)
}

// mapUint64Uint64Value []mapUint64Uint64Value
type mapUint64Uint64Value map[uint64]uint64

func newmapUint64Uint64Value(val mapUint64Uint64Value,
	p *mapUint64Uint64Value) *mapUint64Uint64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint64Uint64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint64Uint64Value) Get() interface{} { return map[uint64]uint64(*slc) }

// String join a string from map
func (slc *mapUint64Uint64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint64Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint64Uint64Var(p *mapUint64Uint64Value, name string, value mapUint64Uint64Value, usage string, required, isset bool) {
	f.Var(newmapUint64Uint64Value(value, p), name, usage, required, isset)
}

// mapUint64Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint64Uint64Var(p *mapUint64Uint64Value, name string, value mapUint64Uint64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint64Uint64Value(value, p), name, usage, required, isset)
}

// mapUint64Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint64Uint64(name string, value mapUint64Uint64Value, usage string, required, isset bool) *mapUint64Uint64Value {
	p := new(mapUint64Uint64Value)
	f.mapUint64Uint64Var(p, name, value, usage, required, isset)
	return p
}

// mapUint64Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint64Uint64(name string, value mapUint64Uint64Value, usage string, required, isset bool) *mapUint64Uint64Value {
	return CommandLine.mapUint64Uint64(name, value, usage, required, isset)
}

// mapUint64Float64Value []mapUint64Float64Value
type mapUint64Float64Value map[uint64]float64

func newmapUint64Float64Value(val mapUint64Float64Value,
	p *mapUint64Float64Value) *mapUint64Float64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint64Float64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint64Float64Value) Get() interface{} { return map[uint64]float64(*slc) }

// String join a string from map
func (slc *mapUint64Float64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint64Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint64Float64Var(p *mapUint64Float64Value, name string, value mapUint64Float64Value, usage string, required, isset bool) {
	f.Var(newmapUint64Float64Value(value, p), name, usage, required, isset)
}

// mapUint64Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint64Float64Var(p *mapUint64Float64Value, name string, value mapUint64Float64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint64Float64Value(value, p), name, usage, required, isset)
}

// mapUint64Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint64Float64(name string, value mapUint64Float64Value, usage string, required, isset bool) *mapUint64Float64Value {
	p := new(mapUint64Float64Value)
	f.mapUint64Float64Var(p, name, value, usage, required, isset)
	return p
}

// mapUint64Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint64Float64(name string, value mapUint64Float64Value, usage string, required, isset bool) *mapUint64Float64Value {
	return CommandLine.mapUint64Float64(name, value, usage, required, isset)
}

// mapUint64Float32Value []mapUint64Float32Value
type mapUint64Float32Value map[uint64]float32

func newmapUint64Float32Value(val mapUint64Float32Value,
	p *mapUint64Float32Value) *mapUint64Float32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint64Float32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint64Float32Value) Get() interface{} { return map[uint64]float32(*slc) }

// String join a string from map
func (slc *mapUint64Float32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint64Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint64Float32Var(p *mapUint64Float32Value, name string, value mapUint64Float32Value, usage string, required, isset bool) {
	f.Var(newmapUint64Float32Value(value, p), name, usage, required, isset)
}

// mapUint64Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint64Float32Var(p *mapUint64Float32Value, name string, value mapUint64Float32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapUint64Float32Value(value, p), name, usage, required, isset)
}

// mapUint64Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint64Float32(name string, value mapUint64Float32Value, usage string, required, isset bool) *mapUint64Float32Value {
	p := new(mapUint64Float32Value)
	f.mapUint64Float32Var(p, name, value, usage, required, isset)
	return p
}

// mapUint64Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint64Float32(name string, value mapUint64Float32Value, usage string, required, isset bool) *mapUint64Float32Value {
	return CommandLine.mapUint64Float32(name, value, usage, required, isset)
}

// mapUint64BoolValue []mapUint64BoolValue
type mapUint64BoolValue map[uint64]bool

func newmapUint64BoolValue(val mapUint64BoolValue,
	p *mapUint64BoolValue) *mapUint64BoolValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint64BoolValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint64BoolValue) Get() interface{} { return map[uint64]bool(*slc) }

// String join a string from map
func (slc *mapUint64BoolValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint64BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint64BoolVar(p *mapUint64BoolValue, name string, value mapUint64BoolValue, usage string, required, isset bool) {
	f.Var(newmapUint64BoolValue(value, p), name, usage, required, isset)
}

// mapUint64BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint64BoolVar(p *mapUint64BoolValue, name string, value mapUint64BoolValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint64BoolValue(value, p), name, usage, required, isset)
}

// mapUint64BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint64Bool(name string, value mapUint64BoolValue, usage string, required, isset bool) *mapUint64BoolValue {
	p := new(mapUint64BoolValue)
	f.mapUint64BoolVar(p, name, value, usage, required, isset)
	return p
}

// mapUint64BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint64Bool(name string, value mapUint64BoolValue, usage string, required, isset bool) *mapUint64BoolValue {
	return CommandLine.mapUint64Bool(name, value, usage, required, isset)
}

// mapUint64StringValue []mapUint64StringValue
type mapUint64StringValue map[uint64]string

func newmapUint64StringValue(val mapUint64StringValue,
	p *mapUint64StringValue) *mapUint64StringValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapUint64StringValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapUint64StringValue) Get() interface{} { return map[uint64]string(*slc) }

// String join a string from map
func (slc *mapUint64StringValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapUint64StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapUint64StringVar(p *mapUint64StringValue, name string, value mapUint64StringValue, usage string, required, isset bool) {
	f.Var(newmapUint64StringValue(value, p), name, usage, required, isset)
}

// mapUint64StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapUint64StringVar(p *mapUint64StringValue, name string, value mapUint64StringValue, usage string, required, isset bool) {
	CommandLine.Var(newmapUint64StringValue(value, p), name, usage, required, isset)
}

// mapUint64StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapUint64String(name string, value mapUint64StringValue, usage string, required, isset bool) *mapUint64StringValue {
	p := new(mapUint64StringValue)
	f.mapUint64StringVar(p, name, value, usage, required, isset)
	return p
}

// mapUint64StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapUint64String(name string, value mapUint64StringValue, usage string, required, isset bool) *mapUint64StringValue {
	return CommandLine.mapUint64String(name, value, usage, required, isset)
}

// mapFloat64DurationValue []mapFloat64DurationValue
type mapFloat64DurationValue map[float64]time.Duration

func newmapFloat64DurationValue(val mapFloat64DurationValue,
	p *mapFloat64DurationValue) *mapFloat64DurationValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat64DurationValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat64DurationValue) Get() interface{} { return map[float64]time.Duration(*slc) }

// String join a string from map
func (slc *mapFloat64DurationValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat64DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat64DurationVar(p *mapFloat64DurationValue, name string, value mapFloat64DurationValue, usage string, required, isset bool) {
	f.Var(newmapFloat64DurationValue(value, p), name, usage, required, isset)
}

// mapFloat64DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat64DurationVar(p *mapFloat64DurationValue, name string, value mapFloat64DurationValue, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat64DurationValue(value, p), name, usage, required, isset)
}

// mapFloat64DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat64Duration(name string, value mapFloat64DurationValue, usage string, required, isset bool) *mapFloat64DurationValue {
	p := new(mapFloat64DurationValue)
	f.mapFloat64DurationVar(p, name, value, usage, required, isset)
	return p
}

// mapFloat64DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat64Duration(name string, value mapFloat64DurationValue, usage string, required, isset bool) *mapFloat64DurationValue {
	return CommandLine.mapFloat64Duration(name, value, usage, required, isset)
}

// mapFloat64IntValue []mapFloat64IntValue
type mapFloat64IntValue map[float64]int

func newmapFloat64IntValue(val mapFloat64IntValue,
	p *mapFloat64IntValue) *mapFloat64IntValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat64IntValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat64IntValue) Get() interface{} { return map[float64]int(*slc) }

// String join a string from map
func (slc *mapFloat64IntValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat64IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat64IntVar(p *mapFloat64IntValue, name string, value mapFloat64IntValue, usage string, required, isset bool) {
	f.Var(newmapFloat64IntValue(value, p), name, usage, required, isset)
}

// mapFloat64IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat64IntVar(p *mapFloat64IntValue, name string, value mapFloat64IntValue, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat64IntValue(value, p), name, usage, required, isset)
}

// mapFloat64IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat64Int(name string, value mapFloat64IntValue, usage string, required, isset bool) *mapFloat64IntValue {
	p := new(mapFloat64IntValue)
	f.mapFloat64IntVar(p, name, value, usage, required, isset)
	return p
}

// mapFloat64IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat64Int(name string, value mapFloat64IntValue, usage string, required, isset bool) *mapFloat64IntValue {
	return CommandLine.mapFloat64Int(name, value, usage, required, isset)
}

// mapFloat64Int8Value []mapFloat64Int8Value
type mapFloat64Int8Value map[float64]int8

func newmapFloat64Int8Value(val mapFloat64Int8Value,
	p *mapFloat64Int8Value) *mapFloat64Int8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat64Int8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat64Int8Value) Get() interface{} { return map[float64]int8(*slc) }

// String join a string from map
func (slc *mapFloat64Int8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat64Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat64Int8Var(p *mapFloat64Int8Value, name string, value mapFloat64Int8Value, usage string, required, isset bool) {
	f.Var(newmapFloat64Int8Value(value, p), name, usage, required, isset)
}

// mapFloat64Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat64Int8Var(p *mapFloat64Int8Value, name string, value mapFloat64Int8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat64Int8Value(value, p), name, usage, required, isset)
}

// mapFloat64Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat64Int8(name string, value mapFloat64Int8Value, usage string, required, isset bool) *mapFloat64Int8Value {
	p := new(mapFloat64Int8Value)
	f.mapFloat64Int8Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat64Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat64Int8(name string, value mapFloat64Int8Value, usage string, required, isset bool) *mapFloat64Int8Value {
	return CommandLine.mapFloat64Int8(name, value, usage, required, isset)
}

// mapFloat64Int16Value []mapFloat64Int16Value
type mapFloat64Int16Value map[float64]int16

func newmapFloat64Int16Value(val mapFloat64Int16Value,
	p *mapFloat64Int16Value) *mapFloat64Int16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat64Int16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat64Int16Value) Get() interface{} { return map[float64]int16(*slc) }

// String join a string from map
func (slc *mapFloat64Int16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat64Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat64Int16Var(p *mapFloat64Int16Value, name string, value mapFloat64Int16Value, usage string, required, isset bool) {
	f.Var(newmapFloat64Int16Value(value, p), name, usage, required, isset)
}

// mapFloat64Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat64Int16Var(p *mapFloat64Int16Value, name string, value mapFloat64Int16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat64Int16Value(value, p), name, usage, required, isset)
}

// mapFloat64Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat64Int16(name string, value mapFloat64Int16Value, usage string, required, isset bool) *mapFloat64Int16Value {
	p := new(mapFloat64Int16Value)
	f.mapFloat64Int16Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat64Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat64Int16(name string, value mapFloat64Int16Value, usage string, required, isset bool) *mapFloat64Int16Value {
	return CommandLine.mapFloat64Int16(name, value, usage, required, isset)
}

// mapFloat64Int32Value []mapFloat64Int32Value
type mapFloat64Int32Value map[float64]int32

func newmapFloat64Int32Value(val mapFloat64Int32Value,
	p *mapFloat64Int32Value) *mapFloat64Int32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat64Int32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat64Int32Value) Get() interface{} { return map[float64]int32(*slc) }

// String join a string from map
func (slc *mapFloat64Int32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat64Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat64Int32Var(p *mapFloat64Int32Value, name string, value mapFloat64Int32Value, usage string, required, isset bool) {
	f.Var(newmapFloat64Int32Value(value, p), name, usage, required, isset)
}

// mapFloat64Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat64Int32Var(p *mapFloat64Int32Value, name string, value mapFloat64Int32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat64Int32Value(value, p), name, usage, required, isset)
}

// mapFloat64Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat64Int32(name string, value mapFloat64Int32Value, usage string, required, isset bool) *mapFloat64Int32Value {
	p := new(mapFloat64Int32Value)
	f.mapFloat64Int32Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat64Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat64Int32(name string, value mapFloat64Int32Value, usage string, required, isset bool) *mapFloat64Int32Value {
	return CommandLine.mapFloat64Int32(name, value, usage, required, isset)
}

// mapFloat64Int64Value []mapFloat64Int64Value
type mapFloat64Int64Value map[float64]int64

func newmapFloat64Int64Value(val mapFloat64Int64Value,
	p *mapFloat64Int64Value) *mapFloat64Int64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat64Int64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat64Int64Value) Get() interface{} { return map[float64]int64(*slc) }

// String join a string from map
func (slc *mapFloat64Int64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat64Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat64Int64Var(p *mapFloat64Int64Value, name string, value mapFloat64Int64Value, usage string, required, isset bool) {
	f.Var(newmapFloat64Int64Value(value, p), name, usage, required, isset)
}

// mapFloat64Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat64Int64Var(p *mapFloat64Int64Value, name string, value mapFloat64Int64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat64Int64Value(value, p), name, usage, required, isset)
}

// mapFloat64Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat64Int64(name string, value mapFloat64Int64Value, usage string, required, isset bool) *mapFloat64Int64Value {
	p := new(mapFloat64Int64Value)
	f.mapFloat64Int64Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat64Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat64Int64(name string, value mapFloat64Int64Value, usage string, required, isset bool) *mapFloat64Int64Value {
	return CommandLine.mapFloat64Int64(name, value, usage, required, isset)
}

// mapFloat64UintValue []mapFloat64UintValue
type mapFloat64UintValue map[float64]uint

func newmapFloat64UintValue(val mapFloat64UintValue,
	p *mapFloat64UintValue) *mapFloat64UintValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat64UintValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat64UintValue) Get() interface{} { return map[float64]uint(*slc) }

// String join a string from map
func (slc *mapFloat64UintValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat64UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat64UintVar(p *mapFloat64UintValue, name string, value mapFloat64UintValue, usage string, required, isset bool) {
	f.Var(newmapFloat64UintValue(value, p), name, usage, required, isset)
}

// mapFloat64UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat64UintVar(p *mapFloat64UintValue, name string, value mapFloat64UintValue, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat64UintValue(value, p), name, usage, required, isset)
}

// mapFloat64UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat64Uint(name string, value mapFloat64UintValue, usage string, required, isset bool) *mapFloat64UintValue {
	p := new(mapFloat64UintValue)
	f.mapFloat64UintVar(p, name, value, usage, required, isset)
	return p
}

// mapFloat64UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat64Uint(name string, value mapFloat64UintValue, usage string, required, isset bool) *mapFloat64UintValue {
	return CommandLine.mapFloat64Uint(name, value, usage, required, isset)
}

// mapFloat64Uint8Value []mapFloat64Uint8Value
type mapFloat64Uint8Value map[float64]uint8

func newmapFloat64Uint8Value(val mapFloat64Uint8Value,
	p *mapFloat64Uint8Value) *mapFloat64Uint8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat64Uint8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat64Uint8Value) Get() interface{} { return map[float64]uint8(*slc) }

// String join a string from map
func (slc *mapFloat64Uint8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat64Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat64Uint8Var(p *mapFloat64Uint8Value, name string, value mapFloat64Uint8Value, usage string, required, isset bool) {
	f.Var(newmapFloat64Uint8Value(value, p), name, usage, required, isset)
}

// mapFloat64Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat64Uint8Var(p *mapFloat64Uint8Value, name string, value mapFloat64Uint8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat64Uint8Value(value, p), name, usage, required, isset)
}

// mapFloat64Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat64Uint8(name string, value mapFloat64Uint8Value, usage string, required, isset bool) *mapFloat64Uint8Value {
	p := new(mapFloat64Uint8Value)
	f.mapFloat64Uint8Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat64Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat64Uint8(name string, value mapFloat64Uint8Value, usage string, required, isset bool) *mapFloat64Uint8Value {
	return CommandLine.mapFloat64Uint8(name, value, usage, required, isset)
}

// mapFloat64Uint16Value []mapFloat64Uint16Value
type mapFloat64Uint16Value map[float64]uint16

func newmapFloat64Uint16Value(val mapFloat64Uint16Value,
	p *mapFloat64Uint16Value) *mapFloat64Uint16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat64Uint16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat64Uint16Value) Get() interface{} { return map[float64]uint16(*slc) }

// String join a string from map
func (slc *mapFloat64Uint16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat64Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat64Uint16Var(p *mapFloat64Uint16Value, name string, value mapFloat64Uint16Value, usage string, required, isset bool) {
	f.Var(newmapFloat64Uint16Value(value, p), name, usage, required, isset)
}

// mapFloat64Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat64Uint16Var(p *mapFloat64Uint16Value, name string, value mapFloat64Uint16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat64Uint16Value(value, p), name, usage, required, isset)
}

// mapFloat64Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat64Uint16(name string, value mapFloat64Uint16Value, usage string, required, isset bool) *mapFloat64Uint16Value {
	p := new(mapFloat64Uint16Value)
	f.mapFloat64Uint16Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat64Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat64Uint16(name string, value mapFloat64Uint16Value, usage string, required, isset bool) *mapFloat64Uint16Value {
	return CommandLine.mapFloat64Uint16(name, value, usage, required, isset)
}

// mapFloat64Uint32Value []mapFloat64Uint32Value
type mapFloat64Uint32Value map[float64]uint32

func newmapFloat64Uint32Value(val mapFloat64Uint32Value,
	p *mapFloat64Uint32Value) *mapFloat64Uint32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat64Uint32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat64Uint32Value) Get() interface{} { return map[float64]uint32(*slc) }

// String join a string from map
func (slc *mapFloat64Uint32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat64Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat64Uint32Var(p *mapFloat64Uint32Value, name string, value mapFloat64Uint32Value, usage string, required, isset bool) {
	f.Var(newmapFloat64Uint32Value(value, p), name, usage, required, isset)
}

// mapFloat64Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat64Uint32Var(p *mapFloat64Uint32Value, name string, value mapFloat64Uint32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat64Uint32Value(value, p), name, usage, required, isset)
}

// mapFloat64Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat64Uint32(name string, value mapFloat64Uint32Value, usage string, required, isset bool) *mapFloat64Uint32Value {
	p := new(mapFloat64Uint32Value)
	f.mapFloat64Uint32Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat64Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat64Uint32(name string, value mapFloat64Uint32Value, usage string, required, isset bool) *mapFloat64Uint32Value {
	return CommandLine.mapFloat64Uint32(name, value, usage, required, isset)
}

// mapFloat64Uint64Value []mapFloat64Uint64Value
type mapFloat64Uint64Value map[float64]uint64

func newmapFloat64Uint64Value(val mapFloat64Uint64Value,
	p *mapFloat64Uint64Value) *mapFloat64Uint64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat64Uint64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat64Uint64Value) Get() interface{} { return map[float64]uint64(*slc) }

// String join a string from map
func (slc *mapFloat64Uint64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat64Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat64Uint64Var(p *mapFloat64Uint64Value, name string, value mapFloat64Uint64Value, usage string, required, isset bool) {
	f.Var(newmapFloat64Uint64Value(value, p), name, usage, required, isset)
}

// mapFloat64Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat64Uint64Var(p *mapFloat64Uint64Value, name string, value mapFloat64Uint64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat64Uint64Value(value, p), name, usage, required, isset)
}

// mapFloat64Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat64Uint64(name string, value mapFloat64Uint64Value, usage string, required, isset bool) *mapFloat64Uint64Value {
	p := new(mapFloat64Uint64Value)
	f.mapFloat64Uint64Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat64Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat64Uint64(name string, value mapFloat64Uint64Value, usage string, required, isset bool) *mapFloat64Uint64Value {
	return CommandLine.mapFloat64Uint64(name, value, usage, required, isset)
}

// mapFloat64Float64Value []mapFloat64Float64Value
type mapFloat64Float64Value map[float64]float64

func newmapFloat64Float64Value(val mapFloat64Float64Value,
	p *mapFloat64Float64Value) *mapFloat64Float64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat64Float64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat64Float64Value) Get() interface{} { return map[float64]float64(*slc) }

// String join a string from map
func (slc *mapFloat64Float64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat64Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat64Float64Var(p *mapFloat64Float64Value, name string, value mapFloat64Float64Value, usage string, required, isset bool) {
	f.Var(newmapFloat64Float64Value(value, p), name, usage, required, isset)
}

// mapFloat64Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat64Float64Var(p *mapFloat64Float64Value, name string, value mapFloat64Float64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat64Float64Value(value, p), name, usage, required, isset)
}

// mapFloat64Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat64Float64(name string, value mapFloat64Float64Value, usage string, required, isset bool) *mapFloat64Float64Value {
	p := new(mapFloat64Float64Value)
	f.mapFloat64Float64Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat64Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat64Float64(name string, value mapFloat64Float64Value, usage string, required, isset bool) *mapFloat64Float64Value {
	return CommandLine.mapFloat64Float64(name, value, usage, required, isset)
}

// mapFloat64Float32Value []mapFloat64Float32Value
type mapFloat64Float32Value map[float64]float32

func newmapFloat64Float32Value(val mapFloat64Float32Value,
	p *mapFloat64Float32Value) *mapFloat64Float32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat64Float32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat64Float32Value) Get() interface{} { return map[float64]float32(*slc) }

// String join a string from map
func (slc *mapFloat64Float32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat64Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat64Float32Var(p *mapFloat64Float32Value, name string, value mapFloat64Float32Value, usage string, required, isset bool) {
	f.Var(newmapFloat64Float32Value(value, p), name, usage, required, isset)
}

// mapFloat64Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat64Float32Var(p *mapFloat64Float32Value, name string, value mapFloat64Float32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat64Float32Value(value, p), name, usage, required, isset)
}

// mapFloat64Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat64Float32(name string, value mapFloat64Float32Value, usage string, required, isset bool) *mapFloat64Float32Value {
	p := new(mapFloat64Float32Value)
	f.mapFloat64Float32Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat64Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat64Float32(name string, value mapFloat64Float32Value, usage string, required, isset bool) *mapFloat64Float32Value {
	return CommandLine.mapFloat64Float32(name, value, usage, required, isset)
}

// mapFloat64BoolValue []mapFloat64BoolValue
type mapFloat64BoolValue map[float64]bool

func newmapFloat64BoolValue(val mapFloat64BoolValue,
	p *mapFloat64BoolValue) *mapFloat64BoolValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat64BoolValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat64BoolValue) Get() interface{} { return map[float64]bool(*slc) }

// String join a string from map
func (slc *mapFloat64BoolValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat64BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat64BoolVar(p *mapFloat64BoolValue, name string, value mapFloat64BoolValue, usage string, required, isset bool) {
	f.Var(newmapFloat64BoolValue(value, p), name, usage, required, isset)
}

// mapFloat64BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat64BoolVar(p *mapFloat64BoolValue, name string, value mapFloat64BoolValue, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat64BoolValue(value, p), name, usage, required, isset)
}

// mapFloat64BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat64Bool(name string, value mapFloat64BoolValue, usage string, required, isset bool) *mapFloat64BoolValue {
	p := new(mapFloat64BoolValue)
	f.mapFloat64BoolVar(p, name, value, usage, required, isset)
	return p
}

// mapFloat64BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat64Bool(name string, value mapFloat64BoolValue, usage string, required, isset bool) *mapFloat64BoolValue {
	return CommandLine.mapFloat64Bool(name, value, usage, required, isset)
}

// mapFloat64StringValue []mapFloat64StringValue
type mapFloat64StringValue map[float64]string

func newmapFloat64StringValue(val mapFloat64StringValue,
	p *mapFloat64StringValue) *mapFloat64StringValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat64StringValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat64StringValue) Get() interface{} { return map[float64]string(*slc) }

// String join a string from map
func (slc *mapFloat64StringValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat64StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat64StringVar(p *mapFloat64StringValue, name string, value mapFloat64StringValue, usage string, required, isset bool) {
	f.Var(newmapFloat64StringValue(value, p), name, usage, required, isset)
}

// mapFloat64StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat64StringVar(p *mapFloat64StringValue, name string, value mapFloat64StringValue, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat64StringValue(value, p), name, usage, required, isset)
}

// mapFloat64StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat64String(name string, value mapFloat64StringValue, usage string, required, isset bool) *mapFloat64StringValue {
	p := new(mapFloat64StringValue)
	f.mapFloat64StringVar(p, name, value, usage, required, isset)
	return p
}

// mapFloat64StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat64String(name string, value mapFloat64StringValue, usage string, required, isset bool) *mapFloat64StringValue {
	return CommandLine.mapFloat64String(name, value, usage, required, isset)
}

// mapFloat32DurationValue []mapFloat32DurationValue
type mapFloat32DurationValue map[float32]time.Duration

func newmapFloat32DurationValue(val mapFloat32DurationValue,
	p *mapFloat32DurationValue) *mapFloat32DurationValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat32DurationValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat32DurationValue) Get() interface{} { return map[float32]time.Duration(*slc) }

// String join a string from map
func (slc *mapFloat32DurationValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat32DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat32DurationVar(p *mapFloat32DurationValue, name string, value mapFloat32DurationValue, usage string, required, isset bool) {
	f.Var(newmapFloat32DurationValue(value, p), name, usage, required, isset)
}

// mapFloat32DurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat32DurationVar(p *mapFloat32DurationValue, name string, value mapFloat32DurationValue, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat32DurationValue(value, p), name, usage, required, isset)
}

// mapFloat32DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat32Duration(name string, value mapFloat32DurationValue, usage string, required, isset bool) *mapFloat32DurationValue {
	p := new(mapFloat32DurationValue)
	f.mapFloat32DurationVar(p, name, value, usage, required, isset)
	return p
}

// mapFloat32DurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat32Duration(name string, value mapFloat32DurationValue, usage string, required, isset bool) *mapFloat32DurationValue {
	return CommandLine.mapFloat32Duration(name, value, usage, required, isset)
}

// mapFloat32IntValue []mapFloat32IntValue
type mapFloat32IntValue map[float32]int

func newmapFloat32IntValue(val mapFloat32IntValue,
	p *mapFloat32IntValue) *mapFloat32IntValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat32IntValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat32IntValue) Get() interface{} { return map[float32]int(*slc) }

// String join a string from map
func (slc *mapFloat32IntValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat32IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat32IntVar(p *mapFloat32IntValue, name string, value mapFloat32IntValue, usage string, required, isset bool) {
	f.Var(newmapFloat32IntValue(value, p), name, usage, required, isset)
}

// mapFloat32IntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat32IntVar(p *mapFloat32IntValue, name string, value mapFloat32IntValue, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat32IntValue(value, p), name, usage, required, isset)
}

// mapFloat32IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat32Int(name string, value mapFloat32IntValue, usage string, required, isset bool) *mapFloat32IntValue {
	p := new(mapFloat32IntValue)
	f.mapFloat32IntVar(p, name, value, usage, required, isset)
	return p
}

// mapFloat32IntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat32Int(name string, value mapFloat32IntValue, usage string, required, isset bool) *mapFloat32IntValue {
	return CommandLine.mapFloat32Int(name, value, usage, required, isset)
}

// mapFloat32Int8Value []mapFloat32Int8Value
type mapFloat32Int8Value map[float32]int8

func newmapFloat32Int8Value(val mapFloat32Int8Value,
	p *mapFloat32Int8Value) *mapFloat32Int8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat32Int8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat32Int8Value) Get() interface{} { return map[float32]int8(*slc) }

// String join a string from map
func (slc *mapFloat32Int8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat32Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat32Int8Var(p *mapFloat32Int8Value, name string, value mapFloat32Int8Value, usage string, required, isset bool) {
	f.Var(newmapFloat32Int8Value(value, p), name, usage, required, isset)
}

// mapFloat32Int8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat32Int8Var(p *mapFloat32Int8Value, name string, value mapFloat32Int8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat32Int8Value(value, p), name, usage, required, isset)
}

// mapFloat32Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat32Int8(name string, value mapFloat32Int8Value, usage string, required, isset bool) *mapFloat32Int8Value {
	p := new(mapFloat32Int8Value)
	f.mapFloat32Int8Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat32Int8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat32Int8(name string, value mapFloat32Int8Value, usage string, required, isset bool) *mapFloat32Int8Value {
	return CommandLine.mapFloat32Int8(name, value, usage, required, isset)
}

// mapFloat32Int16Value []mapFloat32Int16Value
type mapFloat32Int16Value map[float32]int16

func newmapFloat32Int16Value(val mapFloat32Int16Value,
	p *mapFloat32Int16Value) *mapFloat32Int16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat32Int16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat32Int16Value) Get() interface{} { return map[float32]int16(*slc) }

// String join a string from map
func (slc *mapFloat32Int16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat32Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat32Int16Var(p *mapFloat32Int16Value, name string, value mapFloat32Int16Value, usage string, required, isset bool) {
	f.Var(newmapFloat32Int16Value(value, p), name, usage, required, isset)
}

// mapFloat32Int16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat32Int16Var(p *mapFloat32Int16Value, name string, value mapFloat32Int16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat32Int16Value(value, p), name, usage, required, isset)
}

// mapFloat32Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat32Int16(name string, value mapFloat32Int16Value, usage string, required, isset bool) *mapFloat32Int16Value {
	p := new(mapFloat32Int16Value)
	f.mapFloat32Int16Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat32Int16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat32Int16(name string, value mapFloat32Int16Value, usage string, required, isset bool) *mapFloat32Int16Value {
	return CommandLine.mapFloat32Int16(name, value, usage, required, isset)
}

// mapFloat32Int32Value []mapFloat32Int32Value
type mapFloat32Int32Value map[float32]int32

func newmapFloat32Int32Value(val mapFloat32Int32Value,
	p *mapFloat32Int32Value) *mapFloat32Int32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat32Int32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat32Int32Value) Get() interface{} { return map[float32]int32(*slc) }

// String join a string from map
func (slc *mapFloat32Int32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat32Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat32Int32Var(p *mapFloat32Int32Value, name string, value mapFloat32Int32Value, usage string, required, isset bool) {
	f.Var(newmapFloat32Int32Value(value, p), name, usage, required, isset)
}

// mapFloat32Int32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat32Int32Var(p *mapFloat32Int32Value, name string, value mapFloat32Int32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat32Int32Value(value, p), name, usage, required, isset)
}

// mapFloat32Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat32Int32(name string, value mapFloat32Int32Value, usage string, required, isset bool) *mapFloat32Int32Value {
	p := new(mapFloat32Int32Value)
	f.mapFloat32Int32Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat32Int32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat32Int32(name string, value mapFloat32Int32Value, usage string, required, isset bool) *mapFloat32Int32Value {
	return CommandLine.mapFloat32Int32(name, value, usage, required, isset)
}

// mapFloat32Int64Value []mapFloat32Int64Value
type mapFloat32Int64Value map[float32]int64

func newmapFloat32Int64Value(val mapFloat32Int64Value,
	p *mapFloat32Int64Value) *mapFloat32Int64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat32Int64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat32Int64Value) Get() interface{} { return map[float32]int64(*slc) }

// String join a string from map
func (slc *mapFloat32Int64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat32Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat32Int64Var(p *mapFloat32Int64Value, name string, value mapFloat32Int64Value, usage string, required, isset bool) {
	f.Var(newmapFloat32Int64Value(value, p), name, usage, required, isset)
}

// mapFloat32Int64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat32Int64Var(p *mapFloat32Int64Value, name string, value mapFloat32Int64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat32Int64Value(value, p), name, usage, required, isset)
}

// mapFloat32Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat32Int64(name string, value mapFloat32Int64Value, usage string, required, isset bool) *mapFloat32Int64Value {
	p := new(mapFloat32Int64Value)
	f.mapFloat32Int64Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat32Int64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat32Int64(name string, value mapFloat32Int64Value, usage string, required, isset bool) *mapFloat32Int64Value {
	return CommandLine.mapFloat32Int64(name, value, usage, required, isset)
}

// mapFloat32UintValue []mapFloat32UintValue
type mapFloat32UintValue map[float32]uint

func newmapFloat32UintValue(val mapFloat32UintValue,
	p *mapFloat32UintValue) *mapFloat32UintValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat32UintValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat32UintValue) Get() interface{} { return map[float32]uint(*slc) }

// String join a string from map
func (slc *mapFloat32UintValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat32UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat32UintVar(p *mapFloat32UintValue, name string, value mapFloat32UintValue, usage string, required, isset bool) {
	f.Var(newmapFloat32UintValue(value, p), name, usage, required, isset)
}

// mapFloat32UintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat32UintVar(p *mapFloat32UintValue, name string, value mapFloat32UintValue, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat32UintValue(value, p), name, usage, required, isset)
}

// mapFloat32UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat32Uint(name string, value mapFloat32UintValue, usage string, required, isset bool) *mapFloat32UintValue {
	p := new(mapFloat32UintValue)
	f.mapFloat32UintVar(p, name, value, usage, required, isset)
	return p
}

// mapFloat32UintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat32Uint(name string, value mapFloat32UintValue, usage string, required, isset bool) *mapFloat32UintValue {
	return CommandLine.mapFloat32Uint(name, value, usage, required, isset)
}

// mapFloat32Uint8Value []mapFloat32Uint8Value
type mapFloat32Uint8Value map[float32]uint8

func newmapFloat32Uint8Value(val mapFloat32Uint8Value,
	p *mapFloat32Uint8Value) *mapFloat32Uint8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat32Uint8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat32Uint8Value) Get() interface{} { return map[float32]uint8(*slc) }

// String join a string from map
func (slc *mapFloat32Uint8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat32Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat32Uint8Var(p *mapFloat32Uint8Value, name string, value mapFloat32Uint8Value, usage string, required, isset bool) {
	f.Var(newmapFloat32Uint8Value(value, p), name, usage, required, isset)
}

// mapFloat32Uint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat32Uint8Var(p *mapFloat32Uint8Value, name string, value mapFloat32Uint8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat32Uint8Value(value, p), name, usage, required, isset)
}

// mapFloat32Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat32Uint8(name string, value mapFloat32Uint8Value, usage string, required, isset bool) *mapFloat32Uint8Value {
	p := new(mapFloat32Uint8Value)
	f.mapFloat32Uint8Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat32Uint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat32Uint8(name string, value mapFloat32Uint8Value, usage string, required, isset bool) *mapFloat32Uint8Value {
	return CommandLine.mapFloat32Uint8(name, value, usage, required, isset)
}

// mapFloat32Uint16Value []mapFloat32Uint16Value
type mapFloat32Uint16Value map[float32]uint16

func newmapFloat32Uint16Value(val mapFloat32Uint16Value,
	p *mapFloat32Uint16Value) *mapFloat32Uint16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat32Uint16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat32Uint16Value) Get() interface{} { return map[float32]uint16(*slc) }

// String join a string from map
func (slc *mapFloat32Uint16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat32Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat32Uint16Var(p *mapFloat32Uint16Value, name string, value mapFloat32Uint16Value, usage string, required, isset bool) {
	f.Var(newmapFloat32Uint16Value(value, p), name, usage, required, isset)
}

// mapFloat32Uint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat32Uint16Var(p *mapFloat32Uint16Value, name string, value mapFloat32Uint16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat32Uint16Value(value, p), name, usage, required, isset)
}

// mapFloat32Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat32Uint16(name string, value mapFloat32Uint16Value, usage string, required, isset bool) *mapFloat32Uint16Value {
	p := new(mapFloat32Uint16Value)
	f.mapFloat32Uint16Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat32Uint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat32Uint16(name string, value mapFloat32Uint16Value, usage string, required, isset bool) *mapFloat32Uint16Value {
	return CommandLine.mapFloat32Uint16(name, value, usage, required, isset)
}

// mapFloat32Uint32Value []mapFloat32Uint32Value
type mapFloat32Uint32Value map[float32]uint32

func newmapFloat32Uint32Value(val mapFloat32Uint32Value,
	p *mapFloat32Uint32Value) *mapFloat32Uint32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat32Uint32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat32Uint32Value) Get() interface{} { return map[float32]uint32(*slc) }

// String join a string from map
func (slc *mapFloat32Uint32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat32Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat32Uint32Var(p *mapFloat32Uint32Value, name string, value mapFloat32Uint32Value, usage string, required, isset bool) {
	f.Var(newmapFloat32Uint32Value(value, p), name, usage, required, isset)
}

// mapFloat32Uint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat32Uint32Var(p *mapFloat32Uint32Value, name string, value mapFloat32Uint32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat32Uint32Value(value, p), name, usage, required, isset)
}

// mapFloat32Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat32Uint32(name string, value mapFloat32Uint32Value, usage string, required, isset bool) *mapFloat32Uint32Value {
	p := new(mapFloat32Uint32Value)
	f.mapFloat32Uint32Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat32Uint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat32Uint32(name string, value mapFloat32Uint32Value, usage string, required, isset bool) *mapFloat32Uint32Value {
	return CommandLine.mapFloat32Uint32(name, value, usage, required, isset)
}

// mapFloat32Uint64Value []mapFloat32Uint64Value
type mapFloat32Uint64Value map[float32]uint64

func newmapFloat32Uint64Value(val mapFloat32Uint64Value,
	p *mapFloat32Uint64Value) *mapFloat32Uint64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat32Uint64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat32Uint64Value) Get() interface{} { return map[float32]uint64(*slc) }

// String join a string from map
func (slc *mapFloat32Uint64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat32Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat32Uint64Var(p *mapFloat32Uint64Value, name string, value mapFloat32Uint64Value, usage string, required, isset bool) {
	f.Var(newmapFloat32Uint64Value(value, p), name, usage, required, isset)
}

// mapFloat32Uint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat32Uint64Var(p *mapFloat32Uint64Value, name string, value mapFloat32Uint64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat32Uint64Value(value, p), name, usage, required, isset)
}

// mapFloat32Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat32Uint64(name string, value mapFloat32Uint64Value, usage string, required, isset bool) *mapFloat32Uint64Value {
	p := new(mapFloat32Uint64Value)
	f.mapFloat32Uint64Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat32Uint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat32Uint64(name string, value mapFloat32Uint64Value, usage string, required, isset bool) *mapFloat32Uint64Value {
	return CommandLine.mapFloat32Uint64(name, value, usage, required, isset)
}

// mapFloat32Float64Value []mapFloat32Float64Value
type mapFloat32Float64Value map[float32]float64

func newmapFloat32Float64Value(val mapFloat32Float64Value,
	p *mapFloat32Float64Value) *mapFloat32Float64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat32Float64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat32Float64Value) Get() interface{} { return map[float32]float64(*slc) }

// String join a string from map
func (slc *mapFloat32Float64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat32Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat32Float64Var(p *mapFloat32Float64Value, name string, value mapFloat32Float64Value, usage string, required, isset bool) {
	f.Var(newmapFloat32Float64Value(value, p), name, usage, required, isset)
}

// mapFloat32Float64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat32Float64Var(p *mapFloat32Float64Value, name string, value mapFloat32Float64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat32Float64Value(value, p), name, usage, required, isset)
}

// mapFloat32Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat32Float64(name string, value mapFloat32Float64Value, usage string, required, isset bool) *mapFloat32Float64Value {
	p := new(mapFloat32Float64Value)
	f.mapFloat32Float64Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat32Float64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat32Float64(name string, value mapFloat32Float64Value, usage string, required, isset bool) *mapFloat32Float64Value {
	return CommandLine.mapFloat32Float64(name, value, usage, required, isset)
}

// mapFloat32Float32Value []mapFloat32Float32Value
type mapFloat32Float32Value map[float32]float32

func newmapFloat32Float32Value(val mapFloat32Float32Value,
	p *mapFloat32Float32Value) *mapFloat32Float32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat32Float32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat32Float32Value) Get() interface{} { return map[float32]float32(*slc) }

// String join a string from map
func (slc *mapFloat32Float32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat32Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat32Float32Var(p *mapFloat32Float32Value, name string, value mapFloat32Float32Value, usage string, required, isset bool) {
	f.Var(newmapFloat32Float32Value(value, p), name, usage, required, isset)
}

// mapFloat32Float32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat32Float32Var(p *mapFloat32Float32Value, name string, value mapFloat32Float32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat32Float32Value(value, p), name, usage, required, isset)
}

// mapFloat32Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat32Float32(name string, value mapFloat32Float32Value, usage string, required, isset bool) *mapFloat32Float32Value {
	p := new(mapFloat32Float32Value)
	f.mapFloat32Float32Var(p, name, value, usage, required, isset)
	return p
}

// mapFloat32Float32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat32Float32(name string, value mapFloat32Float32Value, usage string, required, isset bool) *mapFloat32Float32Value {
	return CommandLine.mapFloat32Float32(name, value, usage, required, isset)
}

// mapFloat32BoolValue []mapFloat32BoolValue
type mapFloat32BoolValue map[float32]bool

func newmapFloat32BoolValue(val mapFloat32BoolValue,
	p *mapFloat32BoolValue) *mapFloat32BoolValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat32BoolValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat32BoolValue) Get() interface{} { return map[float32]bool(*slc) }

// String join a string from map
func (slc *mapFloat32BoolValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat32BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat32BoolVar(p *mapFloat32BoolValue, name string, value mapFloat32BoolValue, usage string, required, isset bool) {
	f.Var(newmapFloat32BoolValue(value, p), name, usage, required, isset)
}

// mapFloat32BoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat32BoolVar(p *mapFloat32BoolValue, name string, value mapFloat32BoolValue, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat32BoolValue(value, p), name, usage, required, isset)
}

// mapFloat32BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat32Bool(name string, value mapFloat32BoolValue, usage string, required, isset bool) *mapFloat32BoolValue {
	p := new(mapFloat32BoolValue)
	f.mapFloat32BoolVar(p, name, value, usage, required, isset)
	return p
}

// mapFloat32BoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat32Bool(name string, value mapFloat32BoolValue, usage string, required, isset bool) *mapFloat32BoolValue {
	return CommandLine.mapFloat32Bool(name, value, usage, required, isset)
}

// mapFloat32StringValue []mapFloat32StringValue
type mapFloat32StringValue map[float32]string

func newmapFloat32StringValue(val mapFloat32StringValue,
	p *mapFloat32StringValue) *mapFloat32StringValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapFloat32StringValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapFloat32StringValue) Get() interface{} { return map[float32]string(*slc) }

// String join a string from map
func (slc *mapFloat32StringValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapFloat32StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapFloat32StringVar(p *mapFloat32StringValue, name string, value mapFloat32StringValue, usage string, required, isset bool) {
	f.Var(newmapFloat32StringValue(value, p), name, usage, required, isset)
}

// mapFloat32StringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapFloat32StringVar(p *mapFloat32StringValue, name string, value mapFloat32StringValue, usage string, required, isset bool) {
	CommandLine.Var(newmapFloat32StringValue(value, p), name, usage, required, isset)
}

// mapFloat32StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapFloat32String(name string, value mapFloat32StringValue, usage string, required, isset bool) *mapFloat32StringValue {
	p := new(mapFloat32StringValue)
	f.mapFloat32StringVar(p, name, value, usage, required, isset)
	return p
}

// mapFloat32StringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapFloat32String(name string, value mapFloat32StringValue, usage string, required, isset bool) *mapFloat32StringValue {
	return CommandLine.mapFloat32String(name, value, usage, required, isset)
}

// mapBoolDurationValue []mapBoolDurationValue
type mapBoolDurationValue map[bool]time.Duration

func newmapBoolDurationValue(val mapBoolDurationValue,
	p *mapBoolDurationValue) *mapBoolDurationValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapBoolDurationValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapBoolDurationValue) Get() interface{} { return map[bool]time.Duration(*slc) }

// String join a string from map
func (slc *mapBoolDurationValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapBoolDurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapBoolDurationVar(p *mapBoolDurationValue, name string, value mapBoolDurationValue, usage string, required, isset bool) {
	f.Var(newmapBoolDurationValue(value, p), name, usage, required, isset)
}

// mapBoolDurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapBoolDurationVar(p *mapBoolDurationValue, name string, value mapBoolDurationValue, usage string, required, isset bool) {
	CommandLine.Var(newmapBoolDurationValue(value, p), name, usage, required, isset)
}

// mapBoolDurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapBoolDuration(name string, value mapBoolDurationValue, usage string, required, isset bool) *mapBoolDurationValue {
	p := new(mapBoolDurationValue)
	f.mapBoolDurationVar(p, name, value, usage, required, isset)
	return p
}

// mapBoolDurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapBoolDuration(name string, value mapBoolDurationValue, usage string, required, isset bool) *mapBoolDurationValue {
	return CommandLine.mapBoolDuration(name, value, usage, required, isset)
}

// mapBoolIntValue []mapBoolIntValue
type mapBoolIntValue map[bool]int

func newmapBoolIntValue(val mapBoolIntValue,
	p *mapBoolIntValue) *mapBoolIntValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapBoolIntValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapBoolIntValue) Get() interface{} { return map[bool]int(*slc) }

// String join a string from map
func (slc *mapBoolIntValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapBoolIntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapBoolIntVar(p *mapBoolIntValue, name string, value mapBoolIntValue, usage string, required, isset bool) {
	f.Var(newmapBoolIntValue(value, p), name, usage, required, isset)
}

// mapBoolIntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapBoolIntVar(p *mapBoolIntValue, name string, value mapBoolIntValue, usage string, required, isset bool) {
	CommandLine.Var(newmapBoolIntValue(value, p), name, usage, required, isset)
}

// mapBoolIntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapBoolInt(name string, value mapBoolIntValue, usage string, required, isset bool) *mapBoolIntValue {
	p := new(mapBoolIntValue)
	f.mapBoolIntVar(p, name, value, usage, required, isset)
	return p
}

// mapBoolIntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapBoolInt(name string, value mapBoolIntValue, usage string, required, isset bool) *mapBoolIntValue {
	return CommandLine.mapBoolInt(name, value, usage, required, isset)
}

// mapBoolInt8Value []mapBoolInt8Value
type mapBoolInt8Value map[bool]int8

func newmapBoolInt8Value(val mapBoolInt8Value,
	p *mapBoolInt8Value) *mapBoolInt8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapBoolInt8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapBoolInt8Value) Get() interface{} { return map[bool]int8(*slc) }

// String join a string from map
func (slc *mapBoolInt8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapBoolInt8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapBoolInt8Var(p *mapBoolInt8Value, name string, value mapBoolInt8Value, usage string, required, isset bool) {
	f.Var(newmapBoolInt8Value(value, p), name, usage, required, isset)
}

// mapBoolInt8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapBoolInt8Var(p *mapBoolInt8Value, name string, value mapBoolInt8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapBoolInt8Value(value, p), name, usage, required, isset)
}

// mapBoolInt8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapBoolInt8(name string, value mapBoolInt8Value, usage string, required, isset bool) *mapBoolInt8Value {
	p := new(mapBoolInt8Value)
	f.mapBoolInt8Var(p, name, value, usage, required, isset)
	return p
}

// mapBoolInt8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapBoolInt8(name string, value mapBoolInt8Value, usage string, required, isset bool) *mapBoolInt8Value {
	return CommandLine.mapBoolInt8(name, value, usage, required, isset)
}

// mapBoolInt16Value []mapBoolInt16Value
type mapBoolInt16Value map[bool]int16

func newmapBoolInt16Value(val mapBoolInt16Value,
	p *mapBoolInt16Value) *mapBoolInt16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapBoolInt16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapBoolInt16Value) Get() interface{} { return map[bool]int16(*slc) }

// String join a string from map
func (slc *mapBoolInt16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapBoolInt16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapBoolInt16Var(p *mapBoolInt16Value, name string, value mapBoolInt16Value, usage string, required, isset bool) {
	f.Var(newmapBoolInt16Value(value, p), name, usage, required, isset)
}

// mapBoolInt16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapBoolInt16Var(p *mapBoolInt16Value, name string, value mapBoolInt16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapBoolInt16Value(value, p), name, usage, required, isset)
}

// mapBoolInt16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapBoolInt16(name string, value mapBoolInt16Value, usage string, required, isset bool) *mapBoolInt16Value {
	p := new(mapBoolInt16Value)
	f.mapBoolInt16Var(p, name, value, usage, required, isset)
	return p
}

// mapBoolInt16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapBoolInt16(name string, value mapBoolInt16Value, usage string, required, isset bool) *mapBoolInt16Value {
	return CommandLine.mapBoolInt16(name, value, usage, required, isset)
}

// mapBoolInt32Value []mapBoolInt32Value
type mapBoolInt32Value map[bool]int32

func newmapBoolInt32Value(val mapBoolInt32Value,
	p *mapBoolInt32Value) *mapBoolInt32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapBoolInt32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapBoolInt32Value) Get() interface{} { return map[bool]int32(*slc) }

// String join a string from map
func (slc *mapBoolInt32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapBoolInt32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapBoolInt32Var(p *mapBoolInt32Value, name string, value mapBoolInt32Value, usage string, required, isset bool) {
	f.Var(newmapBoolInt32Value(value, p), name, usage, required, isset)
}

// mapBoolInt32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapBoolInt32Var(p *mapBoolInt32Value, name string, value mapBoolInt32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapBoolInt32Value(value, p), name, usage, required, isset)
}

// mapBoolInt32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapBoolInt32(name string, value mapBoolInt32Value, usage string, required, isset bool) *mapBoolInt32Value {
	p := new(mapBoolInt32Value)
	f.mapBoolInt32Var(p, name, value, usage, required, isset)
	return p
}

// mapBoolInt32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapBoolInt32(name string, value mapBoolInt32Value, usage string, required, isset bool) *mapBoolInt32Value {
	return CommandLine.mapBoolInt32(name, value, usage, required, isset)
}

// mapBoolInt64Value []mapBoolInt64Value
type mapBoolInt64Value map[bool]int64

func newmapBoolInt64Value(val mapBoolInt64Value,
	p *mapBoolInt64Value) *mapBoolInt64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapBoolInt64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapBoolInt64Value) Get() interface{} { return map[bool]int64(*slc) }

// String join a string from map
func (slc *mapBoolInt64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapBoolInt64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapBoolInt64Var(p *mapBoolInt64Value, name string, value mapBoolInt64Value, usage string, required, isset bool) {
	f.Var(newmapBoolInt64Value(value, p), name, usage, required, isset)
}

// mapBoolInt64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapBoolInt64Var(p *mapBoolInt64Value, name string, value mapBoolInt64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapBoolInt64Value(value, p), name, usage, required, isset)
}

// mapBoolInt64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapBoolInt64(name string, value mapBoolInt64Value, usage string, required, isset bool) *mapBoolInt64Value {
	p := new(mapBoolInt64Value)
	f.mapBoolInt64Var(p, name, value, usage, required, isset)
	return p
}

// mapBoolInt64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapBoolInt64(name string, value mapBoolInt64Value, usage string, required, isset bool) *mapBoolInt64Value {
	return CommandLine.mapBoolInt64(name, value, usage, required, isset)
}

// mapBoolUintValue []mapBoolUintValue
type mapBoolUintValue map[bool]uint

func newmapBoolUintValue(val mapBoolUintValue,
	p *mapBoolUintValue) *mapBoolUintValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapBoolUintValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapBoolUintValue) Get() interface{} { return map[bool]uint(*slc) }

// String join a string from map
func (slc *mapBoolUintValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapBoolUintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapBoolUintVar(p *mapBoolUintValue, name string, value mapBoolUintValue, usage string, required, isset bool) {
	f.Var(newmapBoolUintValue(value, p), name, usage, required, isset)
}

// mapBoolUintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapBoolUintVar(p *mapBoolUintValue, name string, value mapBoolUintValue, usage string, required, isset bool) {
	CommandLine.Var(newmapBoolUintValue(value, p), name, usage, required, isset)
}

// mapBoolUintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapBoolUint(name string, value mapBoolUintValue, usage string, required, isset bool) *mapBoolUintValue {
	p := new(mapBoolUintValue)
	f.mapBoolUintVar(p, name, value, usage, required, isset)
	return p
}

// mapBoolUintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapBoolUint(name string, value mapBoolUintValue, usage string, required, isset bool) *mapBoolUintValue {
	return CommandLine.mapBoolUint(name, value, usage, required, isset)
}

// mapBoolUint8Value []mapBoolUint8Value
type mapBoolUint8Value map[bool]uint8

func newmapBoolUint8Value(val mapBoolUint8Value,
	p *mapBoolUint8Value) *mapBoolUint8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapBoolUint8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapBoolUint8Value) Get() interface{} { return map[bool]uint8(*slc) }

// String join a string from map
func (slc *mapBoolUint8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapBoolUint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapBoolUint8Var(p *mapBoolUint8Value, name string, value mapBoolUint8Value, usage string, required, isset bool) {
	f.Var(newmapBoolUint8Value(value, p), name, usage, required, isset)
}

// mapBoolUint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapBoolUint8Var(p *mapBoolUint8Value, name string, value mapBoolUint8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapBoolUint8Value(value, p), name, usage, required, isset)
}

// mapBoolUint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapBoolUint8(name string, value mapBoolUint8Value, usage string, required, isset bool) *mapBoolUint8Value {
	p := new(mapBoolUint8Value)
	f.mapBoolUint8Var(p, name, value, usage, required, isset)
	return p
}

// mapBoolUint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapBoolUint8(name string, value mapBoolUint8Value, usage string, required, isset bool) *mapBoolUint8Value {
	return CommandLine.mapBoolUint8(name, value, usage, required, isset)
}

// mapBoolUint16Value []mapBoolUint16Value
type mapBoolUint16Value map[bool]uint16

func newmapBoolUint16Value(val mapBoolUint16Value,
	p *mapBoolUint16Value) *mapBoolUint16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapBoolUint16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapBoolUint16Value) Get() interface{} { return map[bool]uint16(*slc) }

// String join a string from map
func (slc *mapBoolUint16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapBoolUint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapBoolUint16Var(p *mapBoolUint16Value, name string, value mapBoolUint16Value, usage string, required, isset bool) {
	f.Var(newmapBoolUint16Value(value, p), name, usage, required, isset)
}

// mapBoolUint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapBoolUint16Var(p *mapBoolUint16Value, name string, value mapBoolUint16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapBoolUint16Value(value, p), name, usage, required, isset)
}

// mapBoolUint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapBoolUint16(name string, value mapBoolUint16Value, usage string, required, isset bool) *mapBoolUint16Value {
	p := new(mapBoolUint16Value)
	f.mapBoolUint16Var(p, name, value, usage, required, isset)
	return p
}

// mapBoolUint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapBoolUint16(name string, value mapBoolUint16Value, usage string, required, isset bool) *mapBoolUint16Value {
	return CommandLine.mapBoolUint16(name, value, usage, required, isset)
}

// mapBoolUint32Value []mapBoolUint32Value
type mapBoolUint32Value map[bool]uint32

func newmapBoolUint32Value(val mapBoolUint32Value,
	p *mapBoolUint32Value) *mapBoolUint32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapBoolUint32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapBoolUint32Value) Get() interface{} { return map[bool]uint32(*slc) }

// String join a string from map
func (slc *mapBoolUint32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapBoolUint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapBoolUint32Var(p *mapBoolUint32Value, name string, value mapBoolUint32Value, usage string, required, isset bool) {
	f.Var(newmapBoolUint32Value(value, p), name, usage, required, isset)
}

// mapBoolUint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapBoolUint32Var(p *mapBoolUint32Value, name string, value mapBoolUint32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapBoolUint32Value(value, p), name, usage, required, isset)
}

// mapBoolUint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapBoolUint32(name string, value mapBoolUint32Value, usage string, required, isset bool) *mapBoolUint32Value {
	p := new(mapBoolUint32Value)
	f.mapBoolUint32Var(p, name, value, usage, required, isset)
	return p
}

// mapBoolUint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapBoolUint32(name string, value mapBoolUint32Value, usage string, required, isset bool) *mapBoolUint32Value {
	return CommandLine.mapBoolUint32(name, value, usage, required, isset)
}

// mapBoolUint64Value []mapBoolUint64Value
type mapBoolUint64Value map[bool]uint64

func newmapBoolUint64Value(val mapBoolUint64Value,
	p *mapBoolUint64Value) *mapBoolUint64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapBoolUint64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapBoolUint64Value) Get() interface{} { return map[bool]uint64(*slc) }

// String join a string from map
func (slc *mapBoolUint64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapBoolUint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapBoolUint64Var(p *mapBoolUint64Value, name string, value mapBoolUint64Value, usage string, required, isset bool) {
	f.Var(newmapBoolUint64Value(value, p), name, usage, required, isset)
}

// mapBoolUint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapBoolUint64Var(p *mapBoolUint64Value, name string, value mapBoolUint64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapBoolUint64Value(value, p), name, usage, required, isset)
}

// mapBoolUint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapBoolUint64(name string, value mapBoolUint64Value, usage string, required, isset bool) *mapBoolUint64Value {
	p := new(mapBoolUint64Value)
	f.mapBoolUint64Var(p, name, value, usage, required, isset)
	return p
}

// mapBoolUint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapBoolUint64(name string, value mapBoolUint64Value, usage string, required, isset bool) *mapBoolUint64Value {
	return CommandLine.mapBoolUint64(name, value, usage, required, isset)
}

// mapBoolFloat64Value []mapBoolFloat64Value
type mapBoolFloat64Value map[bool]float64

func newmapBoolFloat64Value(val mapBoolFloat64Value,
	p *mapBoolFloat64Value) *mapBoolFloat64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapBoolFloat64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapBoolFloat64Value) Get() interface{} { return map[bool]float64(*slc) }

// String join a string from map
func (slc *mapBoolFloat64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapBoolFloat64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapBoolFloat64Var(p *mapBoolFloat64Value, name string, value mapBoolFloat64Value, usage string, required, isset bool) {
	f.Var(newmapBoolFloat64Value(value, p), name, usage, required, isset)
}

// mapBoolFloat64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapBoolFloat64Var(p *mapBoolFloat64Value, name string, value mapBoolFloat64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapBoolFloat64Value(value, p), name, usage, required, isset)
}

// mapBoolFloat64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapBoolFloat64(name string, value mapBoolFloat64Value, usage string, required, isset bool) *mapBoolFloat64Value {
	p := new(mapBoolFloat64Value)
	f.mapBoolFloat64Var(p, name, value, usage, required, isset)
	return p
}

// mapBoolFloat64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapBoolFloat64(name string, value mapBoolFloat64Value, usage string, required, isset bool) *mapBoolFloat64Value {
	return CommandLine.mapBoolFloat64(name, value, usage, required, isset)
}

// mapBoolFloat32Value []mapBoolFloat32Value
type mapBoolFloat32Value map[bool]float32

func newmapBoolFloat32Value(val mapBoolFloat32Value,
	p *mapBoolFloat32Value) *mapBoolFloat32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapBoolFloat32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapBoolFloat32Value) Get() interface{} { return map[bool]float32(*slc) }

// String join a string from map
func (slc *mapBoolFloat32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapBoolFloat32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapBoolFloat32Var(p *mapBoolFloat32Value, name string, value mapBoolFloat32Value, usage string, required, isset bool) {
	f.Var(newmapBoolFloat32Value(value, p), name, usage, required, isset)
}

// mapBoolFloat32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapBoolFloat32Var(p *mapBoolFloat32Value, name string, value mapBoolFloat32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapBoolFloat32Value(value, p), name, usage, required, isset)
}

// mapBoolFloat32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapBoolFloat32(name string, value mapBoolFloat32Value, usage string, required, isset bool) *mapBoolFloat32Value {
	p := new(mapBoolFloat32Value)
	f.mapBoolFloat32Var(p, name, value, usage, required, isset)
	return p
}

// mapBoolFloat32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapBoolFloat32(name string, value mapBoolFloat32Value, usage string, required, isset bool) *mapBoolFloat32Value {
	return CommandLine.mapBoolFloat32(name, value, usage, required, isset)
}

// mapBoolBoolValue []mapBoolBoolValue
type mapBoolBoolValue map[bool]bool

func newmapBoolBoolValue(val mapBoolBoolValue,
	p *mapBoolBoolValue) *mapBoolBoolValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapBoolBoolValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapBoolBoolValue) Get() interface{} { return map[bool]bool(*slc) }

// String join a string from map
func (slc *mapBoolBoolValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapBoolBoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapBoolBoolVar(p *mapBoolBoolValue, name string, value mapBoolBoolValue, usage string, required, isset bool) {
	f.Var(newmapBoolBoolValue(value, p), name, usage, required, isset)
}

// mapBoolBoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapBoolBoolVar(p *mapBoolBoolValue, name string, value mapBoolBoolValue, usage string, required, isset bool) {
	CommandLine.Var(newmapBoolBoolValue(value, p), name, usage, required, isset)
}

// mapBoolBoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapBoolBool(name string, value mapBoolBoolValue, usage string, required, isset bool) *mapBoolBoolValue {
	p := new(mapBoolBoolValue)
	f.mapBoolBoolVar(p, name, value, usage, required, isset)
	return p
}

// mapBoolBoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapBoolBool(name string, value mapBoolBoolValue, usage string, required, isset bool) *mapBoolBoolValue {
	return CommandLine.mapBoolBool(name, value, usage, required, isset)
}

// mapBoolStringValue []mapBoolStringValue
type mapBoolStringValue map[bool]string

func newmapBoolStringValue(val mapBoolStringValue,
	p *mapBoolStringValue) *mapBoolStringValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapBoolStringValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapBoolStringValue) Get() interface{} { return map[bool]string(*slc) }

// String join a string from map
func (slc *mapBoolStringValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapBoolStringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapBoolStringVar(p *mapBoolStringValue, name string, value mapBoolStringValue, usage string, required, isset bool) {
	f.Var(newmapBoolStringValue(value, p), name, usage, required, isset)
}

// mapBoolStringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapBoolStringVar(p *mapBoolStringValue, name string, value mapBoolStringValue, usage string, required, isset bool) {
	CommandLine.Var(newmapBoolStringValue(value, p), name, usage, required, isset)
}

// mapBoolStringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapBoolString(name string, value mapBoolStringValue, usage string, required, isset bool) *mapBoolStringValue {
	p := new(mapBoolStringValue)
	f.mapBoolStringVar(p, name, value, usage, required, isset)
	return p
}

// mapBoolStringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapBoolString(name string, value mapBoolStringValue, usage string, required, isset bool) *mapBoolStringValue {
	return CommandLine.mapBoolString(name, value, usage, required, isset)
}

// mapStringDurationValue []mapStringDurationValue
type mapStringDurationValue map[string]time.Duration

func newmapStringDurationValue(val mapStringDurationValue,
	p *mapStringDurationValue) *mapStringDurationValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapStringDurationValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapStringDurationValue) Get() interface{} { return map[string]time.Duration(*slc) }

// String join a string from map
func (slc *mapStringDurationValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapStringDurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapStringDurationVar(p *mapStringDurationValue, name string, value mapStringDurationValue, usage string, required, isset bool) {
	f.Var(newmapStringDurationValue(value, p), name, usage, required, isset)
}

// mapStringDurationValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapStringDurationVar(p *mapStringDurationValue, name string, value mapStringDurationValue, usage string, required, isset bool) {
	CommandLine.Var(newmapStringDurationValue(value, p), name, usage, required, isset)
}

// mapStringDurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapStringDuration(name string, value mapStringDurationValue, usage string, required, isset bool) *mapStringDurationValue {
	p := new(mapStringDurationValue)
	f.mapStringDurationVar(p, name, value, usage, required, isset)
	return p
}

// mapStringDurationValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapStringDuration(name string, value mapStringDurationValue, usage string, required, isset bool) *mapStringDurationValue {
	return CommandLine.mapStringDuration(name, value, usage, required, isset)
}

// mapStringIntValue []mapStringIntValue
type mapStringIntValue map[string]int

func newmapStringIntValue(val mapStringIntValue,
	p *mapStringIntValue) *mapStringIntValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapStringIntValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapStringIntValue) Get() interface{} { return map[string]int(*slc) }

// String join a string from map
func (slc *mapStringIntValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapStringIntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapStringIntVar(p *mapStringIntValue, name string, value mapStringIntValue, usage string, required, isset bool) {
	f.Var(newmapStringIntValue(value, p), name, usage, required, isset)
}

// mapStringIntValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapStringIntVar(p *mapStringIntValue, name string, value mapStringIntValue, usage string, required, isset bool) {
	CommandLine.Var(newmapStringIntValue(value, p), name, usage, required, isset)
}

// mapStringIntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapStringInt(name string, value mapStringIntValue, usage string, required, isset bool) *mapStringIntValue {
	p := new(mapStringIntValue)
	f.mapStringIntVar(p, name, value, usage, required, isset)
	return p
}

// mapStringIntValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapStringInt(name string, value mapStringIntValue, usage string, required, isset bool) *mapStringIntValue {
	return CommandLine.mapStringInt(name, value, usage, required, isset)
}

// mapStringInt8Value []mapStringInt8Value
type mapStringInt8Value map[string]int8

func newmapStringInt8Value(val mapStringInt8Value,
	p *mapStringInt8Value) *mapStringInt8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapStringInt8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapStringInt8Value) Get() interface{} { return map[string]int8(*slc) }

// String join a string from map
func (slc *mapStringInt8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapStringInt8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapStringInt8Var(p *mapStringInt8Value, name string, value mapStringInt8Value, usage string, required, isset bool) {
	f.Var(newmapStringInt8Value(value, p), name, usage, required, isset)
}

// mapStringInt8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapStringInt8Var(p *mapStringInt8Value, name string, value mapStringInt8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapStringInt8Value(value, p), name, usage, required, isset)
}

// mapStringInt8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapStringInt8(name string, value mapStringInt8Value, usage string, required, isset bool) *mapStringInt8Value {
	p := new(mapStringInt8Value)
	f.mapStringInt8Var(p, name, value, usage, required, isset)
	return p
}

// mapStringInt8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapStringInt8(name string, value mapStringInt8Value, usage string, required, isset bool) *mapStringInt8Value {
	return CommandLine.mapStringInt8(name, value, usage, required, isset)
}

// mapStringInt16Value []mapStringInt16Value
type mapStringInt16Value map[string]int16

func newmapStringInt16Value(val mapStringInt16Value,
	p *mapStringInt16Value) *mapStringInt16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapStringInt16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapStringInt16Value) Get() interface{} { return map[string]int16(*slc) }

// String join a string from map
func (slc *mapStringInt16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapStringInt16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapStringInt16Var(p *mapStringInt16Value, name string, value mapStringInt16Value, usage string, required, isset bool) {
	f.Var(newmapStringInt16Value(value, p), name, usage, required, isset)
}

// mapStringInt16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapStringInt16Var(p *mapStringInt16Value, name string, value mapStringInt16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapStringInt16Value(value, p), name, usage, required, isset)
}

// mapStringInt16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapStringInt16(name string, value mapStringInt16Value, usage string, required, isset bool) *mapStringInt16Value {
	p := new(mapStringInt16Value)
	f.mapStringInt16Var(p, name, value, usage, required, isset)
	return p
}

// mapStringInt16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapStringInt16(name string, value mapStringInt16Value, usage string, required, isset bool) *mapStringInt16Value {
	return CommandLine.mapStringInt16(name, value, usage, required, isset)
}

// mapStringInt32Value []mapStringInt32Value
type mapStringInt32Value map[string]int32

func newmapStringInt32Value(val mapStringInt32Value,
	p *mapStringInt32Value) *mapStringInt32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapStringInt32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapStringInt32Value) Get() interface{} { return map[string]int32(*slc) }

// String join a string from map
func (slc *mapStringInt32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapStringInt32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapStringInt32Var(p *mapStringInt32Value, name string, value mapStringInt32Value, usage string, required, isset bool) {
	f.Var(newmapStringInt32Value(value, p), name, usage, required, isset)
}

// mapStringInt32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapStringInt32Var(p *mapStringInt32Value, name string, value mapStringInt32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapStringInt32Value(value, p), name, usage, required, isset)
}

// mapStringInt32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapStringInt32(name string, value mapStringInt32Value, usage string, required, isset bool) *mapStringInt32Value {
	p := new(mapStringInt32Value)
	f.mapStringInt32Var(p, name, value, usage, required, isset)
	return p
}

// mapStringInt32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapStringInt32(name string, value mapStringInt32Value, usage string, required, isset bool) *mapStringInt32Value {
	return CommandLine.mapStringInt32(name, value, usage, required, isset)
}

// mapStringInt64Value []mapStringInt64Value
type mapStringInt64Value map[string]int64

func newmapStringInt64Value(val mapStringInt64Value,
	p *mapStringInt64Value) *mapStringInt64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapStringInt64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapStringInt64Value) Get() interface{} { return map[string]int64(*slc) }

// String join a string from map
func (slc *mapStringInt64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapStringInt64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapStringInt64Var(p *mapStringInt64Value, name string, value mapStringInt64Value, usage string, required, isset bool) {
	f.Var(newmapStringInt64Value(value, p), name, usage, required, isset)
}

// mapStringInt64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapStringInt64Var(p *mapStringInt64Value, name string, value mapStringInt64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapStringInt64Value(value, p), name, usage, required, isset)
}

// mapStringInt64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapStringInt64(name string, value mapStringInt64Value, usage string, required, isset bool) *mapStringInt64Value {
	p := new(mapStringInt64Value)
	f.mapStringInt64Var(p, name, value, usage, required, isset)
	return p
}

// mapStringInt64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapStringInt64(name string, value mapStringInt64Value, usage string, required, isset bool) *mapStringInt64Value {
	return CommandLine.mapStringInt64(name, value, usage, required, isset)
}

// mapStringUintValue []mapStringUintValue
type mapStringUintValue map[string]uint

func newmapStringUintValue(val mapStringUintValue,
	p *mapStringUintValue) *mapStringUintValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapStringUintValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapStringUintValue) Get() interface{} { return map[string]uint(*slc) }

// String join a string from map
func (slc *mapStringUintValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapStringUintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapStringUintVar(p *mapStringUintValue, name string, value mapStringUintValue, usage string, required, isset bool) {
	f.Var(newmapStringUintValue(value, p), name, usage, required, isset)
}

// mapStringUintValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapStringUintVar(p *mapStringUintValue, name string, value mapStringUintValue, usage string, required, isset bool) {
	CommandLine.Var(newmapStringUintValue(value, p), name, usage, required, isset)
}

// mapStringUintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapStringUint(name string, value mapStringUintValue, usage string, required, isset bool) *mapStringUintValue {
	p := new(mapStringUintValue)
	f.mapStringUintVar(p, name, value, usage, required, isset)
	return p
}

// mapStringUintValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapStringUint(name string, value mapStringUintValue, usage string, required, isset bool) *mapStringUintValue {
	return CommandLine.mapStringUint(name, value, usage, required, isset)
}

// mapStringUint8Value []mapStringUint8Value
type mapStringUint8Value map[string]uint8

func newmapStringUint8Value(val mapStringUint8Value,
	p *mapStringUint8Value) *mapStringUint8Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapStringUint8Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapStringUint8Value) Get() interface{} { return map[string]uint8(*slc) }

// String join a string from map
func (slc *mapStringUint8Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapStringUint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapStringUint8Var(p *mapStringUint8Value, name string, value mapStringUint8Value, usage string, required, isset bool) {
	f.Var(newmapStringUint8Value(value, p), name, usage, required, isset)
}

// mapStringUint8ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapStringUint8Var(p *mapStringUint8Value, name string, value mapStringUint8Value, usage string, required, isset bool) {
	CommandLine.Var(newmapStringUint8Value(value, p), name, usage, required, isset)
}

// mapStringUint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapStringUint8(name string, value mapStringUint8Value, usage string, required, isset bool) *mapStringUint8Value {
	p := new(mapStringUint8Value)
	f.mapStringUint8Var(p, name, value, usage, required, isset)
	return p
}

// mapStringUint8Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapStringUint8(name string, value mapStringUint8Value, usage string, required, isset bool) *mapStringUint8Value {
	return CommandLine.mapStringUint8(name, value, usage, required, isset)
}

// mapStringUint16Value []mapStringUint16Value
type mapStringUint16Value map[string]uint16

func newmapStringUint16Value(val mapStringUint16Value,
	p *mapStringUint16Value) *mapStringUint16Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapStringUint16Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapStringUint16Value) Get() interface{} { return map[string]uint16(*slc) }

// String join a string from map
func (slc *mapStringUint16Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapStringUint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapStringUint16Var(p *mapStringUint16Value, name string, value mapStringUint16Value, usage string, required, isset bool) {
	f.Var(newmapStringUint16Value(value, p), name, usage, required, isset)
}

// mapStringUint16ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapStringUint16Var(p *mapStringUint16Value, name string, value mapStringUint16Value, usage string, required, isset bool) {
	CommandLine.Var(newmapStringUint16Value(value, p), name, usage, required, isset)
}

// mapStringUint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapStringUint16(name string, value mapStringUint16Value, usage string, required, isset bool) *mapStringUint16Value {
	p := new(mapStringUint16Value)
	f.mapStringUint16Var(p, name, value, usage, required, isset)
	return p
}

// mapStringUint16Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapStringUint16(name string, value mapStringUint16Value, usage string, required, isset bool) *mapStringUint16Value {
	return CommandLine.mapStringUint16(name, value, usage, required, isset)
}

// mapStringUint32Value []mapStringUint32Value
type mapStringUint32Value map[string]uint32

func newmapStringUint32Value(val mapStringUint32Value,
	p *mapStringUint32Value) *mapStringUint32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapStringUint32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapStringUint32Value) Get() interface{} { return map[string]uint32(*slc) }

// String join a string from map
func (slc *mapStringUint32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapStringUint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapStringUint32Var(p *mapStringUint32Value, name string, value mapStringUint32Value, usage string, required, isset bool) {
	f.Var(newmapStringUint32Value(value, p), name, usage, required, isset)
}

// mapStringUint32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapStringUint32Var(p *mapStringUint32Value, name string, value mapStringUint32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapStringUint32Value(value, p), name, usage, required, isset)
}

// mapStringUint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapStringUint32(name string, value mapStringUint32Value, usage string, required, isset bool) *mapStringUint32Value {
	p := new(mapStringUint32Value)
	f.mapStringUint32Var(p, name, value, usage, required, isset)
	return p
}

// mapStringUint32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapStringUint32(name string, value mapStringUint32Value, usage string, required, isset bool) *mapStringUint32Value {
	return CommandLine.mapStringUint32(name, value, usage, required, isset)
}

// mapStringUint64Value []mapStringUint64Value
type mapStringUint64Value map[string]uint64

func newmapStringUint64Value(val mapStringUint64Value,
	p *mapStringUint64Value) *mapStringUint64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapStringUint64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapStringUint64Value) Get() interface{} { return map[string]uint64(*slc) }

// String join a string from map
func (slc *mapStringUint64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapStringUint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapStringUint64Var(p *mapStringUint64Value, name string, value mapStringUint64Value, usage string, required, isset bool) {
	f.Var(newmapStringUint64Value(value, p), name, usage, required, isset)
}

// mapStringUint64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapStringUint64Var(p *mapStringUint64Value, name string, value mapStringUint64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapStringUint64Value(value, p), name, usage, required, isset)
}

// mapStringUint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapStringUint64(name string, value mapStringUint64Value, usage string, required, isset bool) *mapStringUint64Value {
	p := new(mapStringUint64Value)
	f.mapStringUint64Var(p, name, value, usage, required, isset)
	return p
}

// mapStringUint64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapStringUint64(name string, value mapStringUint64Value, usage string, required, isset bool) *mapStringUint64Value {
	return CommandLine.mapStringUint64(name, value, usage, required, isset)
}

// mapStringFloat64Value []mapStringFloat64Value
type mapStringFloat64Value map[string]float64

func newmapStringFloat64Value(val mapStringFloat64Value,
	p *mapStringFloat64Value) *mapStringFloat64Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapStringFloat64Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapStringFloat64Value) Get() interface{} { return map[string]float64(*slc) }

// String join a string from map
func (slc *mapStringFloat64Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapStringFloat64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapStringFloat64Var(p *mapStringFloat64Value, name string, value mapStringFloat64Value, usage string, required, isset bool) {
	f.Var(newmapStringFloat64Value(value, p), name, usage, required, isset)
}

// mapStringFloat64ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapStringFloat64Var(p *mapStringFloat64Value, name string, value mapStringFloat64Value, usage string, required, isset bool) {
	CommandLine.Var(newmapStringFloat64Value(value, p), name, usage, required, isset)
}

// mapStringFloat64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapStringFloat64(name string, value mapStringFloat64Value, usage string, required, isset bool) *mapStringFloat64Value {
	p := new(mapStringFloat64Value)
	f.mapStringFloat64Var(p, name, value, usage, required, isset)
	return p
}

// mapStringFloat64Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapStringFloat64(name string, value mapStringFloat64Value, usage string, required, isset bool) *mapStringFloat64Value {
	return CommandLine.mapStringFloat64(name, value, usage, required, isset)
}

// mapStringFloat32Value []mapStringFloat32Value
type mapStringFloat32Value map[string]float32

func newmapStringFloat32Value(val mapStringFloat32Value,
	p *mapStringFloat32Value) *mapStringFloat32Value {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapStringFloat32Value) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapStringFloat32Value) Get() interface{} { return map[string]float32(*slc) }

// String join a string from map
func (slc *mapStringFloat32Value) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapStringFloat32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapStringFloat32Var(p *mapStringFloat32Value, name string, value mapStringFloat32Value, usage string, required, isset bool) {
	f.Var(newmapStringFloat32Value(value, p), name, usage, required, isset)
}

// mapStringFloat32ValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapStringFloat32Var(p *mapStringFloat32Value, name string, value mapStringFloat32Value, usage string, required, isset bool) {
	CommandLine.Var(newmapStringFloat32Value(value, p), name, usage, required, isset)
}

// mapStringFloat32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapStringFloat32(name string, value mapStringFloat32Value, usage string, required, isset bool) *mapStringFloat32Value {
	p := new(mapStringFloat32Value)
	f.mapStringFloat32Var(p, name, value, usage, required, isset)
	return p
}

// mapStringFloat32Value defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapStringFloat32(name string, value mapStringFloat32Value, usage string, required, isset bool) *mapStringFloat32Value {
	return CommandLine.mapStringFloat32(name, value, usage, required, isset)
}

// mapStringBoolValue []mapStringBoolValue
type mapStringBoolValue map[string]bool

func newmapStringBoolValue(val mapStringBoolValue,
	p *mapStringBoolValue) *mapStringBoolValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapStringBoolValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapStringBoolValue) Get() interface{} { return map[string]bool(*slc) }

// String join a string from map
func (slc *mapStringBoolValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapStringBoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapStringBoolVar(p *mapStringBoolValue, name string, value mapStringBoolValue, usage string, required, isset bool) {
	f.Var(newmapStringBoolValue(value, p), name, usage, required, isset)
}

// mapStringBoolValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapStringBoolVar(p *mapStringBoolValue, name string, value mapStringBoolValue, usage string, required, isset bool) {
	CommandLine.Var(newmapStringBoolValue(value, p), name, usage, required, isset)
}

// mapStringBoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapStringBool(name string, value mapStringBoolValue, usage string, required, isset bool) *mapStringBoolValue {
	p := new(mapStringBoolValue)
	f.mapStringBoolVar(p, name, value, usage, required, isset)
	return p
}

// mapStringBoolValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapStringBool(name string, value mapStringBoolValue, usage string, required, isset bool) *mapStringBoolValue {
	return CommandLine.mapStringBool(name, value, usage, required, isset)
}

// mapStringStringValue []mapStringStringValue
type mapStringStringValue map[string]string

func newmapStringStringValue(val mapStringStringValue,
	p *mapStringStringValue) *mapStringStringValue {
	for k, v := range val {
		(*p)[k] = v
	}
	return p
}

// Set a map after parsing a string
func (slc *mapStringStringValue) Set(s string) error {
	return MapFromText(slc, s)
}

// Get a map interface from the value
func (slc *mapStringStringValue) Get() interface{} { return map[string]string(*slc) }

// String join a string from map
func (slc *mapStringStringValue) String() string {
	t := []string{}
	for k, v := range *slc {
		t = append(t, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(t, ",")
}

// mapStringStringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func (f *FlagSet) mapStringStringVar(p *mapStringStringValue, name string, value mapStringStringValue, usage string, required, isset bool) {
	f.Var(newmapStringStringValue(value, p), name, usage, required, isset)
}

// mapStringStringValueVar defines an map flag with specified name,
// default value, and usage string.  The argument p points to an map
// variable in which to store the value of the flag.
func mapStringStringVar(p *mapStringStringValue, name string, value mapStringStringValue, usage string, required, isset bool) {
	CommandLine.Var(newmapStringStringValue(value, p), name, usage, required, isset)
}

// mapStringStringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func (f *FlagSet) mapStringString(name string, value mapStringStringValue, usage string, required, isset bool) *mapStringStringValue {
	p := new(mapStringStringValue)
	f.mapStringStringVar(p, name, value, usage, required, isset)
	return p
}

// mapStringStringValue defines an map flag with specified name,
// default value, and usage string.  The return value is the address
// of an map variable that stores the value of the flag.
func mapStringString(name string, value mapStringStringValue, usage string, required, isset bool) *mapStringStringValue {
	return CommandLine.mapStringString(name, value, usage, required, isset)
}
