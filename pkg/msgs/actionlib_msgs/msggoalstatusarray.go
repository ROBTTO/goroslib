//nolint:golint
package actionlib_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type GoalStatusArray struct {
	msg.Package `ros:"actionlib_msgs"`
	Header      std_msgs.Header
	StatusList  []GoalStatus
}
