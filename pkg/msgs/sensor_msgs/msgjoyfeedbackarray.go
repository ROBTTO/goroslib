//nolint:golint
package sensor_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type JoyFeedbackArray struct {
	msg.Package `ros:"sensor_msgs"`
	Array       []JoyFeedback
}
