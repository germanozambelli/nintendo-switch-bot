package item

import "github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"

type ItemCategory int

const (
	REMEDY ItemCategory = iota
	BALL
	BATTLE
	BERRY
)

type ItemEffect int

const (
	PP_UP ItemEffect = iota
)

type Item interface {
	Name() string
	Quantity() int
	Category() ItemCategory
	Effect() ItemEffect
	IncreaseQuantity()
}

type ConsumableItem interface {
	Item
	Consume()
}

type SpellApplicable interface {
	Item
	ApplyToSpell(spell *pokemon.Spell)
}

type PokemonApplicable interface {
	Item
	ApplyToPokemon(pokemon *pokemon.Pokemon)
}
