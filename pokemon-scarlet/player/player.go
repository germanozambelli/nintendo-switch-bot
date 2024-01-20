package player

import (
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/bag"
	"time"

	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/logger"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/player"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
)

type Player struct {
	*player.Player
	team   *pokemon.Team
	bag    *bag.Bag
	state  State
	logger logger.Logger
}

func NewPlayer(
	logger logger.Logger,
	basePlayer *player.Player,
	team *pokemon.Team,
	bag *bag.Bag,
) *Player {
	p := &Player{
		basePlayer,
		team,
		bag,
		nil,
		logger,
	}

	p.SetState(NewInFreeWorld(p, p.logger))

	return p
}

func (p *Player) Team() *pokemon.Team {
	return p.team
}

func (p *Player) State() State {
	return p.state
}

func (p *Player) Bag() *bag.Bag {
	return p.bag
}

func (p *Player) SetState(state State) {
	if p.state != nil {
		p.logger.Info(fmt.Sprintf("changing state from %s to %s", p.state.Name(), state.Name()))
	}

	if p.state == nil {
		p.logger.Info(fmt.Sprintf("changing state to %s", state.Name()))
	}

	p.state = state
}

func (p *Player) StartABattle(timeDistanceToEnemy time.Duration) {
	if p.state.Name() != IN_FREE_WORLD {
		panic("can't start a battle if not in free world")
	}

	p.logger.Info("starting a battle")

	p.state.(*InFreeWorld).StartABattle(timeDistanceToEnemy)
}

func (p *Player) ChooseAPokemon(pokemon *pokemon.Pokemon) {
	if p.state.Name() != CHALLENGING_A_POKEMON {
		panic("can't choose a pokemon if not in a battle")
	}

	p.logger.Info(fmt.Sprintf("choosing a pokemon: %s", pokemon.Name()))

	p.state.(*ChallengingAPokemon).ChooseAPokemon(pokemon)
}

func (p *Player) UseSpell(spellName string) {
	if p.state.Name() != CHALLENGING_A_POKEMON {
		panic("can't use a spell if not in a battle")
	}

	p.logger.Info(fmt.Sprintf("using a spell: %s", spellName))

	p.state.(*ChallengingAPokemon).UseSpell(spellName)
}

func (p *Player) RunAway() {
	if p.state.Name() != CHALLENGING_A_POKEMON {
		panic("can't run away if not in a battle")
	}

	p.logger.Info("running away")

	p.state.(*ChallengingAPokemon).RunAway()
}
