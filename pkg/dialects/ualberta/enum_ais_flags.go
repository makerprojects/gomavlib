//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ualberta

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
)

// These flags are used in the AIS_VESSEL.fields bitmask to indicate validity of data in the other message fields. When set, the data is valid.
type AIS_FLAGS = common.AIS_FLAGS

const (
	// 1 = Position accuracy less than 10m, 0 = position accuracy greater than 10m.
	AIS_FLAGS_POSITION_ACCURACY AIS_FLAGS = common.AIS_FLAGS_POSITION_ACCURACY
	AIS_FLAGS_VALID_COG         AIS_FLAGS = common.AIS_FLAGS_VALID_COG
	AIS_FLAGS_VALID_VELOCITY    AIS_FLAGS = common.AIS_FLAGS_VALID_VELOCITY
	// 1 = Velocity over 52.5765m/s (102.2 knots)
	AIS_FLAGS_HIGH_VELOCITY   AIS_FLAGS = common.AIS_FLAGS_HIGH_VELOCITY
	AIS_FLAGS_VALID_TURN_RATE AIS_FLAGS = common.AIS_FLAGS_VALID_TURN_RATE
	// Only the sign of the returned turn rate value is valid, either greater than 5deg/30s or less than -5deg/30s
	AIS_FLAGS_TURN_RATE_SIGN_ONLY AIS_FLAGS = common.AIS_FLAGS_TURN_RATE_SIGN_ONLY
	AIS_FLAGS_VALID_DIMENSIONS    AIS_FLAGS = common.AIS_FLAGS_VALID_DIMENSIONS
	// Distance to bow is larger than 511m
	AIS_FLAGS_LARGE_BOW_DIMENSION AIS_FLAGS = common.AIS_FLAGS_LARGE_BOW_DIMENSION
	// Distance to stern is larger than 511m
	AIS_FLAGS_LARGE_STERN_DIMENSION AIS_FLAGS = common.AIS_FLAGS_LARGE_STERN_DIMENSION
	// Distance to port side is larger than 63m
	AIS_FLAGS_LARGE_PORT_DIMENSION AIS_FLAGS = common.AIS_FLAGS_LARGE_PORT_DIMENSION
	// Distance to starboard side is larger than 63m
	AIS_FLAGS_LARGE_STARBOARD_DIMENSION AIS_FLAGS = common.AIS_FLAGS_LARGE_STARBOARD_DIMENSION
	AIS_FLAGS_VALID_CALLSIGN            AIS_FLAGS = common.AIS_FLAGS_VALID_CALLSIGN
	AIS_FLAGS_VALID_NAME                AIS_FLAGS = common.AIS_FLAGS_VALID_NAME
)
