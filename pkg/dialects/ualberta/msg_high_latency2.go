//autogenerated:yes
//nolint:revive,misspell,govet,lll
package ualberta

// Message appropriate for high latency connections like Iridium (version 2)
type MessageHighLatency2 struct {
	// Timestamp (milliseconds since boot or Unix epoch)
	Timestamp uint32
	// Type of the MAV (quadrotor, helicopter, etc.)
	Type MAV_TYPE `mavenum:"uint8"`
	// Autopilot type / class. Use MAV_AUTOPILOT_INVALID for components that are not flight controllers.
	Autopilot MAV_AUTOPILOT `mavenum:"uint8"`
	// A bitfield for use for autopilot-specific flags (2 byte version).
	CustomMode uint16
	// Latitude
	Latitude int32
	// Longitude
	Longitude int32
	// Altitude above mean sea level
	Altitude int16
	// Altitude setpoint
	TargetAltitude int16
	// Heading
	Heading uint8
	// Heading setpoint
	TargetHeading uint8
	// Distance to target waypoint or position
	TargetDistance uint16
	// Throttle
	Throttle uint8
	// Airspeed
	Airspeed uint8
	// Airspeed setpoint
	AirspeedSp uint8
	// Groundspeed
	Groundspeed uint8
	// Windspeed
	Windspeed uint8
	// Wind heading
	WindHeading uint8
	// Maximum error horizontal position since last message
	Eph uint8
	// Maximum error vertical position since last message
	Epv uint8
	// Air temperature from airspeed sensor
	TemperatureAir int8
	// Maximum climb rate magnitude since last message
	ClimbRate int8
	// Battery level (-1 if field not provided).
	Battery int8
	// Current waypoint number
	WpNum uint16
	// Bitmap of failure flags.
	FailureFlags HL_FAILURE_FLAG `mavenum:"uint16"`
	// Field for custom payload.
	Custom0 int8
	// Field for custom payload.
	Custom1 int8
	// Field for custom payload.
	Custom2 int8
}

// GetID implements the msg.Message interface.
func (*MessageHighLatency2) GetID() uint32 {
	return 235
}
