//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// Backwards compatible version of SERIAL_UDB_EXTRA F5: format
type MessageSerialUdbExtraF5 struct {
	// Serial UDB YAWKP_AILERON Gain for Proporional control of navigation
	SueYawkpAileron float32 `mavname:"sue_YAWKP_AILERON"`
	// Serial UDB YAWKD_AILERON Gain for Rate control of navigation
	SueYawkdAileron float32 `mavname:"sue_YAWKD_AILERON"`
	// Serial UDB Extra ROLLKP Gain for Proportional control of roll stabilization
	SueRollkp float32 `mavname:"sue_ROLLKP"`
	// Serial UDB Extra ROLLKD Gain for Rate control of roll stabilization
	SueRollkd float32 `mavname:"sue_ROLLKD"`
}

// GetID implements the msg.Message interface.
func (*MessageSerialUdbExtraF5) GetID() uint32 {
	return 173
}
