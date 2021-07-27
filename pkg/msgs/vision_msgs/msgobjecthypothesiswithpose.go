//nolint:golint
package vision_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
)

type ObjectHypothesisWithPose struct {
	msg.Package `ros:"vision_msgs"`
	Id          int64
	Score       float64
	Pose        geometry_msgs.PoseWithCovariance
}
