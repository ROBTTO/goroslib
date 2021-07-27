//nolint:golint
package geometry_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type AccelWithCovariance struct {
	msg.Package `ros:"geometry_msgs"`
	Accel       Accel
	Covariance  [36]float64
}
