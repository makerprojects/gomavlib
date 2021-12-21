//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// Regular broadcast for the current latest event sequence number for a component. This is used to check for dropped events.
type MessageCurrentEventSequence struct {
	// Sequence number.
	Sequence uint16
	// Flag bitset.
	Flags MAV_EVENT_CURRENT_SEQUENCE_FLAGS `mavenum:"uint8"`
}

// GetID implements the msg.Message interface.
func (*MessageCurrentEventSequence) GetID() uint32 {
	return 411
}
