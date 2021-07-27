//nolint:golint
package audio_common_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type AudioInfo struct {
	msg.Package  `ros:"audio_common_msgs"`
	Channels     uint8
	SampleRate   uint32
	SampleFormat string
	Bitrate      uint32
	CodingFormat string
}
