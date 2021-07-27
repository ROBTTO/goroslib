//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type FileMakeDirReq struct {
	DirPath string
}

type FileMakeDirRes struct {
	Success bool
	RErrno  int32
}

type FileMakeDir struct {
	msg.Package `ros:"mavros_msgs"`
	FileMakeDirReq
	FileMakeDirRes
}
