//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type FileCloseReq struct {
	FilePath string
}

type FileCloseRes struct {
	Success bool
	RErrno  int32
}

type FileClose struct {
	msg.Package `ros:"mavros_msgs"`
	FileCloseReq
	FileCloseRes
}
