//nolint:golint
package visualization_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
)

const (
	InteractiveMarkerControl_INHERIT        uint8 = 0
	InteractiveMarkerControl_FIXED          uint8 = 1
	InteractiveMarkerControl_VIEW_FACING    uint8 = 2
	InteractiveMarkerControl_NONE           uint8 = 0
	InteractiveMarkerControl_MENU           uint8 = 1
	InteractiveMarkerControl_BUTTON         uint8 = 2
	InteractiveMarkerControl_MOVE_AXIS      uint8 = 3
	InteractiveMarkerControl_MOVE_PLANE     uint8 = 4
	InteractiveMarkerControl_ROTATE_AXIS    uint8 = 5
	InteractiveMarkerControl_MOVE_ROTATE    uint8 = 6
	InteractiveMarkerControl_MOVE_3D        uint8 = 7
	InteractiveMarkerControl_ROTATE_3D      uint8 = 8
	InteractiveMarkerControl_MOVE_ROTATE_3D uint8 = 9
)

type InteractiveMarkerControl struct {
	msg.Package                  `ros:"visualization_msgs"`
	msg.Definitions              `ros:"uint8 INHERIT=0,uint8 FIXED=1,uint8 VIEW_FACING=2,uint8 NONE=0,uint8 MENU=1,uint8 BUTTON=2,uint8 MOVE_AXIS=3,uint8 MOVE_PLANE=4,uint8 ROTATE_AXIS=5,uint8 MOVE_ROTATE=6,uint8 MOVE_3D=7,uint8 ROTATE_3D=8,uint8 MOVE_ROTATE_3D=9"`
	Name                         string
	Orientation                  geometry_msgs.Quaternion
	OrientationMode              uint8
	InteractionMode              uint8
	AlwaysVisible                bool
	Markers                      []Marker
	IndependentMarkerOrientation bool
	Description                  string
}
