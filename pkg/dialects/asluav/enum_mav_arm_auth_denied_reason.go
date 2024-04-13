//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package asluav

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
)

type MAV_ARM_AUTH_DENIED_REASON = common.MAV_ARM_AUTH_DENIED_REASON

const (
	// Not a specific reason
	MAV_ARM_AUTH_DENIED_REASON_GENERIC MAV_ARM_AUTH_DENIED_REASON = common.MAV_ARM_AUTH_DENIED_REASON_GENERIC
	// Authorizer will send the error as string to GCS
	MAV_ARM_AUTH_DENIED_REASON_NONE MAV_ARM_AUTH_DENIED_REASON = common.MAV_ARM_AUTH_DENIED_REASON_NONE
	// At least one waypoint have a invalid value
	MAV_ARM_AUTH_DENIED_REASON_INVALID_WAYPOINT MAV_ARM_AUTH_DENIED_REASON = common.MAV_ARM_AUTH_DENIED_REASON_INVALID_WAYPOINT
	// Timeout in the authorizer process(in case it depends on network)
	MAV_ARM_AUTH_DENIED_REASON_TIMEOUT MAV_ARM_AUTH_DENIED_REASON = common.MAV_ARM_AUTH_DENIED_REASON_TIMEOUT
	// Airspace of the mission in use by another vehicle, second result parameter can have the waypoint id that caused it to be denied.
	MAV_ARM_AUTH_DENIED_REASON_AIRSPACE_IN_USE MAV_ARM_AUTH_DENIED_REASON = common.MAV_ARM_AUTH_DENIED_REASON_AIRSPACE_IN_USE
	// Weather is not good to fly
	MAV_ARM_AUTH_DENIED_REASON_BAD_WEATHER MAV_ARM_AUTH_DENIED_REASON = common.MAV_ARM_AUTH_DENIED_REASON_BAD_WEATHER
)
