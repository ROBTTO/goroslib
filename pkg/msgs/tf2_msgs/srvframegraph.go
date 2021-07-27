//nolint:golint
package tf2_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type FrameGraphReq struct{}

type FrameGraphRes struct {
	FrameYaml string
}

type FrameGraph struct {
	msg.Package `ros:"tf2_msgs"`
	FrameGraphReq
	FrameGraphRes
}
