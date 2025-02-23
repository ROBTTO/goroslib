//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

const (
	StreamRateReq_STREAM_ALL             uint8 = 0
	StreamRateReq_STREAM_RAW_SENSORS     uint8 = 1
	StreamRateReq_STREAM_EXTENDED_STATUS uint8 = 2
	StreamRateReq_STREAM_RC_CHANNELS     uint8 = 3
	StreamRateReq_STREAM_RAW_CONTROLLER  uint8 = 4
	StreamRateReq_STREAM_POSITION        uint8 = 6
	StreamRateReq_STREAM_EXTRA1          uint8 = 10
	StreamRateReq_STREAM_EXTRA2          uint8 = 11
	StreamRateReq_STREAM_EXTRA3          uint8 = 12
)

type StreamRateReq struct {
	msg.Definitions `ros:"uint8 STREAM_ALL=0,uint8 STREAM_RAW_SENSORS=1,uint8 STREAM_EXTENDED_STATUS=2,uint8 STREAM_RC_CHANNELS=3,uint8 STREAM_RAW_CONTROLLER=4,uint8 STREAM_POSITION=6,uint8 STREAM_EXTRA1=10,uint8 STREAM_EXTRA2=11,uint8 STREAM_EXTRA3=12"`
	StreamId        uint8
	MessageRate     uint16
	OnOff           bool
}

type StreamRateRes struct{}

type StreamRate struct {
	msg.Package `ros:"mavros_msgs"`
	StreamRateReq
	StreamRateRes
}
