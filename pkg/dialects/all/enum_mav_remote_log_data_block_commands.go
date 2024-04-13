//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package all

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/ardupilotmega"
)

// Special ACK block numbers control activation of dataflash log streaming.
type MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS = ardupilotmega.MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS

const (
	// UAV to stop sending DataFlash blocks.
	MAV_REMOTE_LOG_DATA_BLOCK_STOP MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS = ardupilotmega.MAV_REMOTE_LOG_DATA_BLOCK_STOP
	// UAV to start sending DataFlash blocks.
	MAV_REMOTE_LOG_DATA_BLOCK_START MAV_REMOTE_LOG_DATA_BLOCK_COMMANDS = ardupilotmega.MAV_REMOTE_LOG_DATA_BLOCK_START
)
