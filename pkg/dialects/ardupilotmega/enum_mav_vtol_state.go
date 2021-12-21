//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

// Enumeration of VTOL states
type MAV_VTOL_STATE int

const (
	// MAV is not configured as VTOL
	MAV_VTOL_STATE_UNDEFINED MAV_VTOL_STATE = 0
	// VTOL is in transition from multicopter to fixed-wing
	MAV_VTOL_STATE_TRANSITION_TO_FW MAV_VTOL_STATE = 1
	// VTOL is in transition from fixed-wing to multicopter
	MAV_VTOL_STATE_TRANSITION_TO_MC MAV_VTOL_STATE = 2
	// VTOL is in multicopter state
	MAV_VTOL_STATE_MC MAV_VTOL_STATE = 3
	// VTOL is in fixed-wing state
	MAV_VTOL_STATE_FW MAV_VTOL_STATE = 4
)

var labels_MAV_VTOL_STATE = map[MAV_VTOL_STATE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_VTOL_STATE) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_VTOL_STATE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_VTOL_STATE = map[string]MAV_VTOL_STATE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_VTOL_STATE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_VTOL_STATE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_VTOL_STATE) String() string {
	if l, ok := labels_MAV_VTOL_STATE[e]; ok {
		return l
	}
	return "invalid value"
}
