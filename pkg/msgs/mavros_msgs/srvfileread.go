//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type FileReadReq struct {
	FilePath string
	Offset   uint64
	Size     uint64
}

type FileReadRes struct {
	Data    []uint8
	Success bool
	RErrno  int32
}

type FileRead struct {
	msg.Package `ros:"mavros_msgs"`
	FileReadReq
	FileReadRes
}
