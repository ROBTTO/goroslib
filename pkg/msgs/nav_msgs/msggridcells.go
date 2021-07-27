//nolint:golint
package nav_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type GridCells struct {
	msg.Package `ros:"nav_msgs"`
	Header      std_msgs.Header
	CellWidth   float32
	CellHeight  float32
	Cells       []geometry_msgs.Point
}
