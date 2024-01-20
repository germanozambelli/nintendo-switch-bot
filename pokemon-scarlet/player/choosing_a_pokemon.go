package player

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/logger"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
	"time"
)

type ChoosingAPokemon struct {
	player *Player
	logger logger.Logger
}

func NewChoosingAPokemon(
	player *Player,
	logger logger.Logger,
) *ChoosingAPokemon {
	return &ChoosingAPokemon{
		player: player,
		logger: logger,
	}
}

func (c *ChoosingAPokemon) Name() StateName {
	return CHOOSING_A_POKEMON
}

func (c *ChoosingAPokemon) ChooseAPokemon(pokemon *pokemon.Pokemon) {
	position := c.player.Team().PositionOf(pokemon)

	c.player.PressDown()
	c.player.Confirm()

	for p := 1; p < position; p++ {
		c.player.PressDown()
	}

	c.player.Confirm()
	c.player.Confirm()
	c.player.Nothing(20 * time.Second)

	c.player.SetState(NewChallengingAPokemon(c.player, pokemon, c.logger))
}
