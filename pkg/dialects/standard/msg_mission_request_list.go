//autogenerated:yes
//nolint:revive,misspell,govet,lll
package standard

// Request the overall list of mission items from the system/component.
type MessageMissionRequestList struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// Mission type.
	MissionType MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageMissionRequestList) GetID() uint32 {
	return 43
}
