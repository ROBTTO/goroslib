//nolint:golint
package control_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type QueryCalibrationStateReq struct{}

type QueryCalibrationStateRes struct {
	IsCalibrated bool
}

type QueryCalibrationState struct {
	msg.Package `ros:"control_msgs"`
	QueryCalibrationStateReq
	QueryCalibrationStateRes
}
