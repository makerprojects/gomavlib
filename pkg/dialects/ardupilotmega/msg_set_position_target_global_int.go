//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ardupilotmega

// Sets a desired vehicle position, velocity, and/or acceleration in a global coordinate system (WGS84). Used by an external controller to command the vehicle (manual controller or other system).
type MessageSetPositionTargetGlobalInt struct {
	// Timestamp (time since system boot). The rationale for the timestamp in the setpoint is to allow the system to compensate for the transport delay of the setpoint. This allows the system to compensate processing latency.
	TimeBootMs uint32
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Valid options are: MAV_FRAME_GLOBAL_INT = 5, MAV_FRAME_GLOBAL_RELATIVE_ALT_INT = 6, MAV_FRAME_GLOBAL_TERRAIN_ALT_INT = 11
	CoordinateFrame MAV_FRAME `mavenum:"uint8"`
	// Bitmap to indicate which dimensions should be ignored by the vehicle.
	TypeMask POSITION_TARGET_TYPEMASK `mavenum:"uint16"`
	// X Position in WGS84 frame
	LatInt int32
	// Y Position in WGS84 frame
	LonInt int32
	// Altitude (MSL, Relative to home, or AGL - depending on frame)
	Alt float32
	// X velocity in NED frame
	Vx float32
	// Y velocity in NED frame
	Vy float32
	// Z velocity in NED frame
	Vz float32
	// X acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afx float32
	// Y acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afy float32
	// Z acceleration or force (if bit 10 of type_mask is set) in NED frame in meter / s^2 or N
	Afz float32
	// yaw setpoint
	Yaw float32
	// yaw rate setpoint
	YawRate float32
}

// GetID implements the msg.Message interface.
func (*MessageSetPositionTargetGlobalInt) GetID() uint32 {
	return 86
}
