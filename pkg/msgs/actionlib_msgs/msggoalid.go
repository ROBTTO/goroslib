//nolint:golint
package actionlib_msgs

import (
	"time"

	"github.com/ROBTTO/goroslib/pkg/msg"
)

type GoalID struct {
	msg.Package `ros:"actionlib_msgs"`
	Stamp       time.Time
	Id          string
}
