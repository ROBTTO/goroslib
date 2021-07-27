//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type LogRequestListReq struct {
	Start uint16
	End   uint16
}

type LogRequestListRes struct {
	Success bool
}

type LogRequestList struct {
	msg.Package `ros:"mavros_msgs"`
	LogRequestListReq
	LogRequestListRes
}
