//autogenerated:yes
//nolint:revive,misspell,govet,lll
package all

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/development"
)

// Information about GCS in control of this MAV. This should be broadcast at low rate (nominally 1 Hz) and emitted when ownership or takeover status change. Control over MAV is requested using MAV_CMD_REQUEST_OPERATOR_CONTROL.
type MessageControlStatus = development.MessageControlStatus
