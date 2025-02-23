//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

const (
	ExtendedState_VTOL_STATE_UNDEFINED        uint8 = 0
	ExtendedState_VTOL_STATE_TRANSITION_TO_FW uint8 = 1
	ExtendedState_VTOL_STATE_TRANSITION_TO_MC uint8 = 2
	ExtendedState_VTOL_STATE_MC               uint8 = 3
	ExtendedState_VTOL_STATE_FW               uint8 = 4
	ExtendedState_LANDED_STATE_UNDEFINED      uint8 = 0
	ExtendedState_LANDED_STATE_ON_GROUND      uint8 = 1
	ExtendedState_LANDED_STATE_IN_AIR         uint8 = 2
	ExtendedState_LANDED_STATE_TAKEOFF        uint8 = 3
	ExtendedState_LANDED_STATE_LANDING        uint8 = 4
)

type ExtendedState struct {
	msg.Package     `ros:"mavros_msgs"`
	msg.Definitions `ros:"uint8 VTOL_STATE_UNDEFINED=0,uint8 VTOL_STATE_TRANSITION_TO_FW=1,uint8 VTOL_STATE_TRANSITION_TO_MC=2,uint8 VTOL_STATE_MC=3,uint8 VTOL_STATE_FW=4,uint8 LANDED_STATE_UNDEFINED=0,uint8 LANDED_STATE_ON_GROUND=1,uint8 LANDED_STATE_IN_AIR=2,uint8 LANDED_STATE_TAKEOFF=3,uint8 LANDED_STATE_LANDING=4"`
	Header          std_msgs.Header
	VtolState       uint8
	LandedState     uint8
}
