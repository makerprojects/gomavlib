//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
)

// Type of GPS fix
type GPS_FIX_TYPE = common.GPS_FIX_TYPE

const (
	// No GPS connected
	GPS_FIX_TYPE_NO_GPS GPS_FIX_TYPE = common.GPS_FIX_TYPE_NO_GPS
	// No position information, GPS is connected
	GPS_FIX_TYPE_NO_FIX GPS_FIX_TYPE = common.GPS_FIX_TYPE_NO_FIX
	// 2D position
	GPS_FIX_TYPE_2D_FIX GPS_FIX_TYPE = common.GPS_FIX_TYPE_2D_FIX
	// 3D position
	GPS_FIX_TYPE_3D_FIX GPS_FIX_TYPE = common.GPS_FIX_TYPE_3D_FIX
	// DGPS/SBAS aided 3D position
	GPS_FIX_TYPE_DGPS GPS_FIX_TYPE = common.GPS_FIX_TYPE_DGPS
	// RTK float, 3D position
	GPS_FIX_TYPE_RTK_FLOAT GPS_FIX_TYPE = common.GPS_FIX_TYPE_RTK_FLOAT
	// RTK Fixed, 3D position
	GPS_FIX_TYPE_RTK_FIXED GPS_FIX_TYPE = common.GPS_FIX_TYPE_RTK_FIXED
	// Static fixed, typically used for base stations
	GPS_FIX_TYPE_STATIC GPS_FIX_TYPE = common.GPS_FIX_TYPE_STATIC
	// PPP, 3D position.
	GPS_FIX_TYPE_PPP GPS_FIX_TYPE = common.GPS_FIX_TYPE_PPP
)
