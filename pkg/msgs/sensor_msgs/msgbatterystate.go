//nolint:golint
package sensor_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

const (
	BatteryState_POWER_SUPPLY_STATUS_UNKNOWN               uint8 = 0
	BatteryState_POWER_SUPPLY_STATUS_CHARGING              uint8 = 1
	BatteryState_POWER_SUPPLY_STATUS_DISCHARGING           uint8 = 2
	BatteryState_POWER_SUPPLY_STATUS_NOT_CHARGING          uint8 = 3
	BatteryState_POWER_SUPPLY_STATUS_FULL                  uint8 = 4
	BatteryState_POWER_SUPPLY_HEALTH_UNKNOWN               uint8 = 0
	BatteryState_POWER_SUPPLY_HEALTH_GOOD                  uint8 = 1
	BatteryState_POWER_SUPPLY_HEALTH_OVERHEAT              uint8 = 2
	BatteryState_POWER_SUPPLY_HEALTH_DEAD                  uint8 = 3
	BatteryState_POWER_SUPPLY_HEALTH_OVERVOLTAGE           uint8 = 4
	BatteryState_POWER_SUPPLY_HEALTH_UNSPEC_FAILURE        uint8 = 5
	BatteryState_POWER_SUPPLY_HEALTH_COLD                  uint8 = 6
	BatteryState_POWER_SUPPLY_HEALTH_WATCHDOG_TIMER_EXPIRE uint8 = 7
	BatteryState_POWER_SUPPLY_HEALTH_SAFETY_TIMER_EXPIRE   uint8 = 8
	BatteryState_POWER_SUPPLY_TECHNOLOGY_UNKNOWN           uint8 = 0
	BatteryState_POWER_SUPPLY_TECHNOLOGY_NIMH              uint8 = 1
	BatteryState_POWER_SUPPLY_TECHNOLOGY_LION              uint8 = 2
	BatteryState_POWER_SUPPLY_TECHNOLOGY_LIPO              uint8 = 3
	BatteryState_POWER_SUPPLY_TECHNOLOGY_LIFE              uint8 = 4
	BatteryState_POWER_SUPPLY_TECHNOLOGY_NICD              uint8 = 5
	BatteryState_POWER_SUPPLY_TECHNOLOGY_LIMN              uint8 = 6
)

type BatteryState struct {
	msg.Package           `ros:"sensor_msgs"`
	msg.Definitions       `ros:"uint8 POWER_SUPPLY_STATUS_UNKNOWN=0,uint8 POWER_SUPPLY_STATUS_CHARGING=1,uint8 POWER_SUPPLY_STATUS_DISCHARGING=2,uint8 POWER_SUPPLY_STATUS_NOT_CHARGING=3,uint8 POWER_SUPPLY_STATUS_FULL=4,uint8 POWER_SUPPLY_HEALTH_UNKNOWN=0,uint8 POWER_SUPPLY_HEALTH_GOOD=1,uint8 POWER_SUPPLY_HEALTH_OVERHEAT=2,uint8 POWER_SUPPLY_HEALTH_DEAD=3,uint8 POWER_SUPPLY_HEALTH_OVERVOLTAGE=4,uint8 POWER_SUPPLY_HEALTH_UNSPEC_FAILURE=5,uint8 POWER_SUPPLY_HEALTH_COLD=6,uint8 POWER_SUPPLY_HEALTH_WATCHDOG_TIMER_EXPIRE=7,uint8 POWER_SUPPLY_HEALTH_SAFETY_TIMER_EXPIRE=8,uint8 POWER_SUPPLY_TECHNOLOGY_UNKNOWN=0,uint8 POWER_SUPPLY_TECHNOLOGY_NIMH=1,uint8 POWER_SUPPLY_TECHNOLOGY_LION=2,uint8 POWER_SUPPLY_TECHNOLOGY_LIPO=3,uint8 POWER_SUPPLY_TECHNOLOGY_LIFE=4,uint8 POWER_SUPPLY_TECHNOLOGY_NICD=5,uint8 POWER_SUPPLY_TECHNOLOGY_LIMN=6"`
	Header                std_msgs.Header
	Voltage               float32
	Current               float32
	Charge                float32
	Capacity              float32
	DesignCapacity        float32
	Percentage            float32
	PowerSupplyStatus     uint8
	PowerSupplyHealth     uint8
	PowerSupplyTechnology uint8
	Present               bool
	CellVoltage           []float32
	Location              string
	SerialNumber          string
}
