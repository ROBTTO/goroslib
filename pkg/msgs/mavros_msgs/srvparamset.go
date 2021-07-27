//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type ParamSetReq struct {
	ParamId string
	Value   ParamValue
}

type ParamSetRes struct {
	Success bool
	Value   ParamValue
}

type ParamSet struct {
	msg.Package `ros:"mavros_msgs"`
	ParamSetReq
	ParamSetRes
}
