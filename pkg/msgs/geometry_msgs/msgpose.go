//nolint:golint
package geometry_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type Pose struct {
	msg.Package `ros:"geometry_msgs"`
	Position    Point
	Orientation Quaternion
}
