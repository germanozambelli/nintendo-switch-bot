package item

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
)

type Etere struct {
	quantity int
}

func NewEtere(
	quantity int,
) *Etere {
	return &Etere{
		quantity: quantity,
	}
}

func (e *Etere) Name() string {
	return "Etere"
}

func (e *Etere) Quantity() int {
	return e.quantity
}

func (e *Etere) Category() pokemon.ItemCategory {
	return pokemon.REMEDY
}

func (e *Etere) Effect() pokemon.ItemEffect {
	return pokemon.PP_UP
}

func (e *Etere) ApplyToSpell(spell *pokemon.Spell) {
	spell.SetPP(10)

	e.quantity = e.quantity - 1
}

func (e *Etere) IncreaseQuantity() {
	e.quantity = e.quantity + 1
}
