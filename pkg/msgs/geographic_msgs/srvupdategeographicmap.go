//nolint:golint
package geographic_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

type UpdateGeographicMapReq struct {
	Updates GeographicMapChanges
}

type UpdateGeographicMapRes struct {
	Success bool
	Status  string
}

type UpdateGeographicMap struct {
	msg.Package `ros:"geographic_msgs"`
	UpdateGeographicMapReq
	UpdateGeographicMapRes
}
