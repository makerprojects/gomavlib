//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

// GPS lataral offset encoding
type UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT int

const (
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_NO_DATA  UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = 0
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_2M  UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = 1
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_4M  UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = 2
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_LEFT_6M  UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = 3
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_0M UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = 4
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_2M UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = 5
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_4M UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = 6
	UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT_RIGHT_6M UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = 7
)

var labels_UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = map[UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT) MarshalText() ([]byte, error) {
	if l, ok := labels_UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT = map[string]UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT) String() string {
	if l, ok := labels_UAVIONIX_ADSB_OUT_CFG_GPS_OFFSET_LAT[e]; ok {
		return l
	}
	return "invalid value"
}
