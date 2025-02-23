//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geometry_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

const (
	GlobalPositionTarget_FRAME_GLOBAL_INT         uint8  = 5
	GlobalPositionTarget_FRAME_GLOBAL_REL_ALT     uint8  = 6
	GlobalPositionTarget_FRAME_GLOBAL_TERRAIN_ALT uint8  = 11
	GlobalPositionTarget_IGNORE_LATITUDE          uint16 = 1
	GlobalPositionTarget_IGNORE_LONGITUDE         uint16 = 2
	GlobalPositionTarget_IGNORE_ALTITUDE          uint16 = 4
	GlobalPositionTarget_IGNORE_VX                uint16 = 8
	GlobalPositionTarget_IGNORE_VY                uint16 = 16
	GlobalPositionTarget_IGNORE_VZ                uint16 = 32
	GlobalPositionTarget_IGNORE_AFX               uint16 = 64
	GlobalPositionTarget_IGNORE_AFY               uint16 = 128
	GlobalPositionTarget_IGNORE_AFZ               uint16 = 256
	GlobalPositionTarget_FORCE                    uint16 = 512
	GlobalPositionTarget_IGNORE_YAW               uint16 = 1024
	GlobalPositionTarget_IGNORE_YAW_RATE          uint16 = 2048
)

type GlobalPositionTarget struct {
	msg.Package         `ros:"mavros_msgs"`
	msg.Definitions     `ros:"uint8 FRAME_GLOBAL_INT=5,uint8 FRAME_GLOBAL_REL_ALT=6,uint8 FRAME_GLOBAL_TERRAIN_ALT=11,uint16 IGNORE_LATITUDE=1,uint16 IGNORE_LONGITUDE=2,uint16 IGNORE_ALTITUDE=4,uint16 IGNORE_VX=8,uint16 IGNORE_VY=16,uint16 IGNORE_VZ=32,uint16 IGNORE_AFX=64,uint16 IGNORE_AFY=128,uint16 IGNORE_AFZ=256,uint16 FORCE=512,uint16 IGNORE_YAW=1024,uint16 IGNORE_YAW_RATE=2048"`
	Header              std_msgs.Header
	CoordinateFrame     uint8
	TypeMask            uint16
	Latitude            float64
	Longitude           float64
	Altitude            float32
	Velocity            geometry_msgs.Vector3
	AccelerationOrForce geometry_msgs.Vector3
	Yaw                 float32
	YawRate             float32
}
