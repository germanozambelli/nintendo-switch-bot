package player

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/bag"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
)

type Player struct {
	team *pokemon.Team
	bag  *bag.Bag
}

func NewPlayer(
	team *pokemon.Team,
	bag *bag.Bag,
) *Player {
	p := &Player{
		team,
		bag,
	}

	return p
}

func (p *Player) Team() *pokemon.Team {
	return p.team
}

func (p *Player) Bag() *bag.Bag {
	return p.bag
}
