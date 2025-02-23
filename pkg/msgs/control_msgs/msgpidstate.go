//nolint:golint
package control_msgs

import (
	"time"

	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type PidState struct {
	msg.Package `ros:"control_msgs"`
	Header      std_msgs.Header
	Timestep    time.Duration
	Error       float64
	ErrorDot    float64
	PError      float64
	IError      float64
	DError      float64
	PTerm       float64
	ITerm       float64
	DTerm       float64
	IMax        float64
	IMin        float64
	Output      float64
}
