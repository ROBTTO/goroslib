//nolint:golint
package uuid_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type UniqueID struct {
	msg.Package `ros:"uuid_msgs"`
	Uuid        [16]uint8
}
