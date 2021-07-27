//nolint:golint
package geometry_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type Point32 struct {
	msg.Package `ros:"geometry_msgs"`
	X           float32
	Y           float32
	Z           float32
}
