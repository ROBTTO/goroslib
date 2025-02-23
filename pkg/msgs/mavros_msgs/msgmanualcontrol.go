//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type ManualControl struct {
	msg.Package `ros:"mavros_msgs"`
	Header      std_msgs.Header
	X           float32
	Y           float32
	Z           float32
	R           float32
	Buttons     uint16
}
