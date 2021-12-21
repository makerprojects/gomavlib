//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package matrixpilot

import (
	"errors"
)

// Action required when performing CMD_PREFLIGHT_STORAGE
type MAV_PREFLIGHT_STORAGE_ACTION int

const (
	// Read all parameters from storage
	MAV_PFS_CMD_READ_ALL MAV_PREFLIGHT_STORAGE_ACTION = 0
	// Write all parameters to storage
	MAV_PFS_CMD_WRITE_ALL MAV_PREFLIGHT_STORAGE_ACTION = 1
	// Clear all  parameters in storage
	MAV_PFS_CMD_CLEAR_ALL MAV_PREFLIGHT_STORAGE_ACTION = 2
	// Read specific parameters from storage
	MAV_PFS_CMD_READ_SPECIFIC MAV_PREFLIGHT_STORAGE_ACTION = 3
	// Write specific parameters to storage
	MAV_PFS_CMD_WRITE_SPECIFIC MAV_PREFLIGHT_STORAGE_ACTION = 4
	// Clear specific parameters in storage
	MAV_PFS_CMD_CLEAR_SPECIFIC MAV_PREFLIGHT_STORAGE_ACTION = 5
	// do nothing
	MAV_PFS_CMD_DO_NOTHING MAV_PREFLIGHT_STORAGE_ACTION = 6
)

var labels_MAV_PREFLIGHT_STORAGE_ACTION = map[MAV_PREFLIGHT_STORAGE_ACTION]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_PREFLIGHT_STORAGE_ACTION) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_PREFLIGHT_STORAGE_ACTION[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_PREFLIGHT_STORAGE_ACTION = map[string]MAV_PREFLIGHT_STORAGE_ACTION{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_PREFLIGHT_STORAGE_ACTION) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_PREFLIGHT_STORAGE_ACTION[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_PREFLIGHT_STORAGE_ACTION) String() string {
	if l, ok := labels_MAV_PREFLIGHT_STORAGE_ACTION[e]; ok {
		return l
	}
	return "invalid value"
}
