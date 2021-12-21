//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package matrixpilot

import (
	"errors"
)

// Yaw behaviour during orbit flight.
type ORBIT_YAW_BEHAVIOUR int

const (
	// Vehicle front points to the center (default).
	ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TO_CIRCLE_CENTER ORBIT_YAW_BEHAVIOUR = 0
	// Vehicle front holds heading when message received.
	ORBIT_YAW_BEHAVIOUR_HOLD_INITIAL_HEADING ORBIT_YAW_BEHAVIOUR = 1
	// Yaw uncontrolled.
	ORBIT_YAW_BEHAVIOUR_UNCONTROLLED ORBIT_YAW_BEHAVIOUR = 2
	// Vehicle front follows flight path (tangential to circle).
	ORBIT_YAW_BEHAVIOUR_HOLD_FRONT_TANGENT_TO_CIRCLE ORBIT_YAW_BEHAVIOUR = 3
	// Yaw controlled by RC input.
	ORBIT_YAW_BEHAVIOUR_RC_CONTROLLED ORBIT_YAW_BEHAVIOUR = 4
)

var labels_ORBIT_YAW_BEHAVIOUR = map[ORBIT_YAW_BEHAVIOUR]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ORBIT_YAW_BEHAVIOUR) MarshalText() ([]byte, error) {
	if l, ok := labels_ORBIT_YAW_BEHAVIOUR[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_ORBIT_YAW_BEHAVIOUR = map[string]ORBIT_YAW_BEHAVIOUR{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ORBIT_YAW_BEHAVIOUR) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_ORBIT_YAW_BEHAVIOUR[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e ORBIT_YAW_BEHAVIOUR) String() string {
	if l, ok := labels_ORBIT_YAW_BEHAVIOUR[e]; ok {
		return l
	}
	return "invalid value"
}
