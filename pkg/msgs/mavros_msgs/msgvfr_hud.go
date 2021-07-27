//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type VFR_HUD struct {
	msg.Package `ros:"mavros_msgs"`
	Header      std_msgs.Header
	Airspeed    float32
	Groundspeed float32
	Heading     int16
	Throttle    float32
	Altitude    float32
	Climb       float32
}
