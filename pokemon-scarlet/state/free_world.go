package state

import (
	"time"
)

type InFreeWorld struct {
	*Game
}

func NewInFreeWorld(
	state *Game,
) *InFreeWorld {
	state.logger.Info("In free world")

	return &InFreeWorld{
		state,
	}
}

func (i *InFreeWorld) StartABattle(timeDistanceToEnemy time.Duration) *ChallengingAPokemon {
	i.controller.PressL()
	i.controller.MoveUp(timeDistanceToEnemy)
	i.controller.Nothing(6 * time.Second)

	return NewChallengingAPokemon(i)
}
