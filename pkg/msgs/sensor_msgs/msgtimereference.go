//nolint:golint
package sensor_msgs

import (
	"time"

	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type TimeReference struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	TimeRef     time.Time
	Source      string
}
