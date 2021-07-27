//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type GPSRTK struct {
	msg.Package      `ros:"mavros_msgs"`
	Header           std_msgs.Header
	RtkReceiverId    uint8
	Wn               int16
	Tow              uint32
	RtkHealth        uint8
	RtkRate          uint8
	Nsats            uint8
	BaselineA        int32
	BaselineB        int32
	BaselineC        int32
	Accuracy         uint32
	IarNumHypotheses int32
}
