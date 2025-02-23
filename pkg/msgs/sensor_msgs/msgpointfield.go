//nolint:golint
package sensor_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

const (
	PointField_INT8    uint8 = 1
	PointField_UINT8   uint8 = 2
	PointField_INT16   uint8 = 3
	PointField_UINT16  uint8 = 4
	PointField_INT32   uint8 = 5
	PointField_UINT32  uint8 = 6
	PointField_FLOAT32 uint8 = 7
	PointField_FLOAT64 uint8 = 8
)

type PointField struct {
	msg.Package     `ros:"sensor_msgs"`
	msg.Definitions `ros:"uint8 INT8=1,uint8 UINT8=2,uint8 INT16=3,uint8 UINT16=4,uint8 INT32=5,uint8 UINT32=6,uint8 FLOAT32=7,uint8 FLOAT64=8"`
	Name            string
	Offset          uint32
	Datatype        uint8
	Count           uint32
}
