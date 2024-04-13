//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package pythonarraytest

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
)

// Flags in the HIL_SENSOR message indicate which fields have updated since the last message
type HIL_SENSOR_UPDATED_FLAGS = common.HIL_SENSOR_UPDATED_FLAGS

const (
	// None of the fields in HIL_SENSOR have been updated
	HIL_SENSOR_UPDATED_NONE HIL_SENSOR_UPDATED_FLAGS = common.HIL_SENSOR_UPDATED_NONE
	// The value in the xacc field has been updated
	HIL_SENSOR_UPDATED_XACC HIL_SENSOR_UPDATED_FLAGS = common.HIL_SENSOR_UPDATED_XACC
	// The value in the yacc field has been updated
	HIL_SENSOR_UPDATED_YACC HIL_SENSOR_UPDATED_FLAGS = common.HIL_SENSOR_UPDATED_YACC
	// The value in the zacc field has been updated
	HIL_SENSOR_UPDATED_ZACC HIL_SENSOR_UPDATED_FLAGS = common.HIL_SENSOR_UPDATED_ZACC
	// The value in the xgyro field has been updated
	HIL_SENSOR_UPDATED_XGYRO HIL_SENSOR_UPDATED_FLAGS = common.HIL_SENSOR_UPDATED_XGYRO
	// The value in the ygyro field has been updated
	HIL_SENSOR_UPDATED_YGYRO HIL_SENSOR_UPDATED_FLAGS = common.HIL_SENSOR_UPDATED_YGYRO
	// The value in the zgyro field has been updated
	HIL_SENSOR_UPDATED_ZGYRO HIL_SENSOR_UPDATED_FLAGS = common.HIL_SENSOR_UPDATED_ZGYRO
	// The value in the xmag field has been updated
	HIL_SENSOR_UPDATED_XMAG HIL_SENSOR_UPDATED_FLAGS = common.HIL_SENSOR_UPDATED_XMAG
	// The value in the ymag field has been updated
	HIL_SENSOR_UPDATED_YMAG HIL_SENSOR_UPDATED_FLAGS = common.HIL_SENSOR_UPDATED_YMAG
	// The value in the zmag field has been updated
	HIL_SENSOR_UPDATED_ZMAG HIL_SENSOR_UPDATED_FLAGS = common.HIL_SENSOR_UPDATED_ZMAG
	// The value in the abs_pressure field has been updated
	HIL_SENSOR_UPDATED_ABS_PRESSURE HIL_SENSOR_UPDATED_FLAGS = common.HIL_SENSOR_UPDATED_ABS_PRESSURE
	// The value in the diff_pressure field has been updated
	HIL_SENSOR_UPDATED_DIFF_PRESSURE HIL_SENSOR_UPDATED_FLAGS = common.HIL_SENSOR_UPDATED_DIFF_PRESSURE
	// The value in the pressure_alt field has been updated
	HIL_SENSOR_UPDATED_PRESSURE_ALT HIL_SENSOR_UPDATED_FLAGS = common.HIL_SENSOR_UPDATED_PRESSURE_ALT
	// The value in the temperature field has been updated
	HIL_SENSOR_UPDATED_TEMPERATURE HIL_SENSOR_UPDATED_FLAGS = common.HIL_SENSOR_UPDATED_TEMPERATURE
	// Full reset of attitude/position/velocities/etc was performed in sim (Bit 31).
	HIL_SENSOR_UPDATED_RESET HIL_SENSOR_UPDATED_FLAGS = common.HIL_SENSOR_UPDATED_RESET
)
