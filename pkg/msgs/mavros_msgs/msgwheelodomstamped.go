//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type WheelOdomStamped struct {
	msg.Package `ros:"mavros_msgs"`
	Header      std_msgs.Header
	Data        []float64
}
