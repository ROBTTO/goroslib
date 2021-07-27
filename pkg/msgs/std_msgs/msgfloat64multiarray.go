//nolint:golint
package std_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type Float64MultiArray struct {
	msg.Package `ros:"std_msgs"`
	Layout      MultiArrayLayout
	Data        []float64
}
