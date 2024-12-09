package bag

import (
	"fmt"

	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/bag/item"
)

type Position int

type Bag struct {
	remedy []item.Item
	ball   []item.Item
	battle []item.Item
	berry  []item.Item
}

func NewBag() *Bag {
	return &Bag{
		remedy: []item.Item{},
		ball:   []item.Item{},
		battle: []item.Item{},
		berry:  []item.Item{},
	}
}

func (b *Bag) MustAddToRemedy(items ...item.Item) *Bag {
	for _, i := range items {
		if i.Category() != item.REMEDY {
			panic(fmt.Errorf("cannot add item %s to remedy bag", i.Name()))
		}

		b.remedy = append(b.remedy, i)
	}

	return b
}

func (b *Bag) MustAddToBerry(items ...item.Item) *Bag {
	for _, i := range items {
		if i.Category() != item.BERRY {
			panic(fmt.Errorf("cannot add item %s to berry bag", i.Name()))
		}
		b.berry = append(b.berry, i)

	}

	return b
}

func (b *Bag) SearchItem(itemEffect item.ItemEffect, minimumQuantity int) (item.Item, Position) {
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
