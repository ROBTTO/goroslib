//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geographic_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type HilStateQuaternion struct {
	msg.Package        `ros:"mavros_msgs"`
	Header             std_msgs.Header
	Orientation        geometry_msgs.Quaternion
	AngularVelocity    geometry_msgs.Vector3
	LinearAcceleration geometry_msgs.Vector3
	LinearVelocity     geometry_msgs.Vector3
	Geo                geographic_msgs.GeoPoint
	IndAirspeed        float32
	TrueAirspeed       float32
}
