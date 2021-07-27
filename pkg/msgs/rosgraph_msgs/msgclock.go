//nolint:golint
package rosgraph_msgs

import (
	"time"

	"github.com/ROBTTO/goroslib/pkg/msg"
)

type Clock struct {
	msg.Package `ros:"rosgraph_msgs"`
	Clock       time.Time
}
