package player

import (
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/logger"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
)

type RestoringSpellPP struct {
	player   *Player
	pokemon  *pokemon.Pokemon
	oldState State
	logger   logger.Logger
}

func NewRestoringSpellPP(
	player *Player,
	pokemon *pokemon.Pokemon,
	oldState State,
	logger logger.Logger,
) *RestoringSpellPP {
	return &RestoringSpellPP{
		player:   player,
		pokemon:  pokemon,
		oldState: oldState,
		logger:   logger,
	}
}

func (r *RestoringSpellPP) Name() StateName {
	return RESTORING_SPELL_PP
}

func (r *RestoringSpellPP) RestoreSpellPP(spell *pokemon.Spell) {
	suitableItem, position := r.player.Bag().SearchItem(pokemon.PP_UP, 2)

	if suitableItem == nil {
		panic(fmt.Sprintf("cannot find suitable item to restore spell %s pp", spell.Name()))
	}

	switch suitableItem.(type) {
	case pokemon.SpellApplicable:
		state := NewUsingSpellApplicableItem(r.player, r.pokemon, r, r.logger)
		r.player.SetState(state)
		state.UseSpellApplicableItem(spell, suitableItem.(pokemon.SpellApplicable), position)
		return
	case pokemon.PokemonApplicable:
		state := NewUsingPokemonApplicableItem(r.player, r.pokemon, r, r.logger)
		r.player.SetState(state)
		state.UsePokemonApplicableItem(suitableItem.(pokemon.PokemonApplicable), position)
		return
	}

}

func (r *RestoringSpellPP) RestorationDone() {
	r.player.SetState(r.oldState)
}
