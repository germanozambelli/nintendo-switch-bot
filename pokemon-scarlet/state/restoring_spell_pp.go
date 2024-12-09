package state

import (
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/bag/item"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
)

type RestoringSpellPP struct {
	*ChallengingAPokemon
}

func NewRestoringSpellPP(
	state *ChallengingAPokemon,
) *RestoringSpellPP {
	state.logger.Info("RestoringSpellPP")

	return &RestoringSpellPP{
		state,
	}
}

func (r *RestoringSpellPP) RestoreSpellPP(spell *pokemon.Spell) {
	suitableItem, position := r.player.Bag().SearchItem(item.PP_UP, 2)

	if suitableItem == nil {
		panic(fmt.Sprintf("cannot find suitable item to restore spell %s pp", spell.Name()))
	}

	switch suitableItem.(type) {
	case item.SpellApplicable:
		state := NewUsingSpellApplicableItem(r)
		state.UseSpellApplicableItem(spell, suitableItem.(item.SpellApplicable), position)
		return
	case item.PokemonApplicable:
		state := NewUsingPokemonApplicableItem(r)
		state.UsePokemonApplicableItem(suitableItem.(item.PokemonApplicable), position)
		return
	}
}
