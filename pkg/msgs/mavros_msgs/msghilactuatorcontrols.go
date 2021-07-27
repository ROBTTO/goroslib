//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type HilActuatorControls struct {
	msg.Package `ros:"mavros_msgs"`
	Header      std_msgs.Header
	Controls    [16]float32
	Mode        uint8
	Flags       uint64
}
