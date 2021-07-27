//nolint:golint
package tf2_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
)

type TFMessage struct {
	msg.Package `ros:"tf2_msgs"`
	Transforms  []geometry_msgs.TransformStamped
}
