//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package standard

import (
	"errors"
)

// Specifies the datatype of a MAVLink extended parameter.
type MAV_PARAM_EXT_TYPE int

const (
	// 8-bit unsigned integer
	MAV_PARAM_EXT_TYPE_UINT8 MAV_PARAM_EXT_TYPE = 1
	// 8-bit signed integer
	MAV_PARAM_EXT_TYPE_INT8 MAV_PARAM_EXT_TYPE = 2
	// 16-bit unsigned integer
	MAV_PARAM_EXT_TYPE_UINT16 MAV_PARAM_EXT_TYPE = 3
	// 16-bit signed integer
	MAV_PARAM_EXT_TYPE_INT16 MAV_PARAM_EXT_TYPE = 4
	// 32-bit unsigned integer
	MAV_PARAM_EXT_TYPE_UINT32 MAV_PARAM_EXT_TYPE = 5
	// 32-bit signed integer
	MAV_PARAM_EXT_TYPE_INT32 MAV_PARAM_EXT_TYPE = 6
	// 64-bit unsigned integer
	MAV_PARAM_EXT_TYPE_UINT64 MAV_PARAM_EXT_TYPE = 7
	// 64-bit signed integer
	MAV_PARAM_EXT_TYPE_INT64 MAV_PARAM_EXT_TYPE = 8
	// 32-bit floating-point
	MAV_PARAM_EXT_TYPE_REAL32 MAV_PARAM_EXT_TYPE = 9
	// 64-bit floating-point
	MAV_PARAM_EXT_TYPE_REAL64 MAV_PARAM_EXT_TYPE = 10
	// Custom Type
	MAV_PARAM_EXT_TYPE_CUSTOM MAV_PARAM_EXT_TYPE = 11
)

var labels_MAV_PARAM_EXT_TYPE = map[MAV_PARAM_EXT_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_PARAM_EXT_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_PARAM_EXT_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_PARAM_EXT_TYPE = map[string]MAV_PARAM_EXT_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_PARAM_EXT_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_PARAM_EXT_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_PARAM_EXT_TYPE) String() string {
	if l, ok := labels_MAV_PARAM_EXT_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
