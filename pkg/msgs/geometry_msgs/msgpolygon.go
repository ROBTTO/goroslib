//nolint:golint
package geometry_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type Polygon struct {
	msg.Package `ros:"geometry_msgs"`
	Points      []Point32
}
