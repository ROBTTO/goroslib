//nolint:golint
package std_msgs

import (
	"time"

	"github.com/ROBTTO/goroslib/pkg/msg"
)

type Time struct {
	msg.Package `ros:"std_msgs"`
	Data        time.Time
}
