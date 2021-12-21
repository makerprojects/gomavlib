//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// Backwards compatible version of SERIAL_UDB_EXTRA F18 format
type MessageSerialUdbExtraF18 struct {
	// SUE Angle of Attack Normal
	AngleOfAttackNormal float32
	// SUE Angle of Attack Inverted
	AngleOfAttackInverted float32
	// SUE Elevator Trim Normal
	ElevatorTrimNormal float32
	// SUE Elevator Trim Inverted
	ElevatorTrimInverted float32
	// SUE reference_speed
	ReferenceSpeed float32
}

// GetID implements the msg.Message interface.
func (*MessageSerialUdbExtraF18) GetID() uint32 {
	return 184
}
