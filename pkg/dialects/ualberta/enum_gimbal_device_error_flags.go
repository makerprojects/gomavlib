//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ualberta

import (
	"errors"
)

// Gimbal device (low level) error flags (bitmap, 0 means no error)
type GIMBAL_DEVICE_ERROR_FLAGS int

const (
	// Gimbal device is limited by hardware roll limit.
	GIMBAL_DEVICE_ERROR_FLAGS_AT_ROLL_LIMIT GIMBAL_DEVICE_ERROR_FLAGS = 1
	// Gimbal device is limited by hardware pitch limit.
	GIMBAL_DEVICE_ERROR_FLAGS_AT_PITCH_LIMIT GIMBAL_DEVICE_ERROR_FLAGS = 2
	// Gimbal device is limited by hardware yaw limit.
	GIMBAL_DEVICE_ERROR_FLAGS_AT_YAW_LIMIT GIMBAL_DEVICE_ERROR_FLAGS = 4
	// There is an error with the gimbal encoders.
	GIMBAL_DEVICE_ERROR_FLAGS_ENCODER_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 8
	// There is an error with the gimbal power source.
	GIMBAL_DEVICE_ERROR_FLAGS_POWER_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 16
	// There is an error with the gimbal motor's.
	GIMBAL_DEVICE_ERROR_FLAGS_MOTOR_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 32
	// There is an error with the gimbal's software.
	GIMBAL_DEVICE_ERROR_FLAGS_SOFTWARE_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 64
	// There is an error with the gimbal's communication.
	GIMBAL_DEVICE_ERROR_FLAGS_COMMS_ERROR GIMBAL_DEVICE_ERROR_FLAGS = 128
	// Gimbal is currently calibrating.
	GIMBAL_DEVICE_ERROR_FLAGS_CALIBRATION_RUNNING GIMBAL_DEVICE_ERROR_FLAGS = 256
)

var labels_GIMBAL_DEVICE_ERROR_FLAGS = map[GIMBAL_DEVICE_ERROR_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GIMBAL_DEVICE_ERROR_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_GIMBAL_DEVICE_ERROR_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GIMBAL_DEVICE_ERROR_FLAGS = map[string]GIMBAL_DEVICE_ERROR_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GIMBAL_DEVICE_ERROR_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GIMBAL_DEVICE_ERROR_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GIMBAL_DEVICE_ERROR_FLAGS) String() string {
	if l, ok := labels_GIMBAL_DEVICE_ERROR_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
