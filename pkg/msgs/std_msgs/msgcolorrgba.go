//nolint:golint
package std_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type ColorRGBA struct {
	msg.Package `ros:"std_msgs"`
	R           float32
	G           float32
	B           float32
	A           float32
}
