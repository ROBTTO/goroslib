//nolint:golint
package control_msgs

import (
	"time"

	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/trajectory_msgs"
)

type FollowJointTrajectoryActionGoal struct {
	Trajectory        trajectory_msgs.JointTrajectory
	PathTolerance     []JointTolerance
	GoalTolerance     []JointTolerance
	GoalTimeTolerance time.Duration
}

const (
	FollowJointTrajectoryActionResult_SUCCESSFUL              int32 = 0
	FollowJointTrajectoryActionResult_INVALID_GOAL            int32 = -1
	FollowJointTrajectoryActionResult_INVALID_JOINTS          int32 = -2
	FollowJointTrajectoryActionResult_OLD_HEADER_TIMESTAMP    int32 = -3
	FollowJointTrajectoryActionResult_PATH_TOLERANCE_VIOLATED int32 = -4
	FollowJointTrajectoryActionResult_GOAL_TOLERANCE_VIOLATED int32 = -5
)

type FollowJointTrajectoryActionResult struct {
	msg.Definitions `ros:"int32 SUCCESSFUL=0,int32 INVALID_GOAL=-1,int32 INVALID_JOINTS=-2,int32 OLD_HEADER_TIMESTAMP=-3,int32 PATH_TOLERANCE_VIOLATED=-4,int32 GOAL_TOLERANCE_VIOLATED=-5"`
	ErrorCode       int32
	ErrorString     string
}

type FollowJointTrajectoryActionFeedback struct {
	Header     std_msgs.Header
	JointNames []string
	Desired    trajectory_msgs.JointTrajectoryPoint
	Actual     trajectory_msgs.JointTrajectoryPoint
	Error      trajectory_msgs.JointTrajectoryPoint
}

type FollowJointTrajectoryAction struct {
	msg.Package `ros:"control_msgs"`
	FollowJointTrajectoryActionGoal
	FollowJointTrajectoryActionResult
	FollowJointTrajectoryActionFeedback
}
