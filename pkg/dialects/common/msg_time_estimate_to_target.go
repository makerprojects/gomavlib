//autogenerated:yes
//nolint:revive,misspell,govet,lll
package common

// Time/duration estimates for various events and actions given the current vehicle state and position.
type MessageTimeEstimateToTarget struct {
	// Estimated time to complete the vehicle's configured "safe return" action from its current position (e.g. RTL, Smart RTL, etc.). -1 indicates that the vehicle is landed, or that no time estimate available.
	SafeReturn int32
	// Estimated time for vehicle to complete the LAND action from its current position. -1 indicates that the vehicle is landed, or that no time estimate available.
	Land int32
	// Estimated time for reaching/completing the currently active mission item. -1 means no time estimate available.
	MissionNextItem int32
	// Estimated time for completing the current mission. -1 means no mission active and/or no estimate available.
	MissionEnd int32
	// Estimated time for completing the current commanded action (i.e. Go To, Takeoff, Land, etc.). -1 means no action active and/or no estimate available.
	CommandedAction int32
}

// GetID implements the msg.Message interface.
func (*MessageTimeEstimateToTarget) GetID() uint32 {
	return 380
}
