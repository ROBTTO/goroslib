//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type RCIn struct {
	msg.Package `ros:"mavros_msgs"`
	Header      std_msgs.Header
	Rssi        uint8
	Channels    []uint16
}
