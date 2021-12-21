//autogenerated:yes
//nolint:revive,misspell,govet,lll
package pythonarraytest

// The current system altitude.
type MessageAltitude struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// This altitude measure is initialized on system boot and monotonic (it is never reset, but represents the local altitude change). The only guarantee on this field is that it will never be reset and is consistent within a flight. The recommended value for this field is the uncorrected barometric altitude at boot time. This altitude will also drift and vary between flights.
	AltitudeMonotonic float32
	// This altitude measure is strictly above mean sea level and might be non-monotonic (it might reset on events like GPS lock or when a new QNH value is set). It should be the altitude to which global altitude waypoints are compared to. Note that it is *not* the GPS altitude, however, most GPS modules already output MSL by default and not the WGS84 altitude.
	AltitudeAmsl float32
	// This is the local altitude in the local coordinate frame. It is not the altitude above home, but in reference to the coordinate origin (0, 0, 0). It is up-positive.
	AltitudeLocal float32
	// This is the altitude above the home position. It resets on each change of the current home position.
	AltitudeRelative float32
	// This is the altitude above terrain. It might be fed by a terrain database or an altimeter. Values smaller than -1000 should be interpreted as unknown.
	AltitudeTerrain float32
	// This is not the altitude, but the clear space below the system according to the fused clearance estimate. It generally should max out at the maximum range of e.g. the laser altimeter. It is generally a moving target. A negative value indicates no measurement available.
	BottomClearance float32
}

// GetID implements the msg.Message interface.
func (*MessageAltitude) GetID() uint32 {
	return 141
}
