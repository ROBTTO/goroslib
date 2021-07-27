//nolint:golint
package control_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type JointTolerance struct {
	msg.Package  `ros:"control_msgs"`
	Name         string
	Position     float64
	Velocity     float64
	Acceleration float64
}
