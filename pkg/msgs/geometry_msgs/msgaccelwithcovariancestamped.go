//nolint:golint
package geometry_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type AccelWithCovarianceStamped struct {
	msg.Package `ros:"geometry_msgs"`
	Header      std_msgs.Header
	Accel       AccelWithCovariance
}
