//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// Erase all logs
type MessageLogErase struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
}

// GetID implements the msg.Message interface.
func (*MessageLogErase) GetID() uint32 {
	return 121
}
