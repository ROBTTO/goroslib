//nolint:golint
package nav_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type OccupancyGrid struct {
	msg.Package `ros:"nav_msgs"`
	Header      std_msgs.Header
	Info        MapMetaData
	Data        []int8
}
