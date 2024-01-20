package pokemon

type Pokemon struct {
	name   string
	spells []*Spell
	item   HoldableItem
}

func NewPokemon(
	name string,
	holdableItem HoldableItem,
	spells ...*Spell,
) *Pokemon {
	return &Pokemon{
		name:   name,
		item:   holdableItem,
		spells: spells,
	}
}

func (p *Pokemon) Name() string {
	return p.name
}

func (p *Pokemon) SpellBy(name string) (*Spell, int) {
	for i, s := range p.spells {
		if s.name == name {
			return s, i + 1
		}
	}

	return nil, 0
}

func (p *Pokemon) Spells() []*Spell {
	return p.spells
}

func (p *Pokemon) Item() HoldableItem {
	return p.item
}
