//nolint:golint
package velodyne_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type VelodyneScan struct {
	msg.Package `ros:"velodyne_msgs"`
	Header      std_msgs.Header
	Packets     []VelodynePacket
}
