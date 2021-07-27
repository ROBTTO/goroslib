//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type FileWriteReq struct {
	FilePath string
	Offset   uint64
	Data     []uint8
}

type FileWriteRes struct {
	Success bool
	RErrno  int32
}

type FileWrite struct {
	msg.Package `ros:"mavros_msgs"`
	FileWriteReq
	FileWriteRes
}
