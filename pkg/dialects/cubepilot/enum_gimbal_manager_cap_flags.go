//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package cubepilot

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
)

// Gimbal manager high level capability flags (bitmap). The first 16 bits are identical to the GIMBAL_DEVICE_CAP_FLAGS. However, the gimbal manager does not need to copy the flags from the gimbal but can also enhance the capabilities and thus add flags.
type GIMBAL_MANAGER_CAP_FLAGS = common.GIMBAL_MANAGER_CAP_FLAGS

const (
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_RETRACT.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_RETRACT GIMBAL_MANAGER_CAP_FLAGS = common.GIMBAL_MANAGER_CAP_FLAGS_HAS_RETRACT
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_NEUTRAL.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_NEUTRAL GIMBAL_MANAGER_CAP_FLAGS = common.GIMBAL_MANAGER_CAP_FLAGS_HAS_NEUTRAL
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_AXIS.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_AXIS GIMBAL_MANAGER_CAP_FLAGS = common.GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_AXIS
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_FOLLOW.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_FOLLOW GIMBAL_MANAGER_CAP_FLAGS = common.GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_FOLLOW
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_ROLL_LOCK.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_LOCK GIMBAL_MANAGER_CAP_FLAGS = common.GIMBAL_MANAGER_CAP_FLAGS_HAS_ROLL_LOCK
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_AXIS.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_AXIS GIMBAL_MANAGER_CAP_FLAGS = common.GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_AXIS
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_FOLLOW.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_FOLLOW GIMBAL_MANAGER_CAP_FLAGS = common.GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_FOLLOW
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_PITCH_LOCK.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_LOCK GIMBAL_MANAGER_CAP_FLAGS = common.GIMBAL_MANAGER_CAP_FLAGS_HAS_PITCH_LOCK
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_AXIS.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_AXIS GIMBAL_MANAGER_CAP_FLAGS = common.GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_AXIS
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_FOLLOW.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_FOLLOW GIMBAL_MANAGER_CAP_FLAGS = common.GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_FOLLOW
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_YAW_LOCK.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_LOCK GIMBAL_MANAGER_CAP_FLAGS = common.GIMBAL_MANAGER_CAP_FLAGS_HAS_YAW_LOCK
	// Based on GIMBAL_DEVICE_CAP_FLAGS_SUPPORTS_INFINITE_YAW.
	GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_INFINITE_YAW GIMBAL_MANAGER_CAP_FLAGS = common.GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_INFINITE_YAW
	// Based on GIMBAL_DEVICE_CAP_FLAGS_SUPPORTS_YAW_IN_EARTH_FRAME.
	GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_YAW_IN_EARTH_FRAME GIMBAL_MANAGER_CAP_FLAGS = common.GIMBAL_MANAGER_CAP_FLAGS_SUPPORTS_YAW_IN_EARTH_FRAME
	// Based on GIMBAL_DEVICE_CAP_FLAGS_HAS_RC_INPUTS.
	GIMBAL_MANAGER_CAP_FLAGS_HAS_RC_INPUTS GIMBAL_MANAGER_CAP_FLAGS = common.GIMBAL_MANAGER_CAP_FLAGS_HAS_RC_INPUTS
	// Gimbal manager supports to point to a local position.
	GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_LOCAL GIMBAL_MANAGER_CAP_FLAGS = common.GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_LOCAL
	// Gimbal manager supports to point to a global latitude, longitude, altitude position.
	GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_GLOBAL GIMBAL_MANAGER_CAP_FLAGS = common.GIMBAL_MANAGER_CAP_FLAGS_CAN_POINT_LOCATION_GLOBAL
)
