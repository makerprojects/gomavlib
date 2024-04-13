//autogenerated:yes
//nolint:revive,misspell,govet,lll
package uavionix

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
)

// Time synchronization message.
// The message is used for both timesync requests and responses.
// The request is sent with `ts1=syncing component timestamp` and `tc1=0`, and may be broadcast or targeted to a specific system/component.
// The response is sent with `ts1=syncing component timestamp` (mirror back unchanged), and `tc1=responding component timestamp`, with the `target_system` and `target_component` set to ids of the original request.
// Systems can determine if they are receiving a request or response based on the value of `tc`.
// If the response has `target_system==target_component==0` the remote system has not been updated to use the component IDs and cannot reliably timesync; the requestor may report an error.
// Timestamps are UNIX Epoch time or time since system boot in nanoseconds (the timestamp format can be inferred by checking for the magnitude of the number; generally it doesn't matter as only the offset is used).
// The message sequence is repeated numerous times with results being filtered/averaged to estimate the offset.
type MessageTimesync = common.MessageTimesync
