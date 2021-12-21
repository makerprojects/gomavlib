//autogenerated:yes
//nolint:revive,misspell,govet,lll
package uavionix

// Superseded by ACTUATOR_OUTPUT_STATUS. The RAW values of the servo outputs (for RC input from the remote, use the RC_CHANNELS messages). The standard PPM modulation is as follows: 1000 microseconds: 0%, 2000 microseconds: 100%.
type MessageServoOutputRaw struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint32
	// Servo output port (set of 8 outputs = 1 port). Flight stacks running on Pixhawk should use: 0 = MAIN, 1 = AUX.
	Port uint8
	// Servo output 1 value
	Servo1Raw uint16
	// Servo output 2 value
	Servo2Raw uint16
	// Servo output 3 value
	Servo3Raw uint16
	// Servo output 4 value
	Servo4Raw uint16
	// Servo output 5 value
	Servo5Raw uint16
	// Servo output 6 value
	Servo6Raw uint16
	// Servo output 7 value
	Servo7Raw uint16
	// Servo output 8 value
	Servo8Raw uint16
	// Servo output 9 value
	Servo9Raw uint16 `mavext:"true"`
	// Servo output 10 value
	Servo10Raw uint16 `mavext:"true"`
	// Servo output 11 value
	Servo11Raw uint16 `mavext:"true"`
	// Servo output 12 value
	Servo12Raw uint16 `mavext:"true"`
	// Servo output 13 value
	Servo13Raw uint16 `mavext:"true"`
	// Servo output 14 value
	Servo14Raw uint16 `mavext:"true"`
	// Servo output 15 value
	Servo15Raw uint16 `mavext:"true"`
	// Servo output 16 value
	Servo16Raw uint16 `mavext:"true"`
}

// GetID implements the msg.Message interface.
func (*MessageServoOutputRaw) GetID() uint32 {
	return 36
}
