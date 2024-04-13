//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package avssuas

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
)

// Camera sources for MAV_CMD_SET_CAMERA_SOURCE
type CAMERA_SOURCE = common.CAMERA_SOURCE

const (
	// Default camera source.
	CAMERA_SOURCE_DEFAULT CAMERA_SOURCE = common.CAMERA_SOURCE_DEFAULT
	// RGB camera source.
	CAMERA_SOURCE_RGB CAMERA_SOURCE = common.CAMERA_SOURCE_RGB
	// IR camera source.
	CAMERA_SOURCE_IR CAMERA_SOURCE = common.CAMERA_SOURCE_IR
	// NDVI camera source.
	CAMERA_SOURCE_NDVI CAMERA_SOURCE = common.CAMERA_SOURCE_NDVI
)
