//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geographic_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type HomePosition struct {
	msg.Package `ros:"mavros_msgs"`
	Header      std_msgs.Header
	Geo         geographic_msgs.GeoPoint
	Position    geometry_msgs.Point
	Orientation geometry_msgs.Quaternion
	Approach    geometry_msgs.Vector3
}
