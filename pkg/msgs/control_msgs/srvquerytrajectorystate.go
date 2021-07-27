//nolint:golint
package control_msgs

import (
	"time"

	"github.com/ROBTTO/goroslib/pkg/msg"
)

type QueryTrajectoryStateReq struct {
	Time time.Time
}

type QueryTrajectoryStateRes struct {
	Name         []string
	Position     []float64
	Velocity     []float64
	Acceleration []float64
}

type QueryTrajectoryState struct {
	msg.Package `ros:"control_msgs"`
	QueryTrajectoryStateReq
	QueryTrajectoryStateRes
}
