//nolint:golint
package nav_msgs

import (
	"time"

	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
)

type MapMetaData struct {
	msg.Package `ros:"nav_msgs"`
	MapLoadTime time.Time
	Resolution  float32
	Width       uint32
	Height      uint32
	Origin      geometry_msgs.Pose
}
