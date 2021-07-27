//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type ParamGetReq struct {
	ParamId string
}

type ParamGetRes struct {
	Success bool
	Value   ParamValue
}

type ParamGet struct {
	msg.Package `ros:"mavros_msgs"`
	ParamGetReq
	ParamGetRes
}
