//nolint:golint
package ackermann_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type AckermannDriveStamped struct {
	msg.Package `ros:"ackermann_msgs"`
	Header      std_msgs.Header
	Drive       AckermannDrive
}
