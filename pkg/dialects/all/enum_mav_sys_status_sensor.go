//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
)

// These encode the sensors whose status is sent as part of the SYS_STATUS message.
type MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR

const (
	// 0x01 3D gyro
	MAV_SYS_STATUS_SENSOR_3D_GYRO MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_3D_GYRO
	// 0x02 3D accelerometer
	MAV_SYS_STATUS_SENSOR_3D_ACCEL MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_3D_ACCEL
	// 0x04 3D magnetometer
	MAV_SYS_STATUS_SENSOR_3D_MAG MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_3D_MAG
	// 0x08 absolute pressure
	MAV_SYS_STATUS_SENSOR_ABSOLUTE_PRESSURE MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_ABSOLUTE_PRESSURE
	// 0x10 differential pressure
	MAV_SYS_STATUS_SENSOR_DIFFERENTIAL_PRESSURE MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_DIFFERENTIAL_PRESSURE
	// 0x20 GPS
	MAV_SYS_STATUS_SENSOR_GPS MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_GPS
	// 0x40 optical flow
	MAV_SYS_STATUS_SENSOR_OPTICAL_FLOW MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_OPTICAL_FLOW
	// 0x80 computer vision position
	MAV_SYS_STATUS_SENSOR_VISION_POSITION MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_VISION_POSITION
	// 0x100 laser based position
	MAV_SYS_STATUS_SENSOR_LASER_POSITION MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_LASER_POSITION
	// 0x200 external ground truth (Vicon or Leica)
	MAV_SYS_STATUS_SENSOR_EXTERNAL_GROUND_TRUTH MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_EXTERNAL_GROUND_TRUTH
	// 0x400 3D angular rate control
	MAV_SYS_STATUS_SENSOR_ANGULAR_RATE_CONTROL MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_ANGULAR_RATE_CONTROL
	// 0x800 attitude stabilization
	MAV_SYS_STATUS_SENSOR_ATTITUDE_STABILIZATION MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_ATTITUDE_STABILIZATION
	// 0x1000 yaw position
	MAV_SYS_STATUS_SENSOR_YAW_POSITION MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_YAW_POSITION
	// 0x2000 z/altitude control
	MAV_SYS_STATUS_SENSOR_Z_ALTITUDE_CONTROL MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_Z_ALTITUDE_CONTROL
	// 0x4000 x/y position control
	MAV_SYS_STATUS_SENSOR_XY_POSITION_CONTROL MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_XY_POSITION_CONTROL
	// 0x8000 motor outputs / control
	MAV_SYS_STATUS_SENSOR_MOTOR_OUTPUTS MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_MOTOR_OUTPUTS
	// 0x10000 RC receiver
	MAV_SYS_STATUS_SENSOR_RC_RECEIVER MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_RC_RECEIVER
	// 0x20000 2nd 3D gyro
	MAV_SYS_STATUS_SENSOR_3D_GYRO2 MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_3D_GYRO2
	// 0x40000 2nd 3D accelerometer
	MAV_SYS_STATUS_SENSOR_3D_ACCEL2 MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_3D_ACCEL2
	// 0x80000 2nd 3D magnetometer
	MAV_SYS_STATUS_SENSOR_3D_MAG2 MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_3D_MAG2
	// 0x100000 geofence
	MAV_SYS_STATUS_GEOFENCE MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_GEOFENCE
	// 0x200000 AHRS subsystem health
	MAV_SYS_STATUS_AHRS MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_AHRS
	// 0x400000 Terrain subsystem health
	MAV_SYS_STATUS_TERRAIN MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_TERRAIN
	// 0x800000 Motors are reversed
	MAV_SYS_STATUS_REVERSE_MOTOR MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_REVERSE_MOTOR
	// 0x1000000 Logging
	MAV_SYS_STATUS_LOGGING MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_LOGGING
	// 0x2000000 Battery
	MAV_SYS_STATUS_SENSOR_BATTERY MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_BATTERY
	// 0x4000000 Proximity
	MAV_SYS_STATUS_SENSOR_PROXIMITY MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_PROXIMITY
	// 0x8000000 Satellite Communication
	MAV_SYS_STATUS_SENSOR_SATCOM MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_SATCOM
	// 0x10000000 pre-arm check status. Always healthy when armed
	MAV_SYS_STATUS_PREARM_CHECK MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_PREARM_CHECK
	// 0x20000000 Avoidance/collision prevention
	MAV_SYS_STATUS_OBSTACLE_AVOIDANCE MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_OBSTACLE_AVOIDANCE
	// 0x40000000 propulsion (actuator, esc, motor or propellor)
	MAV_SYS_STATUS_SENSOR_PROPULSION MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_SENSOR_PROPULSION
	// 0x80000000 Extended bit-field are used for further sensor status bits (needs to be set in onboard_control_sensors_present only)
	MAV_SYS_STATUS_EXTENSION_USED MAV_SYS_STATUS_SENSOR = common.MAV_SYS_STATUS_EXTENSION_USED
)
