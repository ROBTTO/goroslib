//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

const (
	MountConfigureReq_MODE_RETRACT               uint8 = 0
	MountConfigureReq_MODE_NEUTRAL               uint8 = 1
	MountConfigureReq_MODE_MAVLINK_TARGETING     uint8 = 2
	MountConfigureReq_MODE_RC_TARGETING          uint8 = 3
	MountConfigureReq_MODE_GPS_POINT             uint8 = 4
	MountConfigureReq_INPUT_ANGLE_BODY_FRAME     uint8 = 0
	MountConfigureReq_INPUT_ANGULAR_RATE         uint8 = 1
	MountConfigureReq_INPUT_ANGLE_ABSOLUTE_FRAME uint8 = 2
)

type MountConfigureReq struct {
	msg.Definitions `ros:"uint8 MODE_RETRACT=0,uint8 MODE_NEUTRAL=1,uint8 MODE_MAVLINK_TARGETING=2,uint8 MODE_RC_TARGETING=3,uint8 MODE_GPS_POINT=4,uint8 INPUT_ANGLE_BODY_FRAME=0,uint8 INPUT_ANGULAR_RATE=1,uint8 INPUT_ANGLE_ABSOLUTE_FRAME=2"`
	Header          std_msgs.Header
	Mode            uint8
	StabilizeRoll   bool
	StabilizePitch  bool
	StabilizeYaw    bool
	RollInput       uint8
	PitchInput      uint8
	YawInput        uint8
}

type MountConfigureRes struct {
	Success bool
	Result  uint8
}

type MountConfigure struct {
	msg.Package `ros:"mavros_msgs"`
	MountConfigureReq
	MountConfigureRes
}
