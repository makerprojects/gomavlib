//autogenerated:yes
//nolint:revive,misspell,govet,lll
package pythonarraytest

// Send a command with up to seven parameters to the MAV. The command microservice is documented at https://mavlink.io/en/services/command.html
type MessageCommandLong struct {
	// System which should execute the command
	TargetSystem uint8
	// Component which should execute the command, 0 for all components
	TargetComponent uint8
	// Command ID (of command to send).
	Command MAV_CMD `mavenum:"uint16"`
	// 0: First transmission of this command. 1-255: Confirmation transmissions (e.g. for kill command)
	Confirmation uint8
	// Parameter 1 (for the specific command).
	Param1 float32
	// Parameter 2 (for the specific command).
	Param2 float32
	// Parameter 3 (for the specific command).
	Param3 float32
	// Parameter 4 (for the specific command).
	Param4 float32
	// Parameter 5 (for the specific command).
	Param5 float32
	// Parameter 6 (for the specific command).
	Param6 float32
	// Parameter 7 (for the specific command).
	Param7 float32
}

// GetID implements the msg.Message interface.
func (*MessageCommandLong) GetID() uint32 {
	return 76
}
