package player

import (
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/logger"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/bag"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
	"time"
)

type UsingSpellApplicableItem struct {
	player   *Player
	pokemon  *pokemon.Pokemon
	oldState State
	logger   logger.Logger
}

func NewUsingSpellApplicableItem(
	player *Player,
	pokemon *pokemon.Pokemon,
	oldState State,
	logger logger.Logger,
) *UsingSpellApplicableItem {
	return &UsingSpellApplicableItem{
		player:   player,
		pokemon:  pokemon,
		oldState: oldState,
		logger:   logger,
	}
}

func (u *UsingSpellApplicableItem) Name() StateName {
	return USING_SPELL_APPLICABLE_ITEM
}

func (u *UsingSpellApplicableItem) UseSpellApplicableItem(
	spell *pokemon.Spell,
	item pokemon.SpellApplicable,
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

	u.player.Confirm()
	u.player.Confirm()
	u.player.Nothing(1 * time.Second)

	pokemonPosition := u.player.Team().PositionOf(u.pokemon)

	for p := 0; p < pokemonPosition-1; p++ {
		u.player.PressDown()
	}

	u.player.Nothing(1 * time.Second)
	u.player.Confirm()
	u.player.Nothing(1 * time.Second)

	spell, spellPosition := u.pokemon.SpellBy(spell.Name())

	for p := 0; p < spellPosition-1; p++ {
		u.player.PressDown()
	}

	u.player.Confirm()

	u.logger.Info(fmt.Sprintf("applying %s to %s of %s", item.Name(), spell.Name(), u.pokemon.Name()))
	item.ApplyToSpell(spell)

	u.player.Nothing(17 * time.Second)

	u.player.SetState(u.oldState)

	if u.oldState.Name() == RESTORING_SPELL_PP {
		u.oldState.(*RestoringSpellPP).RestorationDone()
		return
	}
}

func (u *UsingSpellApplicableItem) selectRightCategory(category pokemon.ItemCategory) {
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
