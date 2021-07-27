//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

const (
	AttitudeTarget_IGNORE_ROLL_RATE  uint8 = 1
	AttitudeTarget_IGNORE_PITCH_RATE uint8 = 2
	AttitudeTarget_IGNORE_YAW_RATE   uint8 = 4
	AttitudeTarget_IGNORE_THRUST     uint8 = 64
	AttitudeTarget_IGNORE_ATTITUDE   uint8 = 128
)

type AttitudeTarget struct {
	msg.Package     `ros:"mavros_msgs"`
	msg.Definitions `ros:"uint8 IGNORE_ROLL_RATE=1,uint8 IGNORE_PITCH_RATE=2,uint8 IGNORE_YAW_RATE=4,uint8 IGNORE_THRUST=64,uint8 IGNORE_ATTITUDE=128"`
	Header          std_msgs.Header
	TypeMask        uint8
	Orientation     geometry_msgs.Quaternion
	BodyRate        geometry_msgs.Vector3
	Thrust          float32
}
