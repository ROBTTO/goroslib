//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

const (
	ActuatorControl_PX4_MIX_FLIGHT_CONTROL          uint8 = 0
	ActuatorControl_PX4_MIX_FLIGHT_CONTROL_VTOL_ALT uint8 = 1
	ActuatorControl_PX4_MIX_PAYLOAD                 uint8 = 2
	ActuatorControl_PX4_MIX_MANUAL_PASSTHROUGH      uint8 = 3
)

type ActuatorControl struct {
	msg.Package     `ros:"mavros_msgs"`
	msg.Definitions `ros:"uint8 PX4_MIX_FLIGHT_CONTROL=0,uint8 PX4_MIX_FLIGHT_CONTROL_VTOL_ALT=1,uint8 PX4_MIX_PAYLOAD=2,uint8 PX4_MIX_MANUAL_PASSTHROUGH=3"`
	Header          std_msgs.Header
	GroupMix        uint8
	Controls        [8]float32
}
