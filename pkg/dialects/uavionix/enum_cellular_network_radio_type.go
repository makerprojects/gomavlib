//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package uavionix

import (
	"errors"
)

// Cellular network radio type
type CELLULAR_NETWORK_RADIO_TYPE int

const (
	CELLULAR_NETWORK_RADIO_TYPE_NONE  CELLULAR_NETWORK_RADIO_TYPE = 0
	CELLULAR_NETWORK_RADIO_TYPE_GSM   CELLULAR_NETWORK_RADIO_TYPE = 1
	CELLULAR_NETWORK_RADIO_TYPE_CDMA  CELLULAR_NETWORK_RADIO_TYPE = 2
	CELLULAR_NETWORK_RADIO_TYPE_WCDMA CELLULAR_NETWORK_RADIO_TYPE = 3
	CELLULAR_NETWORK_RADIO_TYPE_LTE   CELLULAR_NETWORK_RADIO_TYPE = 4
)

var labels_CELLULAR_NETWORK_RADIO_TYPE = map[CELLULAR_NETWORK_RADIO_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e CELLULAR_NETWORK_RADIO_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_CELLULAR_NETWORK_RADIO_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_CELLULAR_NETWORK_RADIO_TYPE = map[string]CELLULAR_NETWORK_RADIO_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *CELLULAR_NETWORK_RADIO_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_CELLULAR_NETWORK_RADIO_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e CELLULAR_NETWORK_RADIO_TYPE) String() string {
	if l, ok := labels_CELLULAR_NETWORK_RADIO_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
