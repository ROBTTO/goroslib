//nolint:golint
package nav_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
)

type GetPlanReq struct {
	Start     geometry_msgs.PoseStamped
	Goal      geometry_msgs.PoseStamped
	Tolerance float32
}

type GetPlanRes struct {
	Plan Path
}

type GetPlan struct {
	msg.Package `ros:"nav_msgs"`
	GetPlanReq
	GetPlanRes
}
