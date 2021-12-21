//autogenerated:yes
//nolint:revive,misspell,govet,lll
package paparazzi

// The scaled values of the RC channels received: (-100%) -10000, (0%) 0, (100%) 10000. Channels that are inactive should be set to UINT16_MAX.
type MessageRcChannelsScaled struct {
	// Timestamp (time since system boot).
	TimeBootMs uint32
	// Servo output port (set of 8 outputs = 1 port). Flight stacks running on Pixhawk should use: 0 = MAIN, 1 = AUX.
	Port uint8
	// RC channel 1 value scaled.
	Chan1Scaled int16
	// RC channel 2 value scaled.
	Chan2Scaled int16
	// RC channel 3 value scaled.
	Chan3Scaled int16
	// RC channel 4 value scaled.
	Chan4Scaled int16
	// RC channel 5 value scaled.
	Chan5Scaled int16
	// RC channel 6 value scaled.
	Chan6Scaled int16
	// RC channel 7 value scaled.
	Chan7Scaled int16
	// RC channel 8 value scaled.
	Chan8Scaled int16
	// Receive signal strength indicator in device-dependent units/scale. Values: [0-254], UINT8_MAX: invalid/unknown.
	Rssi uint8
}

// GetID implements the msg.Message interface.
func (*MessageRcChannelsScaled) GetID() uint32 {
	return 34
}
