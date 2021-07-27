//nolint:golint
package shape_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
)

type Mesh struct {
	msg.Package `ros:"shape_msgs"`
	Triangles   []MeshTriangle
	Vertices    []geometry_msgs.Point
}
