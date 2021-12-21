//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

type GOPRO_HEARTBEAT_STATUS int

const (
	// No GoPro connected.
	GOPRO_HEARTBEAT_STATUS_DISCONNECTED GOPRO_HEARTBEAT_STATUS = 0
	// The detected GoPro is not HeroBus compatible.
	GOPRO_HEARTBEAT_STATUS_INCOMPATIBLE GOPRO_HEARTBEAT_STATUS = 1
	// A HeroBus compatible GoPro is connected.
	GOPRO_HEARTBEAT_STATUS_CONNECTED GOPRO_HEARTBEAT_STATUS = 2
	// An unrecoverable error was encountered with the connected GoPro, it may require a power cycle.
	GOPRO_HEARTBEAT_STATUS_ERROR GOPRO_HEARTBEAT_STATUS = 3
)

var labels_GOPRO_HEARTBEAT_STATUS = map[GOPRO_HEARTBEAT_STATUS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GOPRO_HEARTBEAT_STATUS) MarshalText() ([]byte, error) {
	if l, ok := labels_GOPRO_HEARTBEAT_STATUS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GOPRO_HEARTBEAT_STATUS = map[string]GOPRO_HEARTBEAT_STATUS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GOPRO_HEARTBEAT_STATUS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GOPRO_HEARTBEAT_STATUS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GOPRO_HEARTBEAT_STATUS) String() string {
	if l, ok := labels_GOPRO_HEARTBEAT_STATUS[e]; ok {
		return l
	}
	return "invalid value"
}
