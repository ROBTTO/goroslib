//nolint:golint
package nav_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type Path struct {
	msg.Package `ros:"nav_msgs"`
	Header      std_msgs.Header
	Poses       []geometry_msgs.PoseStamped
}
