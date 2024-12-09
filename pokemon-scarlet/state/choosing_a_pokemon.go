package state

import (
	"time"

	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
)

type ChoosingAPokemon struct {
	*ChallengingAPokemon
}

func NewChoosingAPokemon(
	state *ChallengingAPokemon,
) *ChoosingAPokemon {
	state.logger.Info("ChoosingAPokemon")

	return &ChoosingAPokemon{
		state,
	}
}

func (c *ChoosingAPokemon) ChooseAPokemon(pokemon *pokemon.Pokemon) {
	position := c.player.Team().PositionOf(pokemon)

	c.controller.PressDown()
	c.controller.Confirm()

	c.controller.Nothing(2 * time.Second)

	for p := 1; p < position; p++ {
		c.controller.PressDown()
	}

	c.controller.Confirm()
	c.controller.Confirm()
	c.controller.Nothing(20 * time.Second)
}
