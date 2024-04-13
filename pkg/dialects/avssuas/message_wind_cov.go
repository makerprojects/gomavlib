//autogenerated:yes
//nolint:revive,misspell,govet,lll
package avssuas

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
)

// Wind estimate from vehicle. Note that despite the name, this message does not actually contain any covariances but instead variability and accuracy fields in terms of standard deviation (1-STD).
type MessageWindCov = common.MessageWindCov
