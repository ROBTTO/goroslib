//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type LogRequestDataReq struct {
	Id     uint16
	Offset uint32
	Count  uint32
}

type LogRequestDataRes struct {
	Success bool
}

type LogRequestData struct {
	msg.Package `ros:"mavros_msgs"`
	LogRequestDataReq
	LogRequestDataRes
}
