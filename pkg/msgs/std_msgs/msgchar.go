//nolint:golint
package std_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type Char struct {
	msg.Package `ros:"std_msgs"`
	Data        uint8 `rostype:"char"`
}
