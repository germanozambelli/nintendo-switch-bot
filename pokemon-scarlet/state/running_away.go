package state

import (
	"time"
)

type RunningAway struct {
	*ChallengingAPokemon
}

func NewRunningAway(
	state *ChallengingAPokemon,
) *RunningAway {
	state.logger.Info("RunningAway")

	return &RunningAway{
		state,
	}
}

func (r *RunningAway) RunAway() {
	r.controller.PressUp()
	r.controller.Confirm()
	r.controller.Nothing(6 * time.Second)
}
