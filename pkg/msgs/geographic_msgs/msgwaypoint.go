//nolint:golint
package geographic_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/uuid_msgs"
)

type WayPoint struct {
	msg.Package `ros:"geographic_msgs"`
	Id          uuid_msgs.UniqueID
	Position    GeoPoint
	Props       []KeyValue
}
