package item

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
)

type ElisirMax struct {
	quantity int
}

func NewElisirMax(quantity int) *ElisirMax {
	return &ElisirMax{quantity: quantity}
}

func (e *ElisirMax) Name() string {
	return "Elisir max"
}

func (e *ElisirMax) Quantity() int {
	return e.quantity
}

func (e *ElisirMax) Category() pokemon.ItemCategory {
	return pokemon.REMEDY
}

func (e *ElisirMax) Effect() pokemon.ItemEffect {
	return pokemon.PP_UP
}

func (e *ElisirMax) ApplyToPokemon(pokemon *pokemon.Pokemon) {
	for _, spell := range pokemon.Spells() {
		spell.SetPP(spell.MaxPP())
	}

	e.quantity = e.quantity - 1
}

func (e *ElisirMax) IncreaseQuantity() {
	e.quantity = e.quantity + 1
}
