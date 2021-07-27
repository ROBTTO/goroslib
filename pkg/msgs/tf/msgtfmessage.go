//nolint:golint
package tf

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
)

type TfMessage struct {
	msg.Package `ros:"tf"`
	Transforms  []geometry_msgs.TransformStamped
}
