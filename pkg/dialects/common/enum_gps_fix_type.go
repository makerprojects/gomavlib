//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package common

import (
	"errors"
)

// Type of GPS fix
type GPS_FIX_TYPE int

const (
	// No GPS connected
	GPS_FIX_TYPE_NO_GPS GPS_FIX_TYPE = 0
	// No position information, GPS is connected
	GPS_FIX_TYPE_NO_FIX GPS_FIX_TYPE = 1
	// 2D position
	GPS_FIX_TYPE_2D_FIX GPS_FIX_TYPE = 2
	// 3D position
	GPS_FIX_TYPE_3D_FIX GPS_FIX_TYPE = 3
	// DGPS/SBAS aided 3D position
	GPS_FIX_TYPE_DGPS GPS_FIX_TYPE = 4
	// RTK float, 3D position
	GPS_FIX_TYPE_RTK_FLOAT GPS_FIX_TYPE = 5
	// RTK Fixed, 3D position
	GPS_FIX_TYPE_RTK_FIXED GPS_FIX_TYPE = 6
	// Static fixed, typically used for base stations
	GPS_FIX_TYPE_STATIC GPS_FIX_TYPE = 7
	// PPP, 3D position.
	GPS_FIX_TYPE_PPP GPS_FIX_TYPE = 8
)

var labels_GPS_FIX_TYPE = map[GPS_FIX_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GPS_FIX_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_GPS_FIX_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GPS_FIX_TYPE = map[string]GPS_FIX_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GPS_FIX_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GPS_FIX_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GPS_FIX_TYPE) String() string {
	if l, ok := labels_GPS_FIX_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
