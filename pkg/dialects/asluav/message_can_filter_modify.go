//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
)

// Modify the filter of what CAN messages to forward over the mavlink. This can be used to make CAN forwarding work well on low bandwidth links. The filtering is applied on bits 8 to 24 of the CAN id (2nd and 3rd bytes) which corresponds to the DroneCAN message ID for DroneCAN. Filters with more than 16 IDs can be constructed by sending multiple CAN_FILTER_MODIFY messages.
type MessageCanFilterModify = common.MessageCanFilterModify
