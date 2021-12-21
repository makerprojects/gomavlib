//autogenerated:yes
//nolint:revive,misspell,govet,lll
package pythonarraytest

// Smart Battery information (static/infrequent update). Use for updates from: smart battery to flight stack, flight stack to GCS. Use BATTERY_STATUS for smart battery frequent updates.
type MessageSmartBatteryInfo struct {
	// Battery ID
	Id uint8
	// Function of the battery
	BatteryFunction MAV_BATTERY_FUNCTION `mavenum:"uint8"`
	// Type (chemistry) of the battery
	Type MAV_BATTERY_TYPE `mavenum:"uint8"`
	// Capacity when full according to manufacturer, -1: field not provided.
	CapacityFullSpecification int32
	// Capacity when full (accounting for battery degradation), -1: field not provided.
	CapacityFull int32
	// Charge/discharge cycle count. UINT16_MAX: field not provided.
	CycleCount uint16
	// Serial number in ASCII characters, 0 terminated. All 0: field not provided.
	SerialNumber string `mavlen:"16"`
	// Static device name in ASCII characters, 0 terminated. All 0: field not provided. Encode as manufacturer name then product name separated using an underscore.
	DeviceName string `mavlen:"50"`
	// Battery weight. 0: field not provided.
	Weight uint16
	// Minimum per-cell voltage when discharging. If not supplied set to UINT16_MAX value.
	DischargeMinimumVoltage uint16
	// Minimum per-cell voltage when charging. If not supplied set to UINT16_MAX value.
	ChargingMinimumVoltage uint16
	// Minimum per-cell voltage when resting. If not supplied set to UINT16_MAX value.
	RestingMinimumVoltage uint16
	// Maximum per-cell voltage when charged. 0: field not provided.
	ChargingMaximumVoltage uint16 `mavext:"true"`
	// Number of battery cells in series. 0: field not provided.
	CellsInSeries uint8 `mavext:"true"`
	// Maximum pack discharge current. 0: field not provided.
	DischargeMaximumCurrent uint32 `mavext:"true"`
	// Maximum pack discharge burst current. 0: field not provided.
	DischargeMaximumBurstCurrent uint32 `mavext:"true"`
	// Manufacture date (DD/MM/YYYY) in ASCII characters, 0 terminated. All 0: field not provided.
	ManufactureDate string `mavext:"true" mavlen:"11"`
}

// GetID implements the msg.Message interface.
func (*MessageSmartBatteryInfo) GetID() uint32 {
	return 370
}
