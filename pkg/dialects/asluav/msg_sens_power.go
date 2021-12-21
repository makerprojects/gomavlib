//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

// Voltage and current sensor data
type MessageSensPower struct {
	// Power board voltage sensor reading
	Adc121VspbVolt float32
	// Power board current sensor reading
	Adc121CspbAmp float32
	// Board current sensor 1 reading
	Adc121Cs1Amp float32
	// Board current sensor 2 reading
	Adc121Cs2Amp float32
}

// GetID implements the msg.Message interface.
func (*MessageSensPower) GetID() uint32 {
	return 8002
}
