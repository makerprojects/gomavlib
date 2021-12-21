//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

// Monitoring of sensorpod status
type MessageSensorpodStatus struct {
	// Timestamp in linuxtime (since 1.1.1970)
	Timestamp uint64
	// Rate of ROS topic 1
	VisensorRate_1 uint8
	// Rate of ROS topic 2
	VisensorRate_2 uint8
	// Rate of ROS topic 3
	VisensorRate_3 uint8
	// Rate of ROS topic 4
	VisensorRate_4 uint8
	// Number of recording nodes
	RecordingNodesCount uint8
	// Temperature of sensorpod CPU in
	CpuTemp uint8
	// Free space available in recordings directory in [Gb] * 1e2
	FreeSpace uint16
}

// GetID implements the msg.Message interface.
func (*MessageSensorpodStatus) GetID() uint32 {
	return 8012
}
