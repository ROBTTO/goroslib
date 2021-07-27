//nolint:golint
package sensor_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type PointCloud struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	Points      []geometry_msgs.Point32
	Channels    []ChannelFloat32
}
