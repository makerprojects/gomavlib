//autogenerated:yes
//nolint:revive,misspell,govet,lll
package paparazzi

// Control vehicle tone generation (buzzer).
type MessagePlayTune struct {
	// System ID
	TargetSystem uint8
	// Component ID
	TargetComponent uint8
	// tune in board specific format
	Tune string `mavlen:"30"`
	// tune extension (appended to tune)
	Tune2 string `mavext:"true" mavlen:"200"`
}

// GetID implements the msg.Message interface.
func (*MessagePlayTune) GetID() uint32 {
	return 258
}
