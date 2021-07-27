//nolint:golint
package sensor_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type MultiDOFJointState struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	JointNames  []string
	Transforms  []geometry_msgs.Transform
	Twist       []geometry_msgs.Twist
	Wrench      []geometry_msgs.Wrench
}
