//nolint:golint
package geographic_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/uuid_msgs"
)

type GeographicMap struct {
	msg.Package `ros:"geographic_msgs"`
	Header      std_msgs.Header
	Id          uuid_msgs.UniqueID
	Bounds      BoundingBox
	Points      []WayPoint
	Features    []MapFeature
	Props       []KeyValue
}
