package item

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
)

type EtereMax struct {
	quantity int
}

func NewEtereMax(quantity int) *EtereMax {
	return &EtereMax{quantity: quantity}
}

func (e *EtereMax) Name() string {
	return "Etere Max"
}

func (e *EtereMax) Quantity() int {
	return e.quantity
}

func (e *EtereMax) Category() pokemon.ItemCategory {
	return pokemon.REMEDY
}

func (e *EtereMax) Effect() pokemon.ItemEffect {
	return pokemon.PP_UP
}

func (e *EtereMax) ApplyToSpell(spell *pokemon.Spell) {
	spell.SetPP(spell.MaxPP())

	e.quantity = e.quantity - 1
}

func (e *EtereMax) IncreaseQuantity() {
	e.quantity = e.quantity + 1
}
