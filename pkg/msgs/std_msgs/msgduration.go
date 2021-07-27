//nolint:golint
package std_msgs

import (
	"time"

	"github.com/ROBTTO/goroslib/pkg/msg"
)

type Duration struct {
	msg.Package `ros:"std_msgs"`
	Data        time.Duration
}
