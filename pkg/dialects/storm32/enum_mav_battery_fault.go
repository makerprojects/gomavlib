//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package storm32

import (
	"github.com/bluenviron/gomavlib/v3/pkg/dialects/common"
)

// Smart battery supply status/fault flags (bitmask) for health indication. The battery must also report either MAV_BATTERY_CHARGE_STATE_FAILED or MAV_BATTERY_CHARGE_STATE_UNHEALTHY if any of these are set.
type MAV_BATTERY_FAULT = common.MAV_BATTERY_FAULT

const (
	// Battery has deep discharged.
	MAV_BATTERY_FAULT_DEEP_DISCHARGE MAV_BATTERY_FAULT = common.MAV_BATTERY_FAULT_DEEP_DISCHARGE
	// Voltage spikes.
	MAV_BATTERY_FAULT_SPIKES MAV_BATTERY_FAULT = common.MAV_BATTERY_FAULT_SPIKES
	// One or more cells have failed. Battery should also report MAV_BATTERY_CHARGE_STATE_FAILE (and should not be used).
	MAV_BATTERY_FAULT_CELL_FAIL MAV_BATTERY_FAULT = common.MAV_BATTERY_FAULT_CELL_FAIL
	// Over-current fault.
	MAV_BATTERY_FAULT_OVER_CURRENT MAV_BATTERY_FAULT = common.MAV_BATTERY_FAULT_OVER_CURRENT
	// Over-temperature fault.
	MAV_BATTERY_FAULT_OVER_TEMPERATURE MAV_BATTERY_FAULT = common.MAV_BATTERY_FAULT_OVER_TEMPERATURE
	// Under-temperature fault.
	MAV_BATTERY_FAULT_UNDER_TEMPERATURE MAV_BATTERY_FAULT = common.MAV_BATTERY_FAULT_UNDER_TEMPERATURE
	// Vehicle voltage is not compatible with this battery (batteries on same power rail should have similar voltage).
	MAV_BATTERY_FAULT_INCOMPATIBLE_VOLTAGE MAV_BATTERY_FAULT = common.MAV_BATTERY_FAULT_INCOMPATIBLE_VOLTAGE
	// Battery firmware is not compatible with current autopilot firmware.
	MAV_BATTERY_FAULT_INCOMPATIBLE_FIRMWARE MAV_BATTERY_FAULT = common.MAV_BATTERY_FAULT_INCOMPATIBLE_FIRMWARE
	// Battery is not compatible due to cell configuration (e.g. 5s1p when vehicle requires 6s).
	BATTERY_FAULT_INCOMPATIBLE_CELLS_CONFIGURATION MAV_BATTERY_FAULT = common.BATTERY_FAULT_INCOMPATIBLE_CELLS_CONFIGURATION
)
