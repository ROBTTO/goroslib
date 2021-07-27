//nolint:golint
package control_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/trajectory_msgs"
)

type JointTrajectoryActionGoal struct {
	Trajectory trajectory_msgs.JointTrajectory
}

type JointTrajectoryActionResult struct{}

type JointTrajectoryActionFeedback struct{}

type JointTrajectoryAction struct {
	msg.Package `ros:"control_msgs"`
	JointTrajectoryActionGoal
	JointTrajectoryActionResult
	JointTrajectoryActionFeedback
}
