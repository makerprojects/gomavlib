//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package matrixpilot

import (
	"errors"
)

// Result from PARAM_EXT_SET message (or a PARAM_SET within a transaction).
type PARAM_ACK int

const (
	// Parameter value ACCEPTED and SET
	PARAM_ACK_ACCEPTED PARAM_ACK = 0
	// Parameter value UNKNOWN/UNSUPPORTED
	PARAM_ACK_VALUE_UNSUPPORTED PARAM_ACK = 1
	// Parameter failed to set
	PARAM_ACK_FAILED PARAM_ACK = 2
	// Parameter value received but not yet set/accepted. A subsequent PARAM_ACK_TRANSACTION or PARAM_EXT_ACK with the final result will follow once operation is completed. This is returned immediately for parameters that take longer to set, indicating taht the the parameter was recieved and does not need to be resent.
	PARAM_ACK_IN_PROGRESS PARAM_ACK = 3
)

var labels_PARAM_ACK = map[PARAM_ACK]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e PARAM_ACK) MarshalText() ([]byte, error) {
	if l, ok := labels_PARAM_ACK[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_PARAM_ACK = map[string]PARAM_ACK{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *PARAM_ACK) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_PARAM_ACK[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e PARAM_ACK) String() string {
	if l, ok := labels_PARAM_ACK[e]; ok {
		return l
	}
	return "invalid value"
}
