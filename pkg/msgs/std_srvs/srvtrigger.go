//nolint:golint
package std_srvs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type TriggerReq struct{}

type TriggerRes struct {
	Success bool
	Message string
}

type Trigger struct {
	msg.Package `ros:"std_srvs"`
	TriggerReq
	TriggerRes
}
