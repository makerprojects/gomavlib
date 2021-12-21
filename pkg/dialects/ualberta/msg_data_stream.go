//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

// Data stream status information.
type MessageDataStream struct {
	// The ID of the requested data stream
	StreamId uint8
	// The message rate
	MessageRate uint16
	// 1 stream is enabled, 0 stream is stopped.
	OnOff uint8
}

// GetID implements the msg.Message interface.
func (*MessageDataStream) GetID() uint32 {
	return 67
}
