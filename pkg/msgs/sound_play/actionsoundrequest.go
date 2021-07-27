//nolint:golint
package sound_play

import (
	"time"

	"github.com/ROBTTO/goroslib/pkg/msg"
)

type SoundRequestActionGoal struct {
	SoundRequest SoundRequest
}

type SoundRequestActionResult struct {
	Playing bool
	Stamp   time.Time
}

type SoundRequestActionFeedback struct {
	Playing bool
	Stamp   time.Time
}

type SoundRequestAction struct {
	msg.Package `ros:"sound_play"`
	SoundRequestActionGoal
	SoundRequestActionResult
	SoundRequestActionFeedback
}
