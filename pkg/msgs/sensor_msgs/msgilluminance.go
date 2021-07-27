//nolint:golint
package sensor_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type Illuminance struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	Illuminance float64
	Variance    float64
}
