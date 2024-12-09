package state

import (
	"fmt"
	"time"

	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
)

type UsingASpell struct {
	*ChallengingAPokemon
}

func NewUsingASpell(
	state *ChallengingAPokemon,
) *UsingASpell {
	state.logger.Info("UsingASpell")

	return &UsingASpell{
		state,
	}
}

func (u *UsingASpell) UseSpell(spell *pokemon.Spell) {
	spell, spellPosition := u.pokemon.SpellBy(spell.Name())

	u.logger.Info(fmt.Sprintf("spell %s has %d pp", spell.Name(), spell.PP()))

	if !spell.HasEnoughPP() {
		panic(fmt.Sprintf("spell %s has not enough pp", spell.Name()))
	}

	u.controller.Confirm()
	u.controller.Nothing(1 * time.Second)

	if u.lastUsedSpell != nil {
		_, lastUsedSpellPosition := u.pokemon.SpellBy(u.lastUsedSpell.Name())

		for p := lastUsedSpellPosition; p < len(u.pokemon.Spells())+1; p++ {
			u.controller.PressDown()
		}
	}

	for p := 0; p < spellPosition-1; p++ {
		u.controller.PressDown()
	}

	u.controller.Confirm()
	spell.DecreasePP()

	u.controller.Nothing(20 * time.Second)
}
