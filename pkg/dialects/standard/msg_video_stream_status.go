//autogenerated:yes
//nolint:revive,misspell,govet,lll
package standard

// Information about the status of a video stream. It may be requested using MAV_CMD_REQUEST_MESSAGE.
type MessageVideoStreamStatus struct {
	// Video Stream ID (1 for first, 2 for second, etc.)
	StreamId uint8
	// Bitmap of stream status flags
	Flags VIDEO_STREAM_STATUS_FLAGS `mavenum:"uint16"`
	// Frame rate
	Framerate float32
	// Horizontal resolution
	ResolutionH uint16
	// Vertical resolution
	ResolutionV uint16
	// Bit rate
	Bitrate uint32
	// Video image rotation clockwise
	Rotation uint16
	// Horizontal Field of view
	Hfov uint16
}

// GetID implements the msg.Message interface.
func (*MessageVideoStreamStatus) GetID() uint32 {
	return 270
}
