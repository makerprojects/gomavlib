//autogenerated:yes
//nolint:revive,misspell,govet,lll
package pythonarraytest

// Handshake message to initiate, control and stop image streaming when using the Image Transmission Protocol: https://mavlink.io/en/services/image_transmission.html.
type MessageDataTransmissionHandshake struct {
	// Type of requested/acknowledged data.
	Type MAVLINK_DATA_STREAM_TYPE `mavenum:"uint8"`
	// total data size (set on ACK only).
	Size uint32
	// Width of a matrix or image.
	Width uint16
	// Height of a matrix or image.
	Height uint16
	// Number of packets being sent (set on ACK only).
	Packets uint16
	// Payload size per packet (normally 253 byte, see DATA field size in message ENCAPSULATED_DATA) (set on ACK only).
	Payload uint8
	// JPEG quality. Values: [1-100].
	JpgQuality uint8
}

// GetID implements the msg.Message interface.
func (*MessageDataTransmissionHandshake) GetID() uint32 {
	return 130
}
