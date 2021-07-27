//nolint:golint
package vision_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/sensor_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type Detection3D struct {
	msg.Package `ros:"vision_msgs"`
	Header      std_msgs.Header
	Results     []ObjectHypothesisWithPose
	Bbox        BoundingBox3D
	SourceCloud sensor_msgs.PointCloud2
}
