//nolint:golint
package std_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type Int64 struct {
	msg.Package `ros:"std_msgs"`
	Data        int64
}
