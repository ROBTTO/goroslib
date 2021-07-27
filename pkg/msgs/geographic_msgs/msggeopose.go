//nolint:golint
package geographic_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
)

type GeoPose struct {
	msg.Package `ros:"geographic_msgs"`
	Position    GeoPoint
	Orientation geometry_msgs.Quaternion
}
