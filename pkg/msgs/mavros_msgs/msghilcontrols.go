//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type HilControls struct {
	msg.Package   `ros:"mavros_msgs"`
	Header        std_msgs.Header
	RollAilerons  float32
	PitchElevator float32
	YawRudder     float32
	Throttle      float32
	Aux1          float32
	Aux2          float32
	Aux3          float32
	Aux4          float32
	Mode          uint8
	NavMode       uint8
}
