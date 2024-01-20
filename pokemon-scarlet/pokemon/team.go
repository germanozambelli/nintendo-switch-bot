package pokemon

import "fmt"

type Team struct {
	pokemons []*Pokemon
}

func NewTeam(pokemons ...*Pokemon) (*Team, error) {
	if len(pokemons) > 6 {
		return nil, fmt.Errorf("a team can't have more than 6 pokemons")
	}

	return &Team{pokemons: pokemons}, nil
}

func (t *Team) PositionOf(pokemon *Pokemon) int {
	for i, p := range t.pokemons {
		if p == pokemon {
			return i + 1
		}
	}

	panic("pokemon not found in the team")
}

func (t *Team) At(position int) *Pokemon {
	if position > len(t.pokemons) {
		panic("position out of range")
	}

	if position < 1 {
		panic("position out of range")
	}

	return t.pokemons[position-1]
}
