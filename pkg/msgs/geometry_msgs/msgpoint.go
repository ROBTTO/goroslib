//nolint:golint
package geometry_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type Point struct {
	msg.Package `ros:"geometry_msgs"`
	X           float64
	Y           float64
	Z           float64
}
