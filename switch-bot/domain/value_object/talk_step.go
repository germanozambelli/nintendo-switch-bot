package value_object

import (
	"time"
)

type TalkStep struct {
	Confirm         bool
	Duration        time.Duration
	DelayOfNextStep time.Duration
}

func NewTalkStep(confirm bool, duration time.Duration, delayOfNextStep time.Duration) TalkStep {
	return TalkStep{Confirm: confirm, Duration: duration, DelayOfNextStep: delayOfNextStep}
}
