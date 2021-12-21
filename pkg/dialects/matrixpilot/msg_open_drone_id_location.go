//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// Data for filling the OpenDroneID Location message. The float data types are 32-bit IEEE 754. The Location message provides the location, altitude, direction and speed of the aircraft.
type MessageOpenDroneIdLocation struct {
	// System ID (0 for broadcast).
	TargetSystem uint8
	// Component ID (0 for broadcast).
	TargetComponent uint8
	// Only used for drone ID data received from other UAs. See detailed description at https://mavlink.io/en/services/opendroneid.html.
	IdOrMac [20]uint8
	// Indicates whether the unmanned aircraft is on the ground or in the air.
	Status MAV_ODID_STATUS `mavenum:"uint8"`
	// Direction over ground (not heading, but direction of movement) measured clockwise from true North: 0 - 35999 centi-degrees. If unknown: 36100 centi-degrees.
	Direction uint16
	// Ground speed. Positive only. If unknown: 25500 cm/s. If speed is larger than 25425 cm/s, use 25425 cm/s.
	SpeedHorizontal uint16
	// The vertical speed. Up is positive. If unknown: 6300 cm/s. If speed is larger than 6200 cm/s, use 6200 cm/s. If lower than -6200 cm/s, use -6200 cm/s.
	SpeedVertical int16
	// Current latitude of the unmanned aircraft. If unknown: 0 (both Lat/Lon).
	Latitude int32
	// Current longitude of the unmanned aircraft. If unknown: 0 (both Lat/Lon).
	Longitude int32
	// The altitude calculated from the barometric pressue. Reference is against 29.92inHg or 1013.2mb. If unknown: -1000 m.
	AltitudeBarometric float32
	// The geodetic altitude as defined by WGS84. If unknown: -1000 m.
	AltitudeGeodetic float32
	// Indicates the reference point for the height field.
	HeightReference MAV_ODID_HEIGHT_REF `mavenum:"uint8"`
	// The current height of the unmanned aircraft above the take-off location or the ground as indicated by height_reference. If unknown: -1000 m.
	Height float32
	// The accuracy of the horizontal position.
	HorizontalAccuracy MAV_ODID_HOR_ACC `mavenum:"uint8"`
	// The accuracy of the vertical position.
	VerticalAccuracy MAV_ODID_VER_ACC `mavenum:"uint8"`
	// The accuracy of the barometric altitude.
	BarometerAccuracy MAV_ODID_VER_ACC `mavenum:"uint8"`
	// The accuracy of the horizontal and vertical speed.
	SpeedAccuracy MAV_ODID_SPEED_ACC `mavenum:"uint8"`
	// Seconds after the full hour with reference to UTC time. Typically the GPS outputs a time-of-week value in milliseconds. First convert that to UTC and then convert for this field using ((float) (time_week_ms % (60*60*1000))) / 1000. If unknown: 0xFFFF.
	Timestamp float32
	// The accuracy of the timestamps.
	TimestampAccuracy MAV_ODID_TIME_ACC `mavenum:"uint8"`
}

// GetID implements the msg.Message interface.
func (*MessageOpenDroneIdLocation) GetID() uint32 {
	return 12901
}
