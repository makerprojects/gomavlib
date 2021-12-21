//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

// Report button state change.
type MessageButtonChange struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Time of last change of button state.
	LastChangeMs uint32
	// Bitmap for state of buttons.
	State uint8
}

// GetID implements the msg.Message interface.
func (*MessageButtonChange) GetID() uint32 {
	return 257
}
