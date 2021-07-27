//nolint:golint
package diagnostic_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type KeyValue struct {
	msg.Package `ros:"diagnostic_msgs"`
	Key         string
	Value       string
}
