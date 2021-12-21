//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

// Telemetry of power generation system. Alternator or mechanical generator.
type MessageGeneratorStatus struct {
	// Status flags.
	Status MAV_GENERATOR_STATUS_FLAG `mavenum:"uint64"`
	// Speed of electrical generator or alternator. UINT16_MAX: field not provided.
	GeneratorSpeed uint16
	// Current into/out of battery. Positive for out. Negative for in. NaN: field not provided.
	BatteryCurrent float32
	// Current going to the UAV. If battery current not available this is the DC current from the generator. Positive for out. Negative for in. NaN: field not provided
	LoadCurrent float32
	// The power being generated. NaN: field not provided
	PowerGenerated float32
	// Voltage of the bus seen at the generator, or battery bus if battery bus is controlled by generator and at a different voltage to main bus.
	BusVoltage float32
	// The temperature of the rectifier or power converter. INT16_MAX: field not provided.
	RectifierTemperature int16
	// The target battery current. Positive for out. Negative for in. NaN: field not provided
	BatCurrentSetpoint float32
	// The temperature of the mechanical motor, fuel cell core or generator. INT16_MAX: field not provided.
	GeneratorTemperature int16
	// Seconds this generator has run since it was rebooted. UINT32_MAX: field not provided.
	Runtime uint32
	// Seconds until this generator requires maintenance.  A negative value indicates maintenance is past-due. INT32_MAX: field not provided.
	TimeUntilMaintenance int32
}

// GetID implements the msg.Message interface.
func (*MessageGeneratorStatus) GetID() uint32 {
	return 373
}
