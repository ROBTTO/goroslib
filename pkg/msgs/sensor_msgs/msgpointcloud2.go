//nolint:golint
package sensor_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type PointCloud2 struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	Height      uint32
	Width       uint32
	Fields      []PointField
	IsBigendian bool
	PointStep   uint32
	RowStep     uint32
	Data        []uint8
	IsDense     bool
}
