package action

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller/button"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/value_object"
	"time"
)

type Talk struct {
	steps      []value_object.TalkStep
	controller controller.Controller
}

func NewTalk(steps []value_object.TalkStep, controller controller.Controller) Talk {
	return Talk{steps: steps, controller: controller}
}

func (talk Talk) Play() {
	for _, step := range talk.steps {
		NewNothing(step.Duration).Play()

		if step.Confirm {
			NewPressButton(button.A, 75*time.Millisecond, talk.controller).Play()
		} else {
			NewPressButton(button.B, 75*time.Millisecond, talk.controller).Play()
		}

		NewNothing(step.DelayOfNextStep).Play()
	}
}
