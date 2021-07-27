//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type FileListReq struct {
	DirPath string
}

type FileListRes struct {
	List    []FileEntry
	Success bool
	RErrno  int32
}

type FileList struct {
	msg.Package `ros:"mavros_msgs"`
	FileListReq
	FileListRes
}
