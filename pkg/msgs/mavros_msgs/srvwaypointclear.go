//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type WaypointClearReq struct{}

type WaypointClearRes struct {
	Success bool
}

type WaypointClear struct {
	msg.Package `ros:"mavros_msgs"`
	WaypointClearReq
	WaypointClearRes
}
