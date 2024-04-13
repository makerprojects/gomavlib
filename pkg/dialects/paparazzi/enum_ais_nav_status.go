//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package paparazzi

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
)

// Navigational status of AIS vessel, enum duplicated from AIS standard, https://gpsd.gitlab.io/gpsd/AIVDM.html
type AIS_NAV_STATUS = common.AIS_NAV_STATUS

const (
	// Under way using engine.
	UNDER_WAY                           AIS_NAV_STATUS = common.UNDER_WAY
	AIS_NAV_ANCHORED                    AIS_NAV_STATUS = common.AIS_NAV_ANCHORED
	AIS_NAV_UN_COMMANDED                AIS_NAV_STATUS = common.AIS_NAV_UN_COMMANDED
	AIS_NAV_RESTRICTED_MANOEUVERABILITY AIS_NAV_STATUS = common.AIS_NAV_RESTRICTED_MANOEUVERABILITY
	AIS_NAV_DRAUGHT_CONSTRAINED         AIS_NAV_STATUS = common.AIS_NAV_DRAUGHT_CONSTRAINED
	AIS_NAV_MOORED                      AIS_NAV_STATUS = common.AIS_NAV_MOORED
	AIS_NAV_AGROUND                     AIS_NAV_STATUS = common.AIS_NAV_AGROUND
	AIS_NAV_FISHING                     AIS_NAV_STATUS = common.AIS_NAV_FISHING
	AIS_NAV_SAILING                     AIS_NAV_STATUS = common.AIS_NAV_SAILING
	AIS_NAV_RESERVED_HSC                AIS_NAV_STATUS = common.AIS_NAV_RESERVED_HSC
	AIS_NAV_RESERVED_WIG                AIS_NAV_STATUS = common.AIS_NAV_RESERVED_WIG
	AIS_NAV_RESERVED_1                  AIS_NAV_STATUS = common.AIS_NAV_RESERVED_1
	AIS_NAV_RESERVED_2                  AIS_NAV_STATUS = common.AIS_NAV_RESERVED_2
	AIS_NAV_RESERVED_3                  AIS_NAV_STATUS = common.AIS_NAV_RESERVED_3
	// Search And Rescue Transponder.
	AIS_NAV_AIS_SART AIS_NAV_STATUS = common.AIS_NAV_AIS_SART
	// Not available (default).
	AIS_NAV_UNKNOWN AIS_NAV_STATUS = common.AIS_NAV_UNKNOWN
)
