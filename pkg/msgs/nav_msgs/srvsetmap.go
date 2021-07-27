//nolint:golint
package nav_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
)

type SetMapReq struct {
	Map         OccupancyGrid
	InitialPose geometry_msgs.PoseWithCovarianceStamped
}

type SetMapRes struct {
	Success bool
}

type SetMap struct {
	msg.Package `ros:"nav_msgs"`
	SetMapReq
	SetMapRes
}
