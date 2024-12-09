package state

import (
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/bag/item"
	"time"

	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/bag"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
)

type UsingSpellApplicableItem struct {
	*RestoringSpellPP
}

func NewUsingSpellApplicableItem(
	state *RestoringSpellPP,
) *UsingSpellApplicableItem {
	state.logger.Info("UsingSpellApplicableItem")

	return &UsingSpellApplicableItem{
		state,
	}
}

func (u *UsingSpellApplicableItem) UseSpellApplicableItem(
	spell *pokemon.Spell,
	item item.SpellApplicable,
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

	u.controller.Confirm()
	u.controller.Confirm()
	u.controller.Nothing(1 * time.Second)

	pokemonPosition := u.player.Team().PositionOf(u.pokemon)

	for p := 0; p < pokemonPosition-1; p++ {
		u.controller.PressDown()
	}

	u.controller.Nothing(1 * time.Second)
	u.controller.Confirm()
	u.controller.Nothing(1 * time.Second)

	spell, spellPosition := u.pokemon.SpellBy(spell.Name())

	for p := 0; p < spellPosition-1; p++ {
		u.controller.PressDown()
	}

	u.controller.Confirm()

	u.logger.Info(fmt.Sprintf("applying %s to %s of %s", item.Name(), spell.Name(), u.pokemon.Name()))

	item.ApplyToSpell(spell)

	u.controller.Nothing(17 * time.Second)
}

func (u *UsingSpellApplicableItem) selectRightCategory(category item.ItemCategory) {
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
