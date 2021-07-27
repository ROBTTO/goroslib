//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type WaypointPullReq struct{}

type WaypointPullRes struct {
	Success    bool
	WpReceived uint32
}

type WaypointPull struct {
	msg.Package `ros:"mavros_msgs"`
	WaypointPullReq
	WaypointPullRes
}
