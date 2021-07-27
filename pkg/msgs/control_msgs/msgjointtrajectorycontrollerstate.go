//nolint:golint
package control_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/trajectory_msgs"
)

type JointTrajectoryControllerState struct {
	msg.Package `ros:"control_msgs"`
	Header      std_msgs.Header
	JointNames  []string
	Desired     trajectory_msgs.JointTrajectoryPoint
	Actual      trajectory_msgs.JointTrajectoryPoint
	Error       trajectory_msgs.JointTrajectoryPoint
}
