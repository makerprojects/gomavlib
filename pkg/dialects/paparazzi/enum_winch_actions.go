//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package paparazzi

import (
	"errors"
)

// Winch actions.
type WINCH_ACTIONS int

const (
	// Relax winch.
	WINCH_RELAXED WINCH_ACTIONS = 0
	// Wind or unwind specified length of cable, optionally using specified rate.
	WINCH_RELATIVE_LENGTH_CONTROL WINCH_ACTIONS = 1
	// Wind or unwind cable at specified rate.
	WINCH_RATE_CONTROL WINCH_ACTIONS = 2
)

var labels_WINCH_ACTIONS = map[WINCH_ACTIONS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e WINCH_ACTIONS) MarshalText() ([]byte, error) {
	if l, ok := labels_WINCH_ACTIONS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_WINCH_ACTIONS = map[string]WINCH_ACTIONS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *WINCH_ACTIONS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_WINCH_ACTIONS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e WINCH_ACTIONS) String() string {
	if l, ok := labels_WINCH_ACTIONS[e]; ok {
		return l
	}
	return "invalid value"
}
