//nolint:golint
package geometry_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type TwistWithCovariance struct {
	msg.Package `ros:"geometry_msgs"`
	Twist       Twist
	Covariance  [36]float64
}
