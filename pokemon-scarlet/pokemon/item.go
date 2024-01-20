package pokemon

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

type SpellApplicable interface {
	Item
	ApplyToSpell(spell *Spell)
}

type PokemonApplicable interface {
	Item
	ApplyToPokemon(pokemon *Pokemon)
}

type ConsumableItem interface {
	Item
	Consume()
}

type HoldableItemApplied bool

type HoldableItem interface {
	Item
	ApplyHoldingEffect(pokemon *Pokemon) HoldableItemApplied
}
