package item

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
)

type Elisir struct {
	quantity int
}

func NewElisir(quantity int) *Elisir {
	return &Elisir{quantity: quantity}
}

func (e *Elisir) Name() string {
	return "Elisir"
}

func (e *Elisir) Quantity() int {
	return e.quantity
}

func (e *Elisir) Category() pokemon.ItemCategory {
	return pokemon.REMEDY
}

func (e *Elisir) Effect() pokemon.ItemEffect {
	return pokemon.PP_UP
}

func (e *Elisir) ApplyToPokemon(pokemon *pokemon.Pokemon) {
	for _, spell := range pokemon.Spells() {
		spell.SetPP(10)
	}

	e.quantity = e.quantity - 1
}

func (e *Elisir) IncreaseQuantity() {
	e.quantity = e.quantity + 1
}
