//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package matrixpilot

import (
	"errors"
)

type MAV_ARM_AUTH_DENIED_REASON int

const (
	// Not a specific reason
	MAV_ARM_AUTH_DENIED_REASON_GENERIC MAV_ARM_AUTH_DENIED_REASON = 0
	// Authorizer will send the error as string to GCS
	MAV_ARM_AUTH_DENIED_REASON_NONE MAV_ARM_AUTH_DENIED_REASON = 1
	// At least one waypoint have a invalid value
	MAV_ARM_AUTH_DENIED_REASON_INVALID_WAYPOINT MAV_ARM_AUTH_DENIED_REASON = 2
	// Timeout in the authorizer process(in case it depends on network)
	MAV_ARM_AUTH_DENIED_REASON_TIMEOUT MAV_ARM_AUTH_DENIED_REASON = 3
	// Airspace of the mission in use by another vehicle, second result parameter can have the waypoint id that caused it to be denied.
	MAV_ARM_AUTH_DENIED_REASON_AIRSPACE_IN_USE MAV_ARM_AUTH_DENIED_REASON = 4
	// Weather is not good to fly
	MAV_ARM_AUTH_DENIED_REASON_BAD_WEATHER MAV_ARM_AUTH_DENIED_REASON = 5
)

var labels_MAV_ARM_AUTH_DENIED_REASON = map[MAV_ARM_AUTH_DENIED_REASON]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_ARM_AUTH_DENIED_REASON) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_ARM_AUTH_DENIED_REASON[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_ARM_AUTH_DENIED_REASON = map[string]MAV_ARM_AUTH_DENIED_REASON{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_ARM_AUTH_DENIED_REASON) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_ARM_AUTH_DENIED_REASON[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_ARM_AUTH_DENIED_REASON) String() string {
	if l, ok := labels_MAV_ARM_AUTH_DENIED_REASON[e]; ok {
		return l
	}
	return "invalid value"
}
