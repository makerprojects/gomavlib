//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

// Sent from simulation to autopilot, avoids in contrast to HIL_STATE singularities. This packet is useful for high throughput applications such as hardware in the loop simulations.
type MessageHilStateQuaternion struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Vehicle attitude expressed as normalized quaternion in w, x, y, z order (with 1 0 0 0 being the null-rotation)
	AttitudeQuaternion [4]float32
	// Body frame roll / phi angular speed
	Rollspeed float32
	// Body frame pitch / theta angular speed
	Pitchspeed float32
	// Body frame yaw / psi angular speed
	Yawspeed float32
	// Latitude
	Lat int32
	// Longitude
	Lon int32
	// Altitude
	Alt int32
	// Ground X Speed (Latitude)
	Vx int16
	// Ground Y Speed (Longitude)
	Vy int16
	// Ground Z Speed (Altitude)
	Vz int16
	// Indicated airspeed
	IndAirspeed uint16
	// True airspeed
	TrueAirspeed uint16
	// X acceleration
	Xacc int16
	// Y acceleration
	Yacc int16
	// Z acceleration
	Zacc int16
}

// GetID implements the msg.Message interface.
func (*MessageHilStateQuaternion) GetID() uint32 {
	return 115
}
