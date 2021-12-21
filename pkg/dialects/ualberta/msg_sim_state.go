//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

// Status of simulation environment, if used
type MessageSimState struct {
	// True attitude quaternion component 1, w (1 in null-rotation)
	Q1 float32
	// True attitude quaternion component 2, x (0 in null-rotation)
	Q2 float32
	// True attitude quaternion component 3, y (0 in null-rotation)
	Q3 float32
	// True attitude quaternion component 4, z (0 in null-rotation)
	Q4 float32
	// Attitude roll expressed as Euler angles, not recommended except for human-readable outputs
	Roll float32
	// Attitude pitch expressed as Euler angles, not recommended except for human-readable outputs
	Pitch float32
	// Attitude yaw expressed as Euler angles, not recommended except for human-readable outputs
	Yaw float32
	// X acceleration
	Xacc float32
	// Y acceleration
	Yacc float32
	// Z acceleration
	Zacc float32
	// Angular speed around X axis
	Xgyro float32
	// Angular speed around Y axis
	Ygyro float32
	// Angular speed around Z axis
	Zgyro float32
	// Latitude
	Lat float32
	// Longitude
	Lon float32
	// Altitude
	Alt float32
	// Horizontal position standard deviation
	StdDevHorz float32
	// Vertical position standard deviation
	StdDevVert float32
	// True velocity in north direction in earth-fixed NED frame
	Vn float32
	// True velocity in east direction in earth-fixed NED frame
	Ve float32
	// True velocity in down direction in earth-fixed NED frame
	Vd float32
}

// GetID implements the msg.Message interface.
func (*MessageSimState) GetID() uint32 {
	return 108
}
