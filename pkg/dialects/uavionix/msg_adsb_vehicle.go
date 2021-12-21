//autogenerated:yes
//nolint:revive,misspell,govet,lll
package uavionix

// The location and information of an ADSB vehicle
type MessageAdsbVehicle struct {
	// ICAO address
	IcaoAddress uint32 `mavname:"ICAO_address"`
	// Latitude
	Lat int32
	// Longitude
	Lon int32
	// ADSB altitude type.
	AltitudeType ADSB_ALTITUDE_TYPE `mavenum:"uint8"`
	// Altitude(ASL)
	Altitude int32
	// Course over ground
	Heading uint16
	// The horizontal velocity
	HorVelocity uint16
	// The vertical velocity. Positive is up
	VerVelocity int16
	// The callsign, 8+null
	Callsign string `mavlen:"9"`
	// ADSB emitter type.
	EmitterType ADSB_EMITTER_TYPE `mavenum:"uint8"`
	// Time since last communication in seconds
	Tslc uint8
	// Bitmap to indicate various statuses including valid data fields
	Flags ADSB_FLAGS `mavenum:"uint16"`
	// Squawk code
	Squawk uint16
}

// GetID implements the msg.Message interface.
func (*MessageAdsbVehicle) GetID() uint32 {
	return 246
}
