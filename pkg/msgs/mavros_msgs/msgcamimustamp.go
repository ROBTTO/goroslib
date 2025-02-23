//nolint:golint
package mavros_msgs

import (
	"time"

	"github.com/ROBTTO/goroslib/pkg/msg"
)

type CamIMUStamp struct {
	msg.Package `ros:"mavros_msgs"`
	FrameStamp  time.Time
	FrameSeqId  int32
}
