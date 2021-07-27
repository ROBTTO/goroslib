//nolint:golint
package geographic_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/uuid_msgs"
)

type GetRoutePlanReq struct {
	Network uuid_msgs.UniqueID
	Start   uuid_msgs.UniqueID
	Goal    uuid_msgs.UniqueID
}

type GetRoutePlanRes struct {
	Success bool
	Status  string
	Plan    RoutePath
}

type GetRoutePlan struct {
	msg.Package `ros:"geographic_msgs"`
	GetRoutePlanReq
	GetRoutePlanRes
}
