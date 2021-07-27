//nolint:golint
package geographic_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/uuid_msgs"
)

type RouteSegment struct {
	msg.Package `ros:"geographic_msgs"`
	Id          uuid_msgs.UniqueID
	Start       uuid_msgs.UniqueID
	End         uuid_msgs.UniqueID
	Props       []KeyValue
}
