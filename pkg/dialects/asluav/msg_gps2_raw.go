//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

// Second GPS data.
type MessageGps2Raw struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// GPS fix type.
	FixType GPS_FIX_TYPE `mavenum:"uint8"`
	// Latitude (WGS84)
	Lat int32
	// Longitude (WGS84)
	Lon int32
	// Altitude (MSL). Positive for up.
	Alt int32
	// GPS HDOP horizontal dilution of position (unitless * 100). If unknown, set to: UINT16_MAX
	Eph uint16
	// GPS VDOP vertical dilution of position (unitless * 100). If unknown, set to: UINT16_MAX
	Epv uint16
	// GPS ground speed. If unknown, set to: UINT16_MAX
	Vel uint16
	// Course over ground (NOT heading, but direction of movement): 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
	Cog uint16
	// Number of satellites visible. If unknown, set to UINT8_MAX
	SatellitesVisible uint8
	// Number of DGPS satellites
	DgpsNumch uint8
	// Age of DGPS info
	DgpsAge uint32
	// Yaw in earth frame from north. Use 0 if this GPS does not provide yaw. Use UINT16_MAX if this GPS is configured to provide yaw and is currently unable to provide it. Use 36000 for north.
	Yaw uint16 `mavext:"true"`
	// Altitude (above WGS84, EGM96 ellipsoid). Positive for up.
	AltEllipsoid int32 `mavext:"true"`
	// Position uncertainty.
	HAcc uint32 `mavext:"true"`
	// Altitude uncertainty.
	VAcc uint32 `mavext:"true"`
	// Speed uncertainty.
	VelAcc uint32 `mavext:"true"`
	// Heading / track uncertainty
	HdgAcc uint32 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageGps2Raw) GetID() uint32 {
	return 124
}
