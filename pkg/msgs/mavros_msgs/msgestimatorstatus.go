//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type EstimatorStatus struct {
	msg.Package               `ros:"mavros_msgs"`
	Header                    std_msgs.Header
	AttitudeStatusFlag        bool
	VelocityHorizStatusFlag   bool
	VelocityVertStatusFlag    bool
	PosHorizRelStatusFlag     bool
	PosHorizAbsStatusFlag     bool
	PosVertAbsStatusFlag      bool
	PosVertAglStatusFlag      bool
	ConstPosModeStatusFlag    bool
	PredPosHorizRelStatusFlag bool
	PredPosHorizAbsStatusFlag bool
	GpsGlitchStatusFlag       bool
	AccelErrorStatusFlag      bool
}
