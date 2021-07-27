//nolint:golint
package sensor_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type LaserEcho struct {
	msg.Package `ros:"sensor_msgs"`
	Echoes      []float32
}
