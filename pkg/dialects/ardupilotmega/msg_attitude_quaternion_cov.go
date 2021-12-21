//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ardupilotmega

// The attitude in the aeronautical frame (right-handed, Z-down, X-front, Y-right), expressed as quaternion. Quaternion order is w, x, y, z and a zero rotation would be expressed as (1 0 0 0).
type MessageAttitudeQuaternionCov struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Quaternion components, w, x, y, z (1 0 0 0 is the null-rotation)
	Q [4]float32
	// Roll angular speed
	Rollspeed float32
	// Pitch angular speed
	Pitchspeed float32
	// Yaw angular speed
	Yawspeed float32
	// Row-major representation of a 3x3 attitude covariance matrix (states: roll, pitch, yaw; first three entries are the first ROW, next three entries are the second row, etc.). If unknown, assign NaN value to first element in the array.
	Covariance [9]float32
}

// GetID implements the msg.Message interface.
func (*MessageAttitudeQuaternionCov) GetID() uint32 {
	return 61
}
