//nolint:golint
package nav_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type GetMapActionGoal struct{}

type GetMapActionResult struct {
	Map OccupancyGrid
}

type GetMapActionFeedback struct{}

type GetMapAction struct {
	msg.Package `ros:"nav_msgs"`
	GetMapActionGoal
	GetMapActionResult
	GetMapActionFeedback
}
