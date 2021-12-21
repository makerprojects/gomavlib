//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ardupilotmega

import (
	"errors"
)

type GOPRO_BURST_RATE int

const (
	// 3 Shots / 1 Second.
	GOPRO_BURST_RATE_3_IN_1_SECOND GOPRO_BURST_RATE = 0
	// 5 Shots / 1 Second.
	GOPRO_BURST_RATE_5_IN_1_SECOND GOPRO_BURST_RATE = 1
	// 10 Shots / 1 Second.
	GOPRO_BURST_RATE_10_IN_1_SECOND GOPRO_BURST_RATE = 2
	// 10 Shots / 2 Second.
	GOPRO_BURST_RATE_10_IN_2_SECOND GOPRO_BURST_RATE = 3
	// 10 Shots / 3 Second (Hero 4 Only).
	GOPRO_BURST_RATE_10_IN_3_SECOND GOPRO_BURST_RATE = 4
	// 30 Shots / 1 Second.
	GOPRO_BURST_RATE_30_IN_1_SECOND GOPRO_BURST_RATE = 5
	// 30 Shots / 2 Second.
	GOPRO_BURST_RATE_30_IN_2_SECOND GOPRO_BURST_RATE = 6
	// 30 Shots / 3 Second.
	GOPRO_BURST_RATE_30_IN_3_SECOND GOPRO_BURST_RATE = 7
	// 30 Shots / 6 Second.
	GOPRO_BURST_RATE_30_IN_6_SECOND GOPRO_BURST_RATE = 8
)

var labels_GOPRO_BURST_RATE = map[GOPRO_BURST_RATE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GOPRO_BURST_RATE) MarshalText() ([]byte, error) {
	if l, ok := labels_GOPRO_BURST_RATE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_GOPRO_BURST_RATE = map[string]GOPRO_BURST_RATE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GOPRO_BURST_RATE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_GOPRO_BURST_RATE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e GOPRO_BURST_RATE) String() string {
	if l, ok := labels_GOPRO_BURST_RATE[e]; ok {
		return l
	}
	return "invalid value"
}
