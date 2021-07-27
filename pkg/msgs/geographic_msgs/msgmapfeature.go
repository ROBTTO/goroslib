//nolint:golint
package geographic_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/uuid_msgs"
)

type MapFeature struct {
	msg.Package `ros:"geographic_msgs"`
	Id          uuid_msgs.UniqueID
	Components  []uuid_msgs.UniqueID
	Props       []KeyValue
}
