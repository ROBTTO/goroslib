//nolint:golint
package trajectory_msgs

import (
	"time"

	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
)

type MultiDOFJointTrajectoryPoint struct {
	msg.Package   `ros:"trajectory_msgs"`
	Transforms    []geometry_msgs.Transform
	Velocities    []geometry_msgs.Twist
	Accelerations []geometry_msgs.Twist
	TimeFromStart time.Duration
}
