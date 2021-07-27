//nolint:golint
package std_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type Float32MultiArray struct {
	msg.Package `ros:"std_msgs"`
	Layout      MultiArrayLayout
	Data        []float32
}
