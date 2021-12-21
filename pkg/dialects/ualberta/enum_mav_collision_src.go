//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ualberta

import (
	"errors"
)

// Source of information about this collision.
type MAV_COLLISION_SRC int

const (
	// ID field references ADSB_VEHICLE packets
	MAV_COLLISION_SRC_ADSB MAV_COLLISION_SRC = 0
	// ID field references MAVLink SRC ID
	MAV_COLLISION_SRC_MAVLINK_GPS_GLOBAL_INT MAV_COLLISION_SRC = 1
)

var labels_MAV_COLLISION_SRC = map[MAV_COLLISION_SRC]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_COLLISION_SRC) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_COLLISION_SRC[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_COLLISION_SRC = map[string]MAV_COLLISION_SRC{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_COLLISION_SRC) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_COLLISION_SRC[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_COLLISION_SRC) String() string {
	if l, ok := labels_MAV_COLLISION_SRC[e]; ok {
		return l
	}
	return "invalid value"
}
