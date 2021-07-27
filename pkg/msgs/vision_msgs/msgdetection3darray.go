//nolint:golint
package vision_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type Detection3DArray struct {
	msg.Package `ros:"vision_msgs"`
	Header      std_msgs.Header
	Detections  []Detection3D
}
