//autogenerated:yes
//nolint:revive,misspell,govet,lll
package common

// The interval between messages for a particular MAVLink message ID. This message is the response to the MAV_CMD_GET_MESSAGE_INTERVAL command. This interface replaces DATA_STREAM.
type MessageMessageInterval struct {
	// The ID of the requested MAVLink message. v1.0 is limited to 254 messages.
	MessageId uint16
	// The interval between two messages. A value of -1 indicates this stream is disabled, 0 indicates it is not available, &gt; 0 indicates the interval at which it is sent.
	IntervalUs int32
}

// GetID implements the msg.Message interface.
func (*MessageMessageInterval) GetID() uint32 {
	return 244
}
