//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type LogRequestEndReq struct{}

type LogRequestEndRes struct {
	Success bool
}

type LogRequestEnd struct {
	msg.Package `ros:"mavros_msgs"`
	LogRequestEndReq
	LogRequestEndRes
}
