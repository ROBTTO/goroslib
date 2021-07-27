//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type FileTruncateReq struct {
	FilePath string
	Length   uint64
}

type FileTruncateRes struct {
	Success bool
	RErrno  int32
}

type FileTruncate struct {
	msg.Package `ros:"mavros_msgs"`
	FileTruncateReq
	FileTruncateRes
}
