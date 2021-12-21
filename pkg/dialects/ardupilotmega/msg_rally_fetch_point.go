//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ardupilotmega

// Request a current rally point from MAV. MAV should respond with a RALLY_POINT message. MAV should not respond if the request is invalid.
type MessageRallyFetchPoint struct {
	// System ID.
	TargetSystem uint8
	// Component ID.
	TargetComponent uint8
	// Point index (first point is 0).
	Idx uint8
}

// GetID implements the msg.Message interface.
func (*MessageRallyFetchPoint) GetID() uint32 {
	return 176
}
