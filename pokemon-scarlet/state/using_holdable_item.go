package state

import (
	"fmt"
	"time"
)

type UsingHoldableItem struct {
	*ChallengingAPokemon
}

func NewUsingHoldableItem(
	state *ChallengingAPokemon,
) *UsingHoldableItem {
	state.logger.Info("UsingHoldableItem")

	return &UsingHoldableItem{
		ChallengingAPokemon: state,
	}
}

func (u *UsingHoldableItem) UseHoldableItem() {
	if u.pokemon.Item() == nil {
		panic(fmt.Sprintf("pokemon %s has no item", u.pokemon.Name()))
	}

	if u.pokemon.Item().ApplyHoldingEffect(u.pokemon) {
		u.controller.Nothing(10 * time.Second)
	}
}
