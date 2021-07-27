//nolint:golint
package geometry_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type PoseWithCovariance struct {
	msg.Package `ros:"geometry_msgs"`
	Pose        Pose
	Covariance  [36]float64
}
