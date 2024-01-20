package player

import (
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/logger"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
	"time"
)

type ChallengingAPokemon struct {
	pokemon *pokemon.Pokemon
	player  *Player
	logger  logger.Logger
}

func NewChallengingAPokemon(
	player *Player,
	pokemon *pokemon.Pokemon,
	logger logger.Logger,
) *ChallengingAPokemon {
	return &ChallengingAPokemon{
		player:  player,
		pokemon: pokemon,
		logger:  logger,
	}
}

func (i *ChallengingAPokemon) Name() StateName {
	return CHALLENGING_A_POKEMON
}

func (i *ChallengingAPokemon) isInTheBattle(pokemon *pokemon.Pokemon) bool {
	return i.pokemon == pokemon
}

func (i *ChallengingAPokemon) StartABattle(timeDistanceToEnemy time.Duration) {
	i.player.PressL()
	i.player.MoveUp(timeDistanceToEnemy)
	i.player.Nothing(6 * time.Second)
}

func (i *ChallengingAPokemon) ChooseAPokemon(pokemon *pokemon.Pokemon) {
	if !i.isInTheBattle(pokemon) {
		choosingPokemonState := NewChoosingAPokemon(i.player, i.logger)
		i.player.SetState(choosingPokemonState)
		choosingPokemonState.ChooseAPokemon(pokemon)
	}

	if i.pokemon.Item() != nil && i.pokemon.Item().Quantity() >= 1 {
		usingHoldableItemState := NewUsingHoldableItem(i.player, i.pokemon, i.logger)
		i.player.SetState(usingHoldableItemState)
		usingHoldableItemState.UseHoldableItem()
	}
}

func (i *ChallengingAPokemon) UseSpell(spellName string) {
	spell, position := i.pokemon.SpellBy(spellName)

	if spell == nil {
		panic(fmt.Sprintf("spell %s not found", spellName))
	}

	i.logger.Info(fmt.Sprintf("spell %s has %d pp", spellName, spell.PP()))

	if !spell.HasEnoughPP() {
		restoreSpellState := NewRestoringSpellPP(i.player, i.pokemon, i, i.logger)
		i.player.SetState(restoreSpellState)
		restoreSpellState.RestoreSpellPP(spell)
	}

	usingASpellState := NewUsingASpell(i.player, i.pokemon, i.logger)
	i.player.SetState(usingASpellState)
	usingASpellState.UseSpell(spell, position)
}

func (i *ChallengingAPokemon) RunAway() {
	runningAwayState := NewRunningAway(i.player, i.logger)
	i.player.SetState(runningAwayState)
	runningAwayState.RunAway()
}
