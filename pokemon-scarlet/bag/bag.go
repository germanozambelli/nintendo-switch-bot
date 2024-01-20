package bag

import (
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
)

type Position int

type Bag struct {
	remedy []pokemon.Item
	ball   []pokemon.Item
	battle []pokemon.Item
	berry  []pokemon.Item
}

func NewBag() *Bag {
	return &Bag{
		remedy: []pokemon.Item{},
		ball:   []pokemon.Item{},
		battle: []pokemon.Item{},
		berry:  []pokemon.Item{},
	}
}

func (b *Bag) MustAddToRemedy(items ...pokemon.Item) *Bag {
	for _, i := range items {
		if i.Category() != pokemon.REMEDY {
			panic(fmt.Errorf("cannot add item %s to remedy bag", i.Name()))
		}

		b.remedy = append(b.remedy, i)
	}

	return b
}

func (b *Bag) MustAddToBerry(items ...pokemon.Item) *Bag {
	for _, i := range items {
		if i.Category() != pokemon.BERRY {
			panic(fmt.Errorf("cannot add item %s to berry bag", i.Name()))
		}
		b.berry = append(b.berry, i)

	}

	return b
}

func (b *Bag) SearchItem(itemEffect pokemon.ItemEffect, minimumQuantity int) (pokemon.Item, Position) {
	for position, i := range b.berry {
		if i.Effect() == itemEffect && i.Quantity() >= minimumQuantity {
			return i, Position(position + 1)
		}
	}

	for position, i := range b.remedy {
		if i.Effect() == itemEffect && i.Quantity() >= minimumQuantity {
			return i, Position(position + 1)
		}
	}

	return nil, -1
}
