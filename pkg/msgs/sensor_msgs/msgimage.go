//nolint:golint
package sensor_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type Image struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	Height      uint32
	Width       uint32
	Encoding    string
	IsBigendian uint8
	Step        uint32
	Data        []uint8
}
