//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type FileChecksumReq struct {
	FilePath string
}

type FileChecksumRes struct {
	Crc32   uint32
	Success bool
	RErrno  int32
}

type FileChecksum struct {
	msg.Package `ros:"mavros_msgs"`
	FileChecksumReq
	FileChecksumRes
}
