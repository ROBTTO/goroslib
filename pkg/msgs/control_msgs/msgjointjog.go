//nolint:golint
package control_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type JointJog struct {
	msg.Package   `ros:"control_msgs"`
	Header        std_msgs.Header
	JointNames    []string
	Displacements []float64
	Velocities    []float64
	Duration      float64
}
