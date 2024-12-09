package state

import (
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
)

type ChallengingAPokemon struct {
	*InFreeWorld
	pokemon       *pokemon.Pokemon
	lastUsedSpell *pokemon.Spell
}

func NewChallengingAPokemon(
	state *InFreeWorld,
) *ChallengingAPokemon {
	state.logger.Info("ChallengingAPokemon")

	return &ChallengingAPokemon{
		InFreeWorld: state,
	}
}

func (c *ChallengingAPokemon) ChooseAPokemon(p *pokemon.Pokemon) *ChallengingAPokemon {
	c.pokemon = c.player.Team().At(1)

	isInTheBattle := func(p *pokemon.Pokemon) bool {
		return c.pokemon == p
	}

	if !isInTheBattle(p) {
		NewChoosingAPokemon(c).
			ChooseAPokemon(p)

		c.pokemon = p
	}

	if c.pokemon.Item() != nil && c.pokemon.Item().Quantity() >= 1 {
		NewUsingHoldableItem(c).UseHoldableItem()
	}

	return c
}

func (c *ChallengingAPokemon) UseSpell(spellName string) *ChallengingAPokemon {
	spell, _ := c.pokemon.SpellBy(spellName)

	if spell == nil {
		panic(fmt.Sprintf("spell %s not found", spellName))
	}

	if !spell.HasEnoughPP() {
		NewRestoringSpellPP(c).RestoreSpellPP(spell)
	}

	NewUsingASpell(c).UseSpell(spell)

	c.lastUsedSpell = spell

	return c
}

func (c *ChallengingAPokemon) RunAway() *RunningAway {
	runningAwayState := NewRunningAway(c)
	runningAwayState.RunAway()
	return runningAwayState
}
