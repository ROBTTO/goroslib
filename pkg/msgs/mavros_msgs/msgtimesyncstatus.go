//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
	"github.com/ROBTTO/goroslib/pkg/msgs/std_msgs"
)

type TimesyncStatus struct {
	msg.Package       `ros:"mavros_msgs"`
	Header            std_msgs.Header
	RemoteTimestampNs uint64
	ObservedOffsetNs  int64
	EstimatedOffsetNs int64
	RoundTripTimeMs   float32
}
