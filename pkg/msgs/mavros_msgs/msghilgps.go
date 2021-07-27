//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/geographic_msgs"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type HilGPS struct {
	msg.Package       `ros:"mavros_msgs"`
	Header            std_msgs.Header
	FixType           uint8
	Geo               geographic_msgs.GeoPoint
	Eph               uint16
	Epv               uint16
	Vel               uint16
	Vn                int16
	Ve                int16
	Vd                int16
	Cog               uint16
	SatellitesVisible uint8
}
