package state

import (
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/bag/item"
	"time"

	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/bag"
)

type UsingPokemonApplicableItem struct {
	*RestoringSpellPP
}

func NewUsingPokemonApplicableItem(
	state *RestoringSpellPP,
) *UsingPokemonApplicableItem {
	state.logger.Info("UsingPokemonApplicableItem")

	return &UsingPokemonApplicableItem{
		state,
	}
}

func (u *UsingPokemonApplicableItem) UsePokemonApplicableItem(
	item item.PokemonApplicable,
	position bag.Position,
) {
	u.controller.PressDown()
	u.controller.PressDown()
	u.controller.Confirm()
	u.controller.Nothing(1 * time.Second)

	u.controller.PressMinus()
	u.controller.PressUp()
	u.controller.PressUp()
	u.controller.Confirm()
	u.controller.Nothing(1 * time.Second)

	u.selectRightCategory(item.Category())

	for p := 0; p < int(position)-1; p++ {
		u.controller.PressDown()
	}

	u.controller.Nothing(1 * time.Second)
	u.controller.Confirm()
	u.controller.Nothing(75 * time.Millisecond)
	u.controller.Confirm()
	u.controller.Nothing(1 * time.Second)
	u.controller.Confirm()
	u.controller.Nothing(1 * time.Second)

	u.logger.Info(fmt.Sprintf("applying %s to %s", item.Name(), u.pokemon.Name()))
	item.ApplyToPokemon(u.pokemon)

	u.controller.Nothing(17 * time.Second)

}

func (u *UsingPokemonApplicableItem) selectRightCategory(category item.ItemCategory) {
	switch category {
	case item.REMEDY:
		return
	case item.BERRY:
		u.controller.PressRight()
		u.controller.PressRight()
		u.controller.PressRight()
		return
	case item.BALL:
		u.controller.PressRight()
		return
	case item.BATTLE:
		u.controller.PressRight()
		u.controller.PressRight()
		return
	}
}
