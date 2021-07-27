//nolint:golint
package trajectory_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type JointTrajectory struct {
	msg.Package `ros:"trajectory_msgs"`
	Header      std_msgs.Header
	JointNames  []string
	Points      []JointTrajectoryPoint
}
