//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// The attitude in the aeronautical frame (right-handed, Z-down, X-front, Y-right).
type MessageAttitude struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Roll angle (-pi..+pi)
	Roll float32
	// Pitch angle (-pi..+pi)
	Pitch float32
	// Yaw angle (-pi..+pi)
	Yaw float32
	// Roll angular speed
	Rollspeed float32
	// Pitch angular speed
	Pitchspeed float32
	// Yaw angular speed
	Yawspeed float32
}

// GetID implements the msg.Message interface.
func (*MessageAttitude) GetID() uint32 {
	return 30
}
