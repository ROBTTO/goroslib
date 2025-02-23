//nolint:golint
package control_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type JointControllerState struct {
	msg.Package     `ros:"control_msgs"`
	Header          std_msgs.Header
	SetPoint        float64
	ProcessValue    float64
	ProcessValueDot float64
	Error           float64
	TimeStep        float64
	Command         float64
	P               float64
	I               float64
	D               float64
	IClamp          float64
	Antiwindup      bool
}
