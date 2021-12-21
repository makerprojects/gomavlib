//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

// Information about a component. For camera components instead use CAMERA_INFORMATION, and for autopilots additionally use AUTOPILOT_VERSION. Components including GCSes should consider supporting requests of this message via MAV_CMD_REQUEST_MESSAGE.
type MessageComponentInformation struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// CRC32 of the TYPE_GENERAL file (can be used by a GCS for file caching).
	GeneralMetadataFileCrc uint32
	// Component definition URI for TYPE_GENERAL. This must be a MAVLink FTP URI and the file might be compressed with xz.
	GeneralMetadataUri string `mavlen:"100"`
	// CRC32 of the TYPE_PERIPHERALS file (can be used by a GCS for file caching).
	PeripheralsMetadataFileCrc uint32
	// (Optional) Component definition URI for TYPE_PERIPHERALS. This must be a MAVLink FTP URI and the file might be compressed with xz.
	PeripheralsMetadataUri string `mavlen:"100"`
}

// GetID implements the msg.Message interface.
func (*MessageComponentInformation) GetID() uint32 {
	return 395
}
