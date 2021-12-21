//autogenerated:yes
//nolint:revive,misspell,govet,lll
package common

// Estimator status message including flags, innovation test ratios and estimated accuracies. The flags message is an integer bitmask containing information on which EKF outputs are valid. See the ESTIMATOR_STATUS_FLAGS enum definition for further information. The innovation test ratios show the magnitude of the sensor innovation divided by the innovation check threshold. Under normal operation the innovation test ratios should be below 0.5 with occasional values up to 1.0. Values greater than 1.0 should be rare under normal operation and indicate that a measurement has been rejected by the filter. The user should be notified if an innovation test ratio greater than 1.0 is recorded. Notifications for values in the range between 0.5 and 1.0 should be optional and controllable by the user.
type MessageEstimatorStatus struct {
	// Timestamp (UNIX Epoch time or time since system boot). The receiving end can infer timestamp format (since 1.1.1970 or since system boot) by checking for the magnitude of the number.
	TimeUsec uint64
	// Bitmap indicating which EKF outputs are valid.
	Flags ESTIMATOR_STATUS_FLAGS `mavenum:"uint16"`
	// Velocity innovation test ratio
	VelRatio float32
	// Horizontal position innovation test ratio
	PosHorizRatio float32
	// Vertical position innovation test ratio
	PosVertRatio float32
	// Magnetometer innovation test ratio
	MagRatio float32
	// Height above terrain innovation test ratio
	HaglRatio float32
	// True airspeed innovation test ratio
	TasRatio float32
	// Horizontal position 1-STD accuracy relative to the EKF local origin
	PosHorizAccuracy float32
	// Vertical position 1-STD accuracy relative to the EKF local origin
	PosVertAccuracy float32
}

// GetID implements the msg.Message interface.
func (*MessageEstimatorStatus) GetID() uint32 {
	return 230
}
