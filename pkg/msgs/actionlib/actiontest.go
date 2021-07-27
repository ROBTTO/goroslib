//nolint:golint
package actionlib

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type TestActionGoal struct {
	Goal int32
}

type TestActionResult struct {
	Result int32
}

type TestActionFeedback struct {
	Feedback int32
}

type TestAction struct {
	msg.Package `ros:"actionlib"`
	TestActionGoal
	TestActionResult
	TestActionFeedback
}
