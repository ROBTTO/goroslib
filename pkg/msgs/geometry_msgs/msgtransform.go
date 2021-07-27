//nolint:golint
package geometry_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type Transform struct {
	msg.Package `ros:"geometry_msgs"`
	Translation Vector3
	Rotation    Quaternion
}
