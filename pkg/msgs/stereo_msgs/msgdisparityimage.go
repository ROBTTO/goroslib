//nolint:golint
package stereo_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/sensor_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type DisparityImage struct {
	msg.Package  `ros:"stereo_msgs"`
	Header       std_msgs.Header
	Image        sensor_msgs.Image
	F            float32
	T            float32 `rosname:"T"`
	ValidWindow  sensor_msgs.RegionOfInterest
	MinDisparity float32
	MaxDisparity float32
	DeltaD       float32
}
