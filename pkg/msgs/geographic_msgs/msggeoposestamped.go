//nolint:golint
package geographic_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type GeoPoseStamped struct {
	msg.Package `ros:"geographic_msgs"`
	Header      std_msgs.Header
	Pose        GeoPose
}
