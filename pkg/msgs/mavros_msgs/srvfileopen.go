//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

const (
	FileOpenReq_MODE_READ   uint8 = 0
	FileOpenReq_MODE_WRITE  uint8 = 1
	FileOpenReq_MODE_CREATE uint8 = 2
)

type FileOpenReq struct {
	msg.Definitions `ros:"uint8 MODE_READ=0,uint8 MODE_WRITE=1,uint8 MODE_CREATE=2"`
	FilePath        string
	Mode            uint8
}

type FileOpenRes struct {
	Size    uint32
	Success bool
	RErrno  int32
}

type FileOpen struct {
	msg.Package `ros:"mavros_msgs"`
	FileOpenReq
	FileOpenRes
}
