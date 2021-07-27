//nolint:golint
package std_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type MultiArrayLayout struct {
	msg.Package `ros:"std_msgs"`
	Dim         []MultiArrayDimension
	DataOffset  uint32
}
