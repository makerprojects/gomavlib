//autogenerated:yes
//nolint:revive,misspell,govet,lll
package asluav

// Extended state information for ASLUAVs
type MessageAsluavStatus struct {
	// Status of the position-indicator LEDs
	LedStatus uint8 `mavname:"LED_status"`
	// Status of the IRIDIUM satellite communication system
	SatcomStatus uint8 `mavname:"SATCOM_status"`
	// Status vector for up to 8 servos
	ServoStatus [8]uint8 `mavname:"Servo_status"`
	// Motor RPM
	MotorRpm float32 `mavname:"Motor_rpm"`
}

// GetID implements the msg.Message interface.
func (*MessageAsluavStatus) GetID() uint32 {
	return 8006
}
