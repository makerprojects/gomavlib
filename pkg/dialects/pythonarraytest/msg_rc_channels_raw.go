//autogenerated:yes
//nolint:revive,misspell,govet,lll
package pythonarraytest

// The RAW values of the RC channels received. The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%. A value of UINT16_MAX implies the channel is unused. Individual receivers/transmitters might violate this specification.
type MessageRcChannelsRaw struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Servo output port (set of 8 outputs = 1 port). Flight stacks running on Pixhawk should use: 0 = MAIN, 1 = AUX.
	Port uint8
	// RC channel 1 value.
	Chan1Raw uint16
	// RC channel 2 value.
	Chan2Raw uint16
	// RC channel 3 value.
	Chan3Raw uint16
	// RC channel 4 value.
	Chan4Raw uint16
	// RC channel 5 value.
	Chan5Raw uint16
	// RC channel 6 value.
	Chan6Raw uint16
	// RC channel 7 value.
	Chan7Raw uint16
	// RC channel 8 value.
	Chan8Raw uint16
	// Receive signal strength indicator in device-dependent units/scale. Values: [0-254], UINT8_MAX: invalid/unknown.
	Rssi uint8
}

// GetID implements the msg.Message interface.
func (*MessageRcChannelsRaw) GetID() uint32 {
	return 35
}
