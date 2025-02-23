//nolint:golint
package std_msgs

import (
	"time"

	"github.com/ROBTTO/goroslib/pkg/msg"
)

type Header struct {
	msg.Package `ros:"std_msgs"`
	Seq         uint32
	Stamp       time.Time
	FrameId     string
}
