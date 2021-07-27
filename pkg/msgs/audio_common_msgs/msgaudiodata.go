//nolint:golint
package audio_common_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type AudioData struct {
	msg.Package `ros:"audio_common_msgs"`
	Data        []uint8
}
