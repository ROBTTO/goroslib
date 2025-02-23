//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type ParamValue struct {
	msg.Package `ros:"mavros_msgs"`
	Integer     int64
	Real        float64
}
