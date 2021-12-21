//autogenerated:yes
//nolint:revive,misspell,govet,lll
package matrixpilot

// Backwards compatible version of SERIAL_UDB_EXTRA F8: format
type MessageSerialUdbExtraF8 struct {
	// Serial UDB Extra HEIGHT_TARGET_MAX
	SueHeightTargetMax float32 `mavname:"sue_HEIGHT_TARGET_MAX"`
	// Serial UDB Extra HEIGHT_TARGET_MIN
	SueHeightTargetMin float32 `mavname:"sue_HEIGHT_TARGET_MIN"`
	// Serial UDB Extra ALT_HOLD_THROTTLE_MIN
	SueAltHoldThrottleMin float32 `mavname:"sue_ALT_HOLD_THROTTLE_MIN"`
	// Serial UDB Extra ALT_HOLD_THROTTLE_MAX
	SueAltHoldThrottleMax float32 `mavname:"sue_ALT_HOLD_THROTTLE_MAX"`
	// Serial UDB Extra ALT_HOLD_PITCH_MIN
	SueAltHoldPitchMin float32 `mavname:"sue_ALT_HOLD_PITCH_MIN"`
	// Serial UDB Extra ALT_HOLD_PITCH_MAX
	SueAltHoldPitchMax float32 `mavname:"sue_ALT_HOLD_PITCH_MAX"`
	// Serial UDB Extra ALT_HOLD_PITCH_HIGH
	SueAltHoldPitchHigh float32 `mavname:"sue_ALT_HOLD_PITCH_HIGH"`
}

// GetID implements the msg.Message interface.
func (*MessageSerialUdbExtraF8) GetID() uint32 {
	return 176
}
