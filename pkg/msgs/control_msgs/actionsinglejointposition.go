//nolint:golint
package control_msgs

import (
	"time"

	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type SingleJointPositionActionGoal struct {
	Position    float64
	MinDuration time.Duration
	MaxVelocity float64
}

type SingleJointPositionActionResult struct{}

type SingleJointPositionActionFeedback struct {
	Header   std_msgs.Header
	Position float64
	Velocity float64
	Error    float64
}

type SingleJointPositionAction struct {
	msg.Package `ros:"control_msgs"`
	SingleJointPositionActionGoal
	SingleJointPositionActionResult
	SingleJointPositionActionFeedback
}
