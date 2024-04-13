//autogenerated:yes
//nolint:revive,misspell,govet,lll
package avssuas

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
)

// Sent from simulation to autopilot. The RAW values of the RC channels received. The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%. Individual receivers/transmitters might violate this specification.
type MessageHilRcInputsRaw = common.MessageHilRcInputsRaw
