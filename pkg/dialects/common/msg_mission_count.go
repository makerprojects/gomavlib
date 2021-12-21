//autogenerated:yes
//nolint:revive,misspell,govet,lll
package common

// This message is emitted as response to MISSION_REQUEST_LIST by the MAV and to initiate a write transaction. The GCS can then request the individual mission item based on the knowledge of the total number of waypoints.
type MessageMissionCount struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Number of mission items in the sequence
	Count uint16
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageMissionCount) GetID() uint32 {
	return 44
}
