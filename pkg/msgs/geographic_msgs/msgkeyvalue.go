//nolint:golint
package geographic_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type KeyValue struct {
	msg.Package `ros:"geographic_msgs"`
	Key         string
	Value       string
}
