//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ardupilotmega

// ICAROUS heartbeat
type MessageIcarousHeartbeat struct {
	// See the FMS_STATE enum.
	Status ICAROUS_FMS_STATE `mavenum:"uint8"`
}

// GetID implements the msg.Message interface.
func (*MessageIcarousHeartbeat) GetID() uint32 {
	return 42000
}
