//nolint:golint
package shape_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type Plane struct {
	msg.Package `ros:"shape_msgs"`
	Coef        [4]float64
}
