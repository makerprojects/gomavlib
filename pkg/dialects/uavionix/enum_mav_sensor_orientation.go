//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package uavionix

import (
	"errors"
)

// Enumeration of sensor orientation, according to its rotations
type MAV_SENSOR_ORIENTATION int

const (
	// Roll: 0, Pitch: 0, Yaw: 0
	MAV_SENSOR_ROTATION_NONE MAV_SENSOR_ORIENTATION = 0
	// Roll: 0, Pitch: 0, Yaw: 45
	MAV_SENSOR_ROTATION_YAW_45 MAV_SENSOR_ORIENTATION = 1
	// Roll: 0, Pitch: 0, Yaw: 90
	MAV_SENSOR_ROTATION_YAW_90 MAV_SENSOR_ORIENTATION = 2
	// Roll: 0, Pitch: 0, Yaw: 135
	MAV_SENSOR_ROTATION_YAW_135 MAV_SENSOR_ORIENTATION = 3
	// Roll: 0, Pitch: 0, Yaw: 180
	MAV_SENSOR_ROTATION_YAW_180 MAV_SENSOR_ORIENTATION = 4
	// Roll: 0, Pitch: 0, Yaw: 225
	MAV_SENSOR_ROTATION_YAW_225 MAV_SENSOR_ORIENTATION = 5
	// Roll: 0, Pitch: 0, Yaw: 270
	MAV_SENSOR_ROTATION_YAW_270 MAV_SENSOR_ORIENTATION = 6
	// Roll: 0, Pitch: 0, Yaw: 315
	MAV_SENSOR_ROTATION_YAW_315 MAV_SENSOR_ORIENTATION = 7
	// Roll: 180, Pitch: 0, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_180 MAV_SENSOR_ORIENTATION = 8
	// Roll: 180, Pitch: 0, Yaw: 45
	MAV_SENSOR_ROTATION_ROLL_180_YAW_45 MAV_SENSOR_ORIENTATION = 9
	// Roll: 180, Pitch: 0, Yaw: 90
	MAV_SENSOR_ROTATION_ROLL_180_YAW_90 MAV_SENSOR_ORIENTATION = 10
	// Roll: 180, Pitch: 0, Yaw: 135
	MAV_SENSOR_ROTATION_ROLL_180_YAW_135 MAV_SENSOR_ORIENTATION = 11
	// Roll: 0, Pitch: 180, Yaw: 0
	MAV_SENSOR_ROTATION_PITCH_180 MAV_SENSOR_ORIENTATION = 12
	// Roll: 180, Pitch: 0, Yaw: 225
	MAV_SENSOR_ROTATION_ROLL_180_YAW_225 MAV_SENSOR_ORIENTATION = 13
	// Roll: 180, Pitch: 0, Yaw: 270
	MAV_SENSOR_ROTATION_ROLL_180_YAW_270 MAV_SENSOR_ORIENTATION = 14
	// Roll: 180, Pitch: 0, Yaw: 315
	MAV_SENSOR_ROTATION_ROLL_180_YAW_315 MAV_SENSOR_ORIENTATION = 15
	// Roll: 90, Pitch: 0, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_90 MAV_SENSOR_ORIENTATION = 16
	// Roll: 90, Pitch: 0, Yaw: 45
	MAV_SENSOR_ROTATION_ROLL_90_YAW_45 MAV_SENSOR_ORIENTATION = 17
	// Roll: 90, Pitch: 0, Yaw: 90
	MAV_SENSOR_ROTATION_ROLL_90_YAW_90 MAV_SENSOR_ORIENTATION = 18
	// Roll: 90, Pitch: 0, Yaw: 135
	MAV_SENSOR_ROTATION_ROLL_90_YAW_135 MAV_SENSOR_ORIENTATION = 19
	// Roll: 270, Pitch: 0, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_270 MAV_SENSOR_ORIENTATION = 20
	// Roll: 270, Pitch: 0, Yaw: 45
	MAV_SENSOR_ROTATION_ROLL_270_YAW_45 MAV_SENSOR_ORIENTATION = 21
	// Roll: 270, Pitch: 0, Yaw: 90
	MAV_SENSOR_ROTATION_ROLL_270_YAW_90 MAV_SENSOR_ORIENTATION = 22
	// Roll: 270, Pitch: 0, Yaw: 135
	MAV_SENSOR_ROTATION_ROLL_270_YAW_135 MAV_SENSOR_ORIENTATION = 23
	// Roll: 0, Pitch: 90, Yaw: 0
	MAV_SENSOR_ROTATION_PITCH_90 MAV_SENSOR_ORIENTATION = 24
	// Roll: 0, Pitch: 270, Yaw: 0
	MAV_SENSOR_ROTATION_PITCH_270 MAV_SENSOR_ORIENTATION = 25
	// Roll: 0, Pitch: 180, Yaw: 90
	MAV_SENSOR_ROTATION_PITCH_180_YAW_90 MAV_SENSOR_ORIENTATION = 26
	// Roll: 0, Pitch: 180, Yaw: 270
	MAV_SENSOR_ROTATION_PITCH_180_YAW_270 MAV_SENSOR_ORIENTATION = 27
	// Roll: 90, Pitch: 90, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_90 MAV_SENSOR_ORIENTATION = 28
	// Roll: 180, Pitch: 90, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_180_PITCH_90 MAV_SENSOR_ORIENTATION = 29
	// Roll: 270, Pitch: 90, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_270_PITCH_90 MAV_SENSOR_ORIENTATION = 30
	// Roll: 90, Pitch: 180, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_180 MAV_SENSOR_ORIENTATION = 31
	// Roll: 270, Pitch: 180, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_270_PITCH_180 MAV_SENSOR_ORIENTATION = 32
	// Roll: 90, Pitch: 270, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_270 MAV_SENSOR_ORIENTATION = 33
	// Roll: 180, Pitch: 270, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_180_PITCH_270 MAV_SENSOR_ORIENTATION = 34
	// Roll: 270, Pitch: 270, Yaw: 0
	MAV_SENSOR_ROTATION_ROLL_270_PITCH_270 MAV_SENSOR_ORIENTATION = 35
	// Roll: 90, Pitch: 180, Yaw: 90
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_180_YAW_90 MAV_SENSOR_ORIENTATION = 36
	// Roll: 90, Pitch: 0, Yaw: 270
	MAV_SENSOR_ROTATION_ROLL_90_YAW_270 MAV_SENSOR_ORIENTATION = 37
	// Roll: 90, Pitch: 68, Yaw: 293
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_68_YAW_293 MAV_SENSOR_ORIENTATION = 38
	// Pitch: 315
	MAV_SENSOR_ROTATION_PITCH_315 MAV_SENSOR_ORIENTATION = 39
	// Roll: 90, Pitch: 315
	MAV_SENSOR_ROTATION_ROLL_90_PITCH_315 MAV_SENSOR_ORIENTATION = 40
	// Custom orientation
	MAV_SENSOR_ROTATION_CUSTOM MAV_SENSOR_ORIENTATION = 100
)

var labels_MAV_SENSOR_ORIENTATION = map[MAV_SENSOR_ORIENTATION]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_SENSOR_ORIENTATION) MarshalText() ([]byte, error) {
	if l, ok := labels_MAV_SENSOR_ORIENTATION[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_MAV_SENSOR_ORIENTATION = map[string]MAV_SENSOR_ORIENTATION{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_SENSOR_ORIENTATION) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_MAV_SENSOR_ORIENTATION[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e MAV_SENSOR_ORIENTATION) String() string {
	if l, ok := labels_MAV_SENSOR_ORIENTATION[e]; ok {
		return l
	}
	return "invalid value"
}
