//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

const (
	FileEntry_TYPE_FILE      uint8 = 0
	FileEntry_TYPE_DIRECTORY uint8 = 1
)

type FileEntry struct {
	msg.Package     `ros:"mavros_msgs"`
	msg.Definitions `ros:"uint8 TYPE_FILE=0,uint8 TYPE_DIRECTORY=1"`
	Name            string
	Type            uint8
	Size            uint64
}
