//nolint:golint
package mavros_msgs

import (
	"github.com/ROBTTO/goroslib/pkg/msg"
)

const (
	PlayTuneV2_QBASIC1_1  uint8 = 1
	PlayTuneV2_MML_MODERN uint8 = 2
)

type PlayTuneV2 struct {
	msg.Package     `ros:"mavros_msgs"`
	msg.Definitions `ros:"uint8 QBASIC1_1=1,uint8 MML_MODERN=2"`
	Format          uint8
	Tune            string
}
