package player

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/logger"
	"time"
)

type RunningAway struct {
	player *Player
	logger logger.Logger
}

func NewRunningAway(
	player *Player,
	logger logger.Logger,
) *RunningAway {
	return &RunningAway{
		player: player,
		logger: logger,
	}
}

func (r *RunningAway) Name() StateName {
	return RUNNING_AWAY
}

func (r *RunningAway) RunAway() {
	r.player.PressUp()
	r.player.Confirm()
	r.player.Nothing(6 * time.Second)

	r.player.SetState(NewInFreeWorld(r.player, r.logger))
}
