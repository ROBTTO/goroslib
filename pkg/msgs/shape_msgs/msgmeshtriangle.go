//nolint:golint
package shape_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type MeshTriangle struct {
	msg.Package   `ros:"shape_msgs"`
	VertexIndices [3]uint32
}
