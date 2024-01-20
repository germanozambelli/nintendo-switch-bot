package player

import (
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/logger"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/bag"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
	"time"
)

type UsingPokemonApplicableItem struct {
	player   *Player
	pokemon  *pokemon.Pokemon
	oldState State
	logger   logger.Logger
}

func NewUsingPokemonApplicableItem(
	player *Player,
	pokemon *pokemon.Pokemon,
	oldState State,
	logger logger.Logger,
) *UsingPokemonApplicableItem {
	return &UsingPokemonApplicableItem{
		player:   player,
		pokemon:  pokemon,
		oldState: oldState,
		logger:   logger,
	}
}

func (u *UsingPokemonApplicableItem) Name() StateName {
	return USING_POKEMON_APPLICABLE_ITEM
}

func (u *UsingPokemonApplicableItem) UsePokemonApplicableItem(
	item pokemon.PokemonApplicable,
	position bag.Position,
) {
	u.player.PressDown()
	u.player.PressDown()
	u.player.Confirm()
	u.player.Nothing(1 * time.Second)

	u.player.PressMinus()
	u.player.PressUp()
	u.player.PressUp()
	u.player.Confirm()
	u.player.Nothing(1 * time.Second)

	u.selectRightCategory(item.Category())

	for p := 0; p < int(position)-1; p++ {
		u.player.PressDown()
	}

	u.player.Nothing(1 * time.Second)
	u.player.Confirm()
	u.player.Nothing(75 * time.Millisecond)
	u.player.Confirm()
	u.player.Nothing(1 * time.Second)
	u.player.Confirm()
	u.player.Nothing(1 * time.Second)

	u.logger.Info(fmt.Sprintf("applying %s to %s", item.Name(), u.pokemon.Name()))
	item.ApplyToPokemon(u.pokemon)

	u.player.Nothing(17 * time.Second)

	u.player.SetState(u.oldState)

	if u.oldState.Name() == RESTORING_SPELL_PP {
		u.oldState.(*RestoringSpellPP).RestorationDone()
		return
	}
}

func (u *UsingPokemonApplicableItem) selectRightCategory(category pokemon.ItemCategory) {
	switch category {
	case pokemon.REMEDY:
		return
	case pokemon.BERRY:
		u.player.PressRight()
		u.player.PressRight()
		u.player.PressRight()
		return
	case pokemon.BALL:
		u.player.PressRight()
		return
	case pokemon.BATTLE:
		u.player.PressRight()
		u.player.PressRight()
		return
	}
}
