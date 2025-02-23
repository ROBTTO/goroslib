//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

const (
	StatusText_EMERGENCY uint8 = 0
	StatusText_ALERT     uint8 = 1
	StatusText_CRITICAL  uint8 = 2
	StatusText_ERROR     uint8 = 3
	StatusText_WARNING   uint8 = 4
	StatusText_NOTICE    uint8 = 5
	StatusText_INFO      uint8 = 6
	StatusText_DEBUG     uint8 = 7
)

type StatusText struct {
	msg.Package     `ros:"mavros_msgs"`
	msg.Definitions `ros:"uint8 EMERGENCY=0,uint8 ALERT=1,uint8 CRITICAL=2,uint8 ERROR=3,uint8 WARNING=4,uint8 NOTICE=5,uint8 INFO=6,uint8 DEBUG=7"`
	Header          std_msgs.Header
	Severity        uint8
	Text            string
}
