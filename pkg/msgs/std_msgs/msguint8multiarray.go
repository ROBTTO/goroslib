//nolint:golint
package std_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type UInt8MultiArray struct {
	msg.Package `ros:"std_msgs"`
	Layout      MultiArrayLayout
	Data        []uint8
}
