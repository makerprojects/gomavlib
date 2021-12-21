//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

// Sent from simulation to autopilot. This packet is useful for high throughput applications such as hardware in the loop simulations.
type MessageHilState struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Roll angle
	Roll float32
	// Pitch angle
	Pitch float32
	// Yaw angle
	Yaw float32
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
	// X acceleration
	Xacc int16
	// Y acceleration
	Yacc int16
	// Z acceleration
	Zacc int16
}

// GetID implements the msg.Message interface.
func (*MessageHilState) GetID() uint32 {
	return 90
}
