//nolint:golint
package tf2_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

const (
	TF2Error_NO_ERROR               uint8 = 0
	TF2Error_LOOKUP_ERROR           uint8 = 1
	TF2Error_CONNECTIVITY_ERROR     uint8 = 2
	TF2Error_EXTRAPOLATION_ERROR    uint8 = 3
	TF2Error_INVALID_ARGUMENT_ERROR uint8 = 4
	TF2Error_TIMEOUT_ERROR          uint8 = 5
	TF2Error_TRANSFORM_ERROR        uint8 = 6
)

type TF2Error struct {
	msg.Package     `ros:"tf2_msgs"`
	msg.Definitions `ros:"uint8 NO_ERROR=0,uint8 LOOKUP_ERROR=1,uint8 CONNECTIVITY_ERROR=2,uint8 EXTRAPOLATION_ERROR=3,uint8 INVALID_ARGUMENT_ERROR=4,uint8 TIMEOUT_ERROR=5,uint8 TRANSFORM_ERROR=6"`
	Error           uint8
	ErrorString     string
}
