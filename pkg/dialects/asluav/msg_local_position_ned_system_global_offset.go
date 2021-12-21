//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

// The offset in X, Y, Z and yaw between the LOCAL_POSITION_NED messages of MAV X and the global coordinate frame in NED coordinates. Coordinate frame is right-handed, Z-axis down (aeronautical frame, NED / north-east-down convention)
type MessageLocalPositionNedSystemGlobalOffset struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// X Position
	X float32
	// Y Position
	Y float32
	// Z Position
	Z float32
	// Roll
	Roll float32
	// Pitch
	Pitch float32
	// Yaw
	Yaw float32
}

// GetID implements the msg.Message interface.
func (*MessageLocalPositionNedSystemGlobalOffset) GetID() uint32 {
	return 89
}
