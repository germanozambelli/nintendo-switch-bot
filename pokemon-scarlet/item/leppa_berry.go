package item

import "github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"

type LeppaBerry struct {
	quantity int
}

func NewLeppaBerry(quantity int) *LeppaBerry {
	return &LeppaBerry{quantity: quantity}
}

func (l *LeppaBerry) Name() string {
	return "Leppa Berry"
}

func (l *LeppaBerry) Quantity() int {
	return l.quantity
}

func (l *LeppaBerry) Category() pokemon.ItemCategory {
	return pokemon.BERRY
}

func (l *LeppaBerry) Effect() pokemon.ItemEffect {
	return pokemon.PP_UP
}

func (l *LeppaBerry) ApplyToSpell(spell *pokemon.Spell) {
	spell.SetPP(10)

	l.quantity = l.quantity - 1
}

func (l *LeppaBerry) ApplyHoldingEffect(p *pokemon.Pokemon) pokemon.HoldableItemApplied {
	var zeroPPSpell *pokemon.Spell
	for _, spell := range p.Spells() {
		if spell.PP() == 0 {
			zeroPPSpell = spell
			break
		}
	}

	if zeroPPSpell == nil {
		return false
	}

	zeroPPSpell.SetPP(10)

	l.quantity = l.quantity - 1

	return true
}

func (l *LeppaBerry) IncreaseQuantity() {
	l.quantity = l.quantity + 1
}
