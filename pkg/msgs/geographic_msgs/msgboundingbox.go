//nolint:golint
package geographic_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type BoundingBox struct {
	msg.Package `ros:"geographic_msgs"`
	MinPt       GeoPoint
	MaxPt       GeoPoint
}
