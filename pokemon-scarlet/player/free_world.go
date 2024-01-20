package player

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/logger"
	"time"
)

type InFreeWorld struct {
	player *Player
	logger logger.Logger
}

func NewInFreeWorld(
	player *Player,
	logger logger.Logger,
) *InFreeWorld {
	return &InFreeWorld{
		player: player,
		logger: logger,
	}
}

func (i *InFreeWorld) Name() StateName {
	return IN_FREE_WORLD
}

func (i *InFreeWorld) StartABattle(timeDistanceToEnemy time.Duration) {
	newState := NewChallengingAPokemon(i.player, i.player.Team().At(1), i.logger)
	i.player.SetState(newState)
	newState.StartABattle(timeDistanceToEnemy)
}
