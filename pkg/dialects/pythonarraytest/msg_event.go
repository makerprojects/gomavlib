//autogenerated:yes
//nolint:revive,misspell,govet,lll
package pythonarraytest

// Event message. Each new event from a particular component gets a new sequence number. The same message might be sent multiple times if (re-)requested. Most events are broadcast, some can be specific to a target component (as receivers keep track of the sequence for missed events, all events need to be broadcast. Thus we use destination_component instead of target_component).
type MessageEvent struct {
	// Component ID
	DestinationComponent uint8
	// System ID
	DestinationSystem uint8
	// Event ID (as defined in the component metadata)
	Id uint32
	// Timestamp (time since system boot when the event happened).
	EventTimeBootMs uint32
	// Sequence number.
	Sequence uint16
	// Log levels: 4 bits MSB: internal (for logging purposes), 4 bits LSB: external. Levels: Emergency = 0, Alert = 1, Critical = 2, Error = 3, Warning = 4, Notice = 5, Info = 6, Debug = 7, Protocol = 8, Disabled = 9
	LogLevels uint8
	// Arguments (depend on event ID).
	Arguments [40]uint8
}

// GetID implements the msg.Message interface.
func (*MessageEvent) GetID() uint32 {
	return 410
}
