//nolint:golint
package geometry_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type Twist struct {
	msg.Package `ros:"geometry_msgs"`
	Linear      Vector3
	Angular     Vector3
}
