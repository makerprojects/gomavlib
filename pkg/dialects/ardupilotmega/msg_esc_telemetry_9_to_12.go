//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ardupilotmega

// ESC Telemetry Data for ESCs 9 to 12, matching data sent by BLHeli ESCs.
type MessageEscTelemetry_9To_12 struct {
	// Temperature.
	Temperature [4]uint8
	// Voltage.
	Voltage [4]uint16
	// Current.
	Current [4]uint16
	// Total current.
	Totalcurrent [4]uint16
	// RPM (eRPM).
	Rpm [4]uint16
	// count of telemetry packets received (wraps at 65535).
	Count [4]uint16
}

// GetID implements the msg.Message interface.
func (*MessageEscTelemetry_9To_12) GetID() uint32 {
	return 11032
}
