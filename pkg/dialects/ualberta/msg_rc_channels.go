//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

// The PPM values of the RC channels received. The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%.  A value of UINT16_MAX implies the channel is unused. Individual receivers/transmitters might violate this specification.
type MessageRcChannels struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Total number of RC channels being received. This can be larger than 18, indicating that more channels are available but not given in this message. This value should be 0 when no RC channels are available.
	Chancount uint8
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
	// RC channel 9 value.
	Chan9Raw uint16
	// RC channel 10 value.
	Chan10Raw uint16
	// RC channel 11 value.
	Chan11Raw uint16
	// RC channel 12 value.
	Chan12Raw uint16
	// RC channel 13 value.
	Chan13Raw uint16
	// RC channel 14 value.
	Chan14Raw uint16
	// RC channel 15 value.
	Chan15Raw uint16
	// RC channel 16 value.
	Chan16Raw uint16
	// RC channel 17 value.
	Chan17Raw uint16
	// RC channel 18 value.
	Chan18Raw uint16
	// Receive signal strength indicator in device-dependent units/scale. Values: [0-254], UINT8_MAX: invalid/unknown.
	Rssi uint8
}

// GetID implements the msg.Message interface.
func (*MessageRcChannels) GetID() uint32 {
	return 65
}
