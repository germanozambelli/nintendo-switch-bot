package player

import (
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/logger"
	"time"

	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
)

type UsingASpell struct {
	player  *Player
	pokemon *pokemon.Pokemon
	logger  logger.Logger
}

func NewUsingASpell(
	player *Player,
	pokemon *pokemon.Pokemon,
	logger logger.Logger,
) *UsingASpell {
	return &UsingASpell{
		player:  player,
		pokemon: pokemon,
		logger:  logger,
	}
}

func (u *UsingASpell) Name() StateName {
	return USING_A_SPELL
}

func (u *UsingASpell) UseSpell(spell *pokemon.Spell, position int) {
	if !spell.HasEnoughPP() {
		panic(fmt.Sprintf("spell %s has not enough pp", spell.Name()))
	}

	u.player.Confirm()
	u.player.Nothing(1 * time.Second)

	for p := 0; p < position-1; p++ {
		u.player.PressDown()
	}

	u.player.Confirm()
	spell.DecreasePP()

	u.player.Nothing(17 * time.Second)

	u.player.SetState(NewChallengingAPokemon(u.player, u.pokemon, u.logger))
}
