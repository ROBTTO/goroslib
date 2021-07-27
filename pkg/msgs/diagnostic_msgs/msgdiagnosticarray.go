//nolint:golint
package diagnostic_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type DiagnosticArray struct {
	msg.Package `ros:"diagnostic_msgs"`
	Header      std_msgs.Header
	Status      []DiagnosticStatus
}
