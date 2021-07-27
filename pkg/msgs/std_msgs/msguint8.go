//nolint:golint
package std_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type UInt8 struct {
	msg.Package `ros:"std_msgs"`
	Data        uint8
}
