//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type WaypointSetCurrentReq struct {
	WpSeq uint16
}

type WaypointSetCurrentRes struct {
	Success bool
}

type WaypointSetCurrent struct {
	msg.Package `ros:"mavros_msgs"`
	WaypointSetCurrentReq
	WaypointSetCurrentRes
}
