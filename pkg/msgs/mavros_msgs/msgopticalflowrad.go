//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type OpticalFlowRad struct {
	msg.Package         `ros:"mavros_msgs"`
	Header              std_msgs.Header
	IntegrationTimeUs   uint32
	IntegratedX         float32
	IntegratedY         float32
	IntegratedXgyro     float32
	IntegratedYgyro     float32
	IntegratedZgyro     float32
	Temperature         int16
	Quality             uint8
	TimeDeltaDistanceUs uint32
	Distance            float32
}
