//nolint:golint
package visualization_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

const (
	InteractiveMarkerFeedback_KEEP_ALIVE   uint8 = 0
	InteractiveMarkerFeedback_POSE_UPDATE  uint8 = 1
	InteractiveMarkerFeedback_MENU_SELECT  uint8 = 2
	InteractiveMarkerFeedback_BUTTON_CLICK uint8 = 3
	InteractiveMarkerFeedback_MOUSE_DOWN   uint8 = 4
	InteractiveMarkerFeedback_MOUSE_UP     uint8 = 5
)

type InteractiveMarkerFeedback struct {
	msg.Package     `ros:"visualization_msgs"`
	msg.Definitions `ros:"uint8 KEEP_ALIVE=0,uint8 POSE_UPDATE=1,uint8 MENU_SELECT=2,uint8 BUTTON_CLICK=3,uint8 MOUSE_DOWN=4,uint8 MOUSE_UP=5"`
	Header          std_msgs.Header
	ClientId        string
	MarkerName      string
	ControlName     string
	EventType       uint8
	Pose            geometry_msgs.Pose
	MenuEntryId     uint32
	MousePoint      geometry_msgs.Point
	MousePointValid bool
}
