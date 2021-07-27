//nolint:golint
package std_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type MultiArrayDimension struct {
	msg.Package `ros:"std_msgs"`
	Label       string
	Size        uint32
	Stride      uint32
}
