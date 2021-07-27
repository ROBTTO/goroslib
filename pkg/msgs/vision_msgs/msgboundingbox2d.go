//nolint:golint
package vision_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
)

type BoundingBox2D struct {
	msg.Package `ros:"vision_msgs"`
	Center      geometry_msgs.Pose2D
	SizeX       float64
	SizeY       float64
}
