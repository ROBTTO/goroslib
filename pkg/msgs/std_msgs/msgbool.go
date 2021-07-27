//nolint:golint
package std_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type Bool struct {
	msg.Package `ros:"std_msgs"`
	Data        bool
}
