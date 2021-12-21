//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package ualberta

import (
	"errors"
)

// Camera capability flags (Bitmap)
type CAMERA_CAP_FLAGS int

const (
	// Camera is able to record video
	CAMERA_CAP_FLAGS_CAPTURE_VIDEO CAMERA_CAP_FLAGS = 1
	// Camera is able to capture images
	CAMERA_CAP_FLAGS_CAPTURE_IMAGE CAMERA_CAP_FLAGS = 2
	// Camera has separate Video and Image/Photo modes (MAV_CMD_SET_CAMERA_MODE)
	CAMERA_CAP_FLAGS_HAS_MODES CAMERA_CAP_FLAGS = 4
	// Camera can capture images while in video mode
	CAMERA_CAP_FLAGS_CAN_CAPTURE_IMAGE_IN_VIDEO_MODE CAMERA_CAP_FLAGS = 8
	// Camera can capture videos while in Photo/Image mode
	CAMERA_CAP_FLAGS_CAN_CAPTURE_VIDEO_IN_IMAGE_MODE CAMERA_CAP_FLAGS = 16
	// Camera has image survey mode (MAV_CMD_SET_CAMERA_MODE)
	CAMERA_CAP_FLAGS_HAS_IMAGE_SURVEY_MODE CAMERA_CAP_FLAGS = 32
	// Camera has basic zoom control (MAV_CMD_SET_CAMERA_ZOOM)
	CAMERA_CAP_FLAGS_HAS_BASIC_ZOOM CAMERA_CAP_FLAGS = 64
	// Camera has basic focus control (MAV_CMD_SET_CAMERA_FOCUS)
	CAMERA_CAP_FLAGS_HAS_BASIC_FOCUS CAMERA_CAP_FLAGS = 128
	// Camera has video streaming capabilities (request VIDEO_STREAM_INFORMATION with MAV_CMD_REQUEST_MESSAGE for video streaming info)
	CAMERA_CAP_FLAGS_HAS_VIDEO_STREAM CAMERA_CAP_FLAGS = 256
	// Camera supports tracking of a point on the camera view.
	CAMERA_CAP_FLAGS_HAS_TRACKING_POINT CAMERA_CAP_FLAGS = 512
	// Camera supports tracking of a selection rectangle on the camera view.
	CAMERA_CAP_FLAGS_HAS_TRACKING_RECTANGLE CAMERA_CAP_FLAGS = 1024
	// Camera supports tracking geo status (CAMERA_TRACKING_GEO_STATUS).
	CAMERA_CAP_FLAGS_HAS_TRACKING_GEO_STATUS CAMERA_CAP_FLAGS = 2048
)

var labels_CAMERA_CAP_FLAGS = map[CAMERA_CAP_FLAGS]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e CAMERA_CAP_FLAGS) MarshalText() ([]byte, error) {
	if l, ok := labels_CAMERA_CAP_FLAGS[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_CAMERA_CAP_FLAGS = map[string]CAMERA_CAP_FLAGS{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *CAMERA_CAP_FLAGS) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_CAMERA_CAP_FLAGS[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e CAMERA_CAP_FLAGS) String() string {
	if l, ok := labels_CAMERA_CAP_FLAGS[e]; ok {
		return l
	}
	return "invalid value"
}
