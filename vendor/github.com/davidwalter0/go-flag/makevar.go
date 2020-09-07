package flag

import (
	"fmt"
	"strings"
	"time"
)

// TypeName typename from environment variable TYPENAME
// KeyName typename from environment variable KEYNAME
// SetterTypeName "typename" + "Value"
// MapCastType and MapSetterTypeName from transform of type & key
// Special case time.Duration

// MakeVar type cast to type with interface matching Set, Value
func MakeVar(addr interface{}, name, defaultValue, usage, override string, required, isset bool) {
	defaultValue = strings.TrimSpace(defaultValue)
	override = strings.TrimSpace(override)
	switch ptr := addr.(type) {

	case *map[time.Duration]time.Duration:
		var varType = (*mapDurationDurationValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationDurationValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationDurationValue %T %v\n", ptr, err))
			}
		}

	case *map[time.Duration]int:
		var varType = (*mapDurationIntValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationIntValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationIntValue %T %v\n", ptr, err))
			}
		}

	case *map[time.Duration]int8:
		var varType = (*mapDurationInt8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationInt8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationInt8Value %T %v\n", ptr, err))
			}
		}

	case *map[time.Duration]int16:
		var varType = (*mapDurationInt16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationInt16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationInt16Value %T %v\n", ptr, err))
			}
		}

	case *map[time.Duration]int32:
		var varType = (*mapDurationInt32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationInt32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationInt32Value %T %v\n", ptr, err))
			}
		}

	case *map[time.Duration]int64:
		var varType = (*mapDurationInt64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationInt64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationInt64Value %T %v\n", ptr, err))
			}
		}

	case *map[time.Duration]uint:
		var varType = (*mapDurationUintValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationUintValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationUintValue %T %v\n", ptr, err))
			}
		}

	case *map[time.Duration]uint8:
		var varType = (*mapDurationUint8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationUint8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationUint8Value %T %v\n", ptr, err))
			}
		}

	case *map[time.Duration]uint16:
		var varType = (*mapDurationUint16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationUint16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationUint16Value %T %v\n", ptr, err))
			}
		}

	case *map[time.Duration]uint32:
		var varType = (*mapDurationUint32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationUint32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationUint32Value %T %v\n", ptr, err))
			}
		}

	case *map[time.Duration]uint64:
		var varType = (*mapDurationUint64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationUint64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationUint64Value %T %v\n", ptr, err))
			}
		}

	case *map[time.Duration]float64:
		var varType = (*mapDurationFloat64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationFloat64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationFloat64Value %T %v\n", ptr, err))
			}
		}

	case *map[time.Duration]float32:
		var varType = (*mapDurationFloat32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationFloat32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationFloat32Value %T %v\n", ptr, err))
			}
		}

	case *map[time.Duration]bool:
		var varType = (*mapDurationBoolValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationBoolValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationBoolValue %T %v\n", ptr, err))
			}
		}

	case *map[time.Duration]string:
		var varType = (*mapDurationStringValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationStringValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapDurationStringValue %T %v\n", ptr, err))
			}
		}

	case *map[int]time.Duration:
		var varType = (*mapIntDurationValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapIntDurationValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapIntDurationValue %T %v\n", ptr, err))
			}
		}

	case *map[int]int:
		var varType = (*mapIntIntValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapIntIntValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapIntIntValue %T %v\n", ptr, err))
			}
		}

	case *map[int]int8:
		var varType = (*mapIntInt8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapIntInt8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapIntInt8Value %T %v\n", ptr, err))
			}
		}

	case *map[int]int16:
		var varType = (*mapIntInt16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapIntInt16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapIntInt16Value %T %v\n", ptr, err))
			}
		}

	case *map[int]int32:
		var varType = (*mapIntInt32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapIntInt32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapIntInt32Value %T %v\n", ptr, err))
			}
		}

	case *map[int]int64:
		var varType = (*mapIntInt64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapIntInt64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapIntInt64Value %T %v\n", ptr, err))
			}
		}

	case *map[int]uint:
		var varType = (*mapIntUintValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapIntUintValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapIntUintValue %T %v\n", ptr, err))
			}
		}

	case *map[int]uint8:
		var varType = (*mapIntUint8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapIntUint8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapIntUint8Value %T %v\n", ptr, err))
			}
		}

	case *map[int]uint16:
		var varType = (*mapIntUint16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapIntUint16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapIntUint16Value %T %v\n", ptr, err))
			}
		}

	case *map[int]uint32:
		var varType = (*mapIntUint32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapIntUint32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapIntUint32Value %T %v\n", ptr, err))
			}
		}

	case *map[int]uint64:
		var varType = (*mapIntUint64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapIntUint64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapIntUint64Value %T %v\n", ptr, err))
			}
		}

	case *map[int]float64:
		var varType = (*mapIntFloat64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapIntFloat64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapIntFloat64Value %T %v\n", ptr, err))
			}
		}

	case *map[int]float32:
		var varType = (*mapIntFloat32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapIntFloat32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapIntFloat32Value %T %v\n", ptr, err))
			}
		}

	case *map[int]bool:
		var varType = (*mapIntBoolValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapIntBoolValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapIntBoolValue %T %v\n", ptr, err))
			}
		}

	case *map[int]string:
		var varType = (*mapIntStringValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapIntStringValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapIntStringValue %T %v\n", ptr, err))
			}
		}

	case *map[int8]time.Duration:
		var varType = (*mapInt8DurationValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8DurationValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8DurationValue %T %v\n", ptr, err))
			}
		}

	case *map[int8]int:
		var varType = (*mapInt8IntValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8IntValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8IntValue %T %v\n", ptr, err))
			}
		}

	case *map[int8]int8:
		var varType = (*mapInt8Int8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Int8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Int8Value %T %v\n", ptr, err))
			}
		}

	case *map[int8]int16:
		var varType = (*mapInt8Int16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Int16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Int16Value %T %v\n", ptr, err))
			}
		}

	case *map[int8]int32:
		var varType = (*mapInt8Int32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Int32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Int32Value %T %v\n", ptr, err))
			}
		}

	case *map[int8]int64:
		var varType = (*mapInt8Int64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Int64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Int64Value %T %v\n", ptr, err))
			}
		}

	case *map[int8]uint:
		var varType = (*mapInt8UintValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8UintValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8UintValue %T %v\n", ptr, err))
			}
		}

	case *map[int8]uint8:
		var varType = (*mapInt8Uint8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Uint8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Uint8Value %T %v\n", ptr, err))
			}
		}

	case *map[int8]uint16:
		var varType = (*mapInt8Uint16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Uint16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Uint16Value %T %v\n", ptr, err))
			}
		}

	case *map[int8]uint32:
		var varType = (*mapInt8Uint32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Uint32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Uint32Value %T %v\n", ptr, err))
			}
		}

	case *map[int8]uint64:
		var varType = (*mapInt8Uint64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Uint64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Uint64Value %T %v\n", ptr, err))
			}
		}

	case *map[int8]float64:
		var varType = (*mapInt8Float64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Float64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Float64Value %T %v\n", ptr, err))
			}
		}

	case *map[int8]float32:
		var varType = (*mapInt8Float32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Float32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8Float32Value %T %v\n", ptr, err))
			}
		}

	case *map[int8]bool:
		var varType = (*mapInt8BoolValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8BoolValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8BoolValue %T %v\n", ptr, err))
			}
		}

	case *map[int8]string:
		var varType = (*mapInt8StringValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8StringValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt8StringValue %T %v\n", ptr, err))
			}
		}

	case *map[int16]time.Duration:
		var varType = (*mapInt16DurationValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16DurationValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16DurationValue %T %v\n", ptr, err))
			}
		}

	case *map[int16]int:
		var varType = (*mapInt16IntValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16IntValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16IntValue %T %v\n", ptr, err))
			}
		}

	case *map[int16]int8:
		var varType = (*mapInt16Int8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Int8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Int8Value %T %v\n", ptr, err))
			}
		}

	case *map[int16]int16:
		var varType = (*mapInt16Int16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Int16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Int16Value %T %v\n", ptr, err))
			}
		}

	case *map[int16]int32:
		var varType = (*mapInt16Int32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Int32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Int32Value %T %v\n", ptr, err))
			}
		}

	case *map[int16]int64:
		var varType = (*mapInt16Int64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Int64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Int64Value %T %v\n", ptr, err))
			}
		}

	case *map[int16]uint:
		var varType = (*mapInt16UintValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16UintValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16UintValue %T %v\n", ptr, err))
			}
		}

	case *map[int16]uint8:
		var varType = (*mapInt16Uint8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Uint8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Uint8Value %T %v\n", ptr, err))
			}
		}

	case *map[int16]uint16:
		var varType = (*mapInt16Uint16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Uint16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Uint16Value %T %v\n", ptr, err))
			}
		}

	case *map[int16]uint32:
		var varType = (*mapInt16Uint32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Uint32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Uint32Value %T %v\n", ptr, err))
			}
		}

	case *map[int16]uint64:
		var varType = (*mapInt16Uint64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Uint64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Uint64Value %T %v\n", ptr, err))
			}
		}

	case *map[int16]float64:
		var varType = (*mapInt16Float64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Float64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Float64Value %T %v\n", ptr, err))
			}
		}

	case *map[int16]float32:
		var varType = (*mapInt16Float32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Float32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16Float32Value %T %v\n", ptr, err))
			}
		}

	case *map[int16]bool:
		var varType = (*mapInt16BoolValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16BoolValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16BoolValue %T %v\n", ptr, err))
			}
		}

	case *map[int16]string:
		var varType = (*mapInt16StringValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16StringValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt16StringValue %T %v\n", ptr, err))
			}
		}

	case *map[int32]time.Duration:
		var varType = (*mapInt32DurationValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32DurationValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32DurationValue %T %v\n", ptr, err))
			}
		}

	case *map[int32]int:
		var varType = (*mapInt32IntValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32IntValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32IntValue %T %v\n", ptr, err))
			}
		}

	case *map[int32]int8:
		var varType = (*mapInt32Int8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Int8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Int8Value %T %v\n", ptr, err))
			}
		}

	case *map[int32]int16:
		var varType = (*mapInt32Int16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Int16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Int16Value %T %v\n", ptr, err))
			}
		}

	case *map[int32]int32:
		var varType = (*mapInt32Int32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Int32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Int32Value %T %v\n", ptr, err))
			}
		}

	case *map[int32]int64:
		var varType = (*mapInt32Int64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Int64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Int64Value %T %v\n", ptr, err))
			}
		}

	case *map[int32]uint:
		var varType = (*mapInt32UintValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32UintValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32UintValue %T %v\n", ptr, err))
			}
		}

	case *map[int32]uint8:
		var varType = (*mapInt32Uint8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Uint8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Uint8Value %T %v\n", ptr, err))
			}
		}

	case *map[int32]uint16:
		var varType = (*mapInt32Uint16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Uint16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Uint16Value %T %v\n", ptr, err))
			}
		}

	case *map[int32]uint32:
		var varType = (*mapInt32Uint32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Uint32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Uint32Value %T %v\n", ptr, err))
			}
		}

	case *map[int32]uint64:
		var varType = (*mapInt32Uint64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Uint64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Uint64Value %T %v\n", ptr, err))
			}
		}

	case *map[int32]float64:
		var varType = (*mapInt32Float64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Float64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Float64Value %T %v\n", ptr, err))
			}
		}

	case *map[int32]float32:
		var varType = (*mapInt32Float32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Float32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32Float32Value %T %v\n", ptr, err))
			}
		}

	case *map[int32]bool:
		var varType = (*mapInt32BoolValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32BoolValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32BoolValue %T %v\n", ptr, err))
			}
		}

	case *map[int32]string:
		var varType = (*mapInt32StringValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32StringValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt32StringValue %T %v\n", ptr, err))
			}
		}

	case *map[int64]time.Duration:
		var varType = (*mapInt64DurationValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64DurationValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64DurationValue %T %v\n", ptr, err))
			}
		}

	case *map[int64]int:
		var varType = (*mapInt64IntValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64IntValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64IntValue %T %v\n", ptr, err))
			}
		}

	case *map[int64]int8:
		var varType = (*mapInt64Int8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Int8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Int8Value %T %v\n", ptr, err))
			}
		}

	case *map[int64]int16:
		var varType = (*mapInt64Int16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Int16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Int16Value %T %v\n", ptr, err))
			}
		}

	case *map[int64]int32:
		var varType = (*mapInt64Int32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Int32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Int32Value %T %v\n", ptr, err))
			}
		}

	case *map[int64]int64:
		var varType = (*mapInt64Int64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Int64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Int64Value %T %v\n", ptr, err))
			}
		}

	case *map[int64]uint:
		var varType = (*mapInt64UintValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64UintValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64UintValue %T %v\n", ptr, err))
			}
		}

	case *map[int64]uint8:
		var varType = (*mapInt64Uint8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Uint8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Uint8Value %T %v\n", ptr, err))
			}
		}

	case *map[int64]uint16:
		var varType = (*mapInt64Uint16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Uint16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Uint16Value %T %v\n", ptr, err))
			}
		}

	case *map[int64]uint32:
		var varType = (*mapInt64Uint32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Uint32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Uint32Value %T %v\n", ptr, err))
			}
		}

	case *map[int64]uint64:
		var varType = (*mapInt64Uint64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Uint64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Uint64Value %T %v\n", ptr, err))
			}
		}

	case *map[int64]float64:
		var varType = (*mapInt64Float64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Float64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Float64Value %T %v\n", ptr, err))
			}
		}

	case *map[int64]float32:
		var varType = (*mapInt64Float32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Float32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64Float32Value %T %v\n", ptr, err))
			}
		}

	case *map[int64]bool:
		var varType = (*mapInt64BoolValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64BoolValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64BoolValue %T %v\n", ptr, err))
			}
		}

	case *map[int64]string:
		var varType = (*mapInt64StringValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64StringValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapInt64StringValue %T %v\n", ptr, err))
			}
		}

	case *map[uint]time.Duration:
		var varType = (*mapUintDurationValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUintDurationValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUintDurationValue %T %v\n", ptr, err))
			}
		}

	case *map[uint]int:
		var varType = (*mapUintIntValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUintIntValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUintIntValue %T %v\n", ptr, err))
			}
		}

	case *map[uint]int8:
		var varType = (*mapUintInt8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUintInt8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUintInt8Value %T %v\n", ptr, err))
			}
		}

	case *map[uint]int16:
		var varType = (*mapUintInt16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUintInt16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUintInt16Value %T %v\n", ptr, err))
			}
		}

	case *map[uint]int32:
		var varType = (*mapUintInt32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUintInt32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUintInt32Value %T %v\n", ptr, err))
			}
		}

	case *map[uint]int64:
		var varType = (*mapUintInt64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUintInt64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUintInt64Value %T %v\n", ptr, err))
			}
		}

	case *map[uint]uint:
		var varType = (*mapUintUintValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUintUintValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUintUintValue %T %v\n", ptr, err))
			}
		}

	case *map[uint]uint8:
		var varType = (*mapUintUint8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUintUint8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUintUint8Value %T %v\n", ptr, err))
			}
		}

	case *map[uint]uint16:
		var varType = (*mapUintUint16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUintUint16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUintUint16Value %T %v\n", ptr, err))
			}
		}

	case *map[uint]uint32:
		var varType = (*mapUintUint32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUintUint32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUintUint32Value %T %v\n", ptr, err))
			}
		}

	case *map[uint]uint64:
		var varType = (*mapUintUint64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUintUint64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUintUint64Value %T %v\n", ptr, err))
			}
		}

	case *map[uint]float64:
		var varType = (*mapUintFloat64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUintFloat64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUintFloat64Value %T %v\n", ptr, err))
			}
		}

	case *map[uint]float32:
		var varType = (*mapUintFloat32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUintFloat32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUintFloat32Value %T %v\n", ptr, err))
			}
		}

	case *map[uint]bool:
		var varType = (*mapUintBoolValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUintBoolValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUintBoolValue %T %v\n", ptr, err))
			}
		}

	case *map[uint]string:
		var varType = (*mapUintStringValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUintStringValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUintStringValue %T %v\n", ptr, err))
			}
		}

	case *map[uint8]time.Duration:
		var varType = (*mapUint8DurationValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8DurationValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8DurationValue %T %v\n", ptr, err))
			}
		}

	case *map[uint8]int:
		var varType = (*mapUint8IntValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8IntValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8IntValue %T %v\n", ptr, err))
			}
		}

	case *map[uint8]int8:
		var varType = (*mapUint8Int8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Int8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Int8Value %T %v\n", ptr, err))
			}
		}

	case *map[uint8]int16:
		var varType = (*mapUint8Int16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Int16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Int16Value %T %v\n", ptr, err))
			}
		}

	case *map[uint8]int32:
		var varType = (*mapUint8Int32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Int32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Int32Value %T %v\n", ptr, err))
			}
		}

	case *map[uint8]int64:
		var varType = (*mapUint8Int64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Int64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Int64Value %T %v\n", ptr, err))
			}
		}

	case *map[uint8]uint:
		var varType = (*mapUint8UintValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8UintValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8UintValue %T %v\n", ptr, err))
			}
		}

	case *map[uint8]uint8:
		var varType = (*mapUint8Uint8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Uint8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Uint8Value %T %v\n", ptr, err))
			}
		}

	case *map[uint8]uint16:
		var varType = (*mapUint8Uint16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Uint16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Uint16Value %T %v\n", ptr, err))
			}
		}

	case *map[uint8]uint32:
		var varType = (*mapUint8Uint32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Uint32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Uint32Value %T %v\n", ptr, err))
			}
		}

	case *map[uint8]uint64:
		var varType = (*mapUint8Uint64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Uint64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Uint64Value %T %v\n", ptr, err))
			}
		}

	case *map[uint8]float64:
		var varType = (*mapUint8Float64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Float64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Float64Value %T %v\n", ptr, err))
			}
		}

	case *map[uint8]float32:
		var varType = (*mapUint8Float32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Float32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8Float32Value %T %v\n", ptr, err))
			}
		}

	case *map[uint8]bool:
		var varType = (*mapUint8BoolValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8BoolValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8BoolValue %T %v\n", ptr, err))
			}
		}

	case *map[uint8]string:
		var varType = (*mapUint8StringValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8StringValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint8StringValue %T %v\n", ptr, err))
			}
		}

	case *map[uint16]time.Duration:
		var varType = (*mapUint16DurationValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16DurationValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16DurationValue %T %v\n", ptr, err))
			}
		}

	case *map[uint16]int:
		var varType = (*mapUint16IntValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16IntValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16IntValue %T %v\n", ptr, err))
			}
		}

	case *map[uint16]int8:
		var varType = (*mapUint16Int8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Int8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Int8Value %T %v\n", ptr, err))
			}
		}

	case *map[uint16]int16:
		var varType = (*mapUint16Int16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Int16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Int16Value %T %v\n", ptr, err))
			}
		}

	case *map[uint16]int32:
		var varType = (*mapUint16Int32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Int32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Int32Value %T %v\n", ptr, err))
			}
		}

	case *map[uint16]int64:
		var varType = (*mapUint16Int64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Int64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Int64Value %T %v\n", ptr, err))
			}
		}

	case *map[uint16]uint:
		var varType = (*mapUint16UintValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16UintValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16UintValue %T %v\n", ptr, err))
			}
		}

	case *map[uint16]uint8:
		var varType = (*mapUint16Uint8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Uint8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Uint8Value %T %v\n", ptr, err))
			}
		}

	case *map[uint16]uint16:
		var varType = (*mapUint16Uint16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Uint16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Uint16Value %T %v\n", ptr, err))
			}
		}

	case *map[uint16]uint32:
		var varType = (*mapUint16Uint32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Uint32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Uint32Value %T %v\n", ptr, err))
			}
		}

	case *map[uint16]uint64:
		var varType = (*mapUint16Uint64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Uint64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Uint64Value %T %v\n", ptr, err))
			}
		}

	case *map[uint16]float64:
		var varType = (*mapUint16Float64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Float64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Float64Value %T %v\n", ptr, err))
			}
		}

	case *map[uint16]float32:
		var varType = (*mapUint16Float32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Float32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16Float32Value %T %v\n", ptr, err))
			}
		}

	case *map[uint16]bool:
		var varType = (*mapUint16BoolValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16BoolValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16BoolValue %T %v\n", ptr, err))
			}
		}

	case *map[uint16]string:
		var varType = (*mapUint16StringValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16StringValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint16StringValue %T %v\n", ptr, err))
			}
		}

	case *map[uint32]time.Duration:
		var varType = (*mapUint32DurationValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32DurationValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32DurationValue %T %v\n", ptr, err))
			}
		}

	case *map[uint32]int:
		var varType = (*mapUint32IntValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32IntValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32IntValue %T %v\n", ptr, err))
			}
		}

	case *map[uint32]int8:
		var varType = (*mapUint32Int8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Int8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Int8Value %T %v\n", ptr, err))
			}
		}

	case *map[uint32]int16:
		var varType = (*mapUint32Int16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Int16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Int16Value %T %v\n", ptr, err))
			}
		}

	case *map[uint32]int32:
		var varType = (*mapUint32Int32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Int32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Int32Value %T %v\n", ptr, err))
			}
		}

	case *map[uint32]int64:
		var varType = (*mapUint32Int64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Int64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Int64Value %T %v\n", ptr, err))
			}
		}

	case *map[uint32]uint:
		var varType = (*mapUint32UintValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32UintValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32UintValue %T %v\n", ptr, err))
			}
		}

	case *map[uint32]uint8:
		var varType = (*mapUint32Uint8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Uint8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Uint8Value %T %v\n", ptr, err))
			}
		}

	case *map[uint32]uint16:
		var varType = (*mapUint32Uint16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Uint16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Uint16Value %T %v\n", ptr, err))
			}
		}

	case *map[uint32]uint32:
		var varType = (*mapUint32Uint32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Uint32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Uint32Value %T %v\n", ptr, err))
			}
		}

	case *map[uint32]uint64:
		var varType = (*mapUint32Uint64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Uint64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Uint64Value %T %v\n", ptr, err))
			}
		}

	case *map[uint32]float64:
		var varType = (*mapUint32Float64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Float64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Float64Value %T %v\n", ptr, err))
			}
		}

	case *map[uint32]float32:
		var varType = (*mapUint32Float32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Float32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32Float32Value %T %v\n", ptr, err))
			}
		}

	case *map[uint32]bool:
		var varType = (*mapUint32BoolValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32BoolValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32BoolValue %T %v\n", ptr, err))
			}
		}

	case *map[uint32]string:
		var varType = (*mapUint32StringValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32StringValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint32StringValue %T %v\n", ptr, err))
			}
		}

	case *map[uint64]time.Duration:
		var varType = (*mapUint64DurationValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64DurationValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64DurationValue %T %v\n", ptr, err))
			}
		}

	case *map[uint64]int:
		var varType = (*mapUint64IntValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64IntValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64IntValue %T %v\n", ptr, err))
			}
		}

	case *map[uint64]int8:
		var varType = (*mapUint64Int8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Int8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Int8Value %T %v\n", ptr, err))
			}
		}

	case *map[uint64]int16:
		var varType = (*mapUint64Int16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Int16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Int16Value %T %v\n", ptr, err))
			}
		}

	case *map[uint64]int32:
		var varType = (*mapUint64Int32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Int32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Int32Value %T %v\n", ptr, err))
			}
		}

	case *map[uint64]int64:
		var varType = (*mapUint64Int64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Int64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Int64Value %T %v\n", ptr, err))
			}
		}

	case *map[uint64]uint:
		var varType = (*mapUint64UintValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64UintValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64UintValue %T %v\n", ptr, err))
			}
		}

	case *map[uint64]uint8:
		var varType = (*mapUint64Uint8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Uint8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Uint8Value %T %v\n", ptr, err))
			}
		}

	case *map[uint64]uint16:
		var varType = (*mapUint64Uint16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Uint16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Uint16Value %T %v\n", ptr, err))
			}
		}

	case *map[uint64]uint32:
		var varType = (*mapUint64Uint32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Uint32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Uint32Value %T %v\n", ptr, err))
			}
		}

	case *map[uint64]uint64:
		var varType = (*mapUint64Uint64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Uint64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Uint64Value %T %v\n", ptr, err))
			}
		}

	case *map[uint64]float64:
		var varType = (*mapUint64Float64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Float64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Float64Value %T %v\n", ptr, err))
			}
		}

	case *map[uint64]float32:
		var varType = (*mapUint64Float32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Float32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64Float32Value %T %v\n", ptr, err))
			}
		}

	case *map[uint64]bool:
		var varType = (*mapUint64BoolValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64BoolValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64BoolValue %T %v\n", ptr, err))
			}
		}

	case *map[uint64]string:
		var varType = (*mapUint64StringValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64StringValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapUint64StringValue %T %v\n", ptr, err))
			}
		}

	case *map[float64]time.Duration:
		var varType = (*mapFloat64DurationValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64DurationValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64DurationValue %T %v\n", ptr, err))
			}
		}

	case *map[float64]int:
		var varType = (*mapFloat64IntValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64IntValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64IntValue %T %v\n", ptr, err))
			}
		}

	case *map[float64]int8:
		var varType = (*mapFloat64Int8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Int8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Int8Value %T %v\n", ptr, err))
			}
		}

	case *map[float64]int16:
		var varType = (*mapFloat64Int16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Int16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Int16Value %T %v\n", ptr, err))
			}
		}

	case *map[float64]int32:
		var varType = (*mapFloat64Int32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Int32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Int32Value %T %v\n", ptr, err))
			}
		}

	case *map[float64]int64:
		var varType = (*mapFloat64Int64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Int64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Int64Value %T %v\n", ptr, err))
			}
		}

	case *map[float64]uint:
		var varType = (*mapFloat64UintValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64UintValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64UintValue %T %v\n", ptr, err))
			}
		}

	case *map[float64]uint8:
		var varType = (*mapFloat64Uint8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Uint8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Uint8Value %T %v\n", ptr, err))
			}
		}

	case *map[float64]uint16:
		var varType = (*mapFloat64Uint16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Uint16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Uint16Value %T %v\n", ptr, err))
			}
		}

	case *map[float64]uint32:
		var varType = (*mapFloat64Uint32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Uint32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Uint32Value %T %v\n", ptr, err))
			}
		}

	case *map[float64]uint64:
		var varType = (*mapFloat64Uint64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Uint64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Uint64Value %T %v\n", ptr, err))
			}
		}

	case *map[float64]float64:
		var varType = (*mapFloat64Float64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Float64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Float64Value %T %v\n", ptr, err))
			}
		}

	case *map[float64]float32:
		var varType = (*mapFloat64Float32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Float32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64Float32Value %T %v\n", ptr, err))
			}
		}

	case *map[float64]bool:
		var varType = (*mapFloat64BoolValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64BoolValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64BoolValue %T %v\n", ptr, err))
			}
		}

	case *map[float64]string:
		var varType = (*mapFloat64StringValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64StringValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat64StringValue %T %v\n", ptr, err))
			}
		}

	case *map[float32]time.Duration:
		var varType = (*mapFloat32DurationValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32DurationValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32DurationValue %T %v\n", ptr, err))
			}
		}

	case *map[float32]int:
		var varType = (*mapFloat32IntValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32IntValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32IntValue %T %v\n", ptr, err))
			}
		}

	case *map[float32]int8:
		var varType = (*mapFloat32Int8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Int8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Int8Value %T %v\n", ptr, err))
			}
		}

	case *map[float32]int16:
		var varType = (*mapFloat32Int16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Int16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Int16Value %T %v\n", ptr, err))
			}
		}

	case *map[float32]int32:
		var varType = (*mapFloat32Int32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Int32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Int32Value %T %v\n", ptr, err))
			}
		}

	case *map[float32]int64:
		var varType = (*mapFloat32Int64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Int64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Int64Value %T %v\n", ptr, err))
			}
		}

	case *map[float32]uint:
		var varType = (*mapFloat32UintValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32UintValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32UintValue %T %v\n", ptr, err))
			}
		}

	case *map[float32]uint8:
		var varType = (*mapFloat32Uint8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Uint8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Uint8Value %T %v\n", ptr, err))
			}
		}

	case *map[float32]uint16:
		var varType = (*mapFloat32Uint16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Uint16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Uint16Value %T %v\n", ptr, err))
			}
		}

	case *map[float32]uint32:
		var varType = (*mapFloat32Uint32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Uint32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Uint32Value %T %v\n", ptr, err))
			}
		}

	case *map[float32]uint64:
		var varType = (*mapFloat32Uint64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Uint64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Uint64Value %T %v\n", ptr, err))
			}
		}

	case *map[float32]float64:
		var varType = (*mapFloat32Float64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Float64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Float64Value %T %v\n", ptr, err))
			}
		}

	case *map[float32]float32:
		var varType = (*mapFloat32Float32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Float32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32Float32Value %T %v\n", ptr, err))
			}
		}

	case *map[float32]bool:
		var varType = (*mapFloat32BoolValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32BoolValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32BoolValue %T %v\n", ptr, err))
			}
		}

	case *map[float32]string:
		var varType = (*mapFloat32StringValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32StringValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapFloat32StringValue %T %v\n", ptr, err))
			}
		}

	case *map[bool]time.Duration:
		var varType = (*mapBoolDurationValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolDurationValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolDurationValue %T %v\n", ptr, err))
			}
		}

	case *map[bool]int:
		var varType = (*mapBoolIntValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolIntValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolIntValue %T %v\n", ptr, err))
			}
		}

	case *map[bool]int8:
		var varType = (*mapBoolInt8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolInt8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolInt8Value %T %v\n", ptr, err))
			}
		}

	case *map[bool]int16:
		var varType = (*mapBoolInt16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolInt16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolInt16Value %T %v\n", ptr, err))
			}
		}

	case *map[bool]int32:
		var varType = (*mapBoolInt32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolInt32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolInt32Value %T %v\n", ptr, err))
			}
		}

	case *map[bool]int64:
		var varType = (*mapBoolInt64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolInt64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolInt64Value %T %v\n", ptr, err))
			}
		}

	case *map[bool]uint:
		var varType = (*mapBoolUintValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolUintValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolUintValue %T %v\n", ptr, err))
			}
		}

	case *map[bool]uint8:
		var varType = (*mapBoolUint8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolUint8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolUint8Value %T %v\n", ptr, err))
			}
		}

	case *map[bool]uint16:
		var varType = (*mapBoolUint16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolUint16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolUint16Value %T %v\n", ptr, err))
			}
		}

	case *map[bool]uint32:
		var varType = (*mapBoolUint32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolUint32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolUint32Value %T %v\n", ptr, err))
			}
		}

	case *map[bool]uint64:
		var varType = (*mapBoolUint64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolUint64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolUint64Value %T %v\n", ptr, err))
			}
		}

	case *map[bool]float64:
		var varType = (*mapBoolFloat64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolFloat64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolFloat64Value %T %v\n", ptr, err))
			}
		}

	case *map[bool]float32:
		var varType = (*mapBoolFloat32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolFloat32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolFloat32Value %T %v\n", ptr, err))
			}
		}

	case *map[bool]bool:
		var varType = (*mapBoolBoolValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolBoolValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolBoolValue %T %v\n", ptr, err))
			}
		}

	case *map[bool]string:
		var varType = (*mapBoolStringValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolStringValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapBoolStringValue %T %v\n", ptr, err))
			}
		}

	case *map[string]time.Duration:
		var varType = (*mapStringDurationValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapStringDurationValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapStringDurationValue %T %v\n", ptr, err))
			}
		}

	case *map[string]int:
		var varType = (*mapStringIntValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapStringIntValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapStringIntValue %T %v\n", ptr, err))
			}
		}

	case *map[string]int8:
		var varType = (*mapStringInt8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapStringInt8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapStringInt8Value %T %v\n", ptr, err))
			}
		}

	case *map[string]int16:
		var varType = (*mapStringInt16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapStringInt16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapStringInt16Value %T %v\n", ptr, err))
			}
		}

	case *map[string]int32:
		var varType = (*mapStringInt32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapStringInt32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapStringInt32Value %T %v\n", ptr, err))
			}
		}

	case *map[string]int64:
		var varType = (*mapStringInt64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapStringInt64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapStringInt64Value %T %v\n", ptr, err))
			}
		}

	case *map[string]uint:
		var varType = (*mapStringUintValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapStringUintValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapStringUintValue %T %v\n", ptr, err))
			}
		}

	case *map[string]uint8:
		var varType = (*mapStringUint8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapStringUint8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapStringUint8Value %T %v\n", ptr, err))
			}
		}

	case *map[string]uint16:
		var varType = (*mapStringUint16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapStringUint16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapStringUint16Value %T %v\n", ptr, err))
			}
		}

	case *map[string]uint32:
		var varType = (*mapStringUint32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapStringUint32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapStringUint32Value %T %v\n", ptr, err))
			}
		}

	case *map[string]uint64:
		var varType = (*mapStringUint64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapStringUint64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapStringUint64Value %T %v\n", ptr, err))
			}
		}

	case *map[string]float64:
		var varType = (*mapStringFloat64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapStringFloat64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapStringFloat64Value %T %v\n", ptr, err))
			}
		}

	case *map[string]float32:
		var varType = (*mapStringFloat32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapStringFloat32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapStringFloat32Value %T %v\n", ptr, err))
			}
		}

	case *map[string]bool:
		var varType = (*mapStringBoolValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapStringBoolValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapStringBoolValue %T %v\n", ptr, err))
			}
		}

	case *map[string]string:
		var varType = (*mapStringStringValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting mapStringStringValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting mapStringStringValue %T %v\n", ptr, err))
			}
		}

	case *time.Duration:
		var varType = (*durationValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting durationValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting durationValue %T %v\n", ptr, err))
			}
		}

	case *int:
		var varType = (*intValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting intValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting intValue %T %v\n", ptr, err))
			}
		}

	case *int8:
		var varType = (*int8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting int8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting int8Value %T %v\n", ptr, err))
			}
		}

	case *int16:
		var varType = (*int16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting int16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting int16Value %T %v\n", ptr, err))
			}
		}

	case *int32:
		var varType = (*int32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting int32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting int32Value %T %v\n", ptr, err))
			}
		}

	case *int64:
		var varType = (*int64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting int64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting int64Value %T %v\n", ptr, err))
			}
		}

	case *uint:
		var varType = (*uintValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting uintValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting uintValue %T %v\n", ptr, err))
			}
		}

	case *uint8:
		var varType = (*uint8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting uint8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting uint8Value %T %v\n", ptr, err))
			}
		}

	case *uint16:
		var varType = (*uint16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting uint16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting uint16Value %T %v\n", ptr, err))
			}
		}

	case *uint32:
		var varType = (*uint32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting uint32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting uint32Value %T %v\n", ptr, err))
			}
		}

	case *uint64:
		var varType = (*uint64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting uint64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting uint64Value %T %v\n", ptr, err))
			}
		}

	case *float64:
		var varType = (*float64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting float64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting float64Value %T %v\n", ptr, err))
			}
		}

	case *float32:
		var varType = (*float32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting float32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting float32Value %T %v\n", ptr, err))
			}
		}

	case *bool:
		var varType = (*boolValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting boolValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting boolValue %T %v\n", ptr, err))
			}
		}

	case *string:
		var varType = (*stringValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting stringValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting stringValue %T %v\n", ptr, err))
			}
		}

	case *[]time.Duration:
		var varType = (*sliceDurationValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting sliceDurationValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting sliceDurationValue %T %v\n", ptr, err))
			}
		}

	case *[]int:
		var varType = (*sliceIntValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting sliceIntValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting sliceIntValue %T %v\n", ptr, err))
			}
		}

	case *[]int8:
		var varType = (*sliceInt8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting sliceInt8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting sliceInt8Value %T %v\n", ptr, err))
			}
		}

	case *[]int16:
		var varType = (*sliceInt16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting sliceInt16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting sliceInt16Value %T %v\n", ptr, err))
			}
		}

	case *[]int32:
		var varType = (*sliceInt32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting sliceInt32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting sliceInt32Value %T %v\n", ptr, err))
			}
		}

	case *[]int64:
		var varType = (*sliceInt64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting sliceInt64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting sliceInt64Value %T %v\n", ptr, err))
			}
		}

	case *[]uint:
		var varType = (*sliceUintValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting sliceUintValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting sliceUintValue %T %v\n", ptr, err))
			}
		}

	case *[]uint8:
		var varType = (*sliceUint8Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting sliceUint8Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting sliceUint8Value %T %v\n", ptr, err))
			}
		}

	case *[]uint16:
		var varType = (*sliceUint16Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting sliceUint16Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting sliceUint16Value %T %v\n", ptr, err))
			}
		}

	case *[]uint32:
		var varType = (*sliceUint32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting sliceUint32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting sliceUint32Value %T %v\n", ptr, err))
			}
		}

	case *[]uint64:
		var varType = (*sliceUint64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting sliceUint64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting sliceUint64Value %T %v\n", ptr, err))
			}
		}

	case *[]float64:
		var varType = (*sliceFloat64Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting sliceFloat64Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting sliceFloat64Value %T %v\n", ptr, err))
			}
		}

	case *[]float32:
		var varType = (*sliceFloat32Value)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting sliceFloat32Value %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting sliceFloat32Value %T %v\n", ptr, err))
			}
		}

	case *[]bool:
		var varType = (*sliceBoolValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting sliceBoolValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting sliceBoolValue %T %v\n", ptr, err))
			}
		}

	case *[]string:
		var varType = (*sliceStringValue)(ptr)
		if len(defaultValue) > 0 {
			if err := varType.Set(defaultValue); err != nil {
				panic(fmt.Sprintf("Error setting sliceStringValue %T %v\n", ptr, err))
			}
		}
		CommandLine.Var(varType, name, usage, required, isset)
		if len(override) > 0 {
			if err := varType.Set(override); err != nil {
				panic(fmt.Sprintf("Error setting sliceStringValue %T %v\n", ptr, err))
			}
		}

	default:
		// panic("Unhandled Type %T\n", ptr)
		panic(fmt.Sprintf("Unhandled Type %v %T\n", ptr, ptr))
	}
}
