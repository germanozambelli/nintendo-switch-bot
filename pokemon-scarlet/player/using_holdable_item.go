package player

import (
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/logger"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
	"time"
)

type UsingHoldableItem struct {
	player  *Player
	pokemon *pokemon.Pokemon
	logger  logger.Logger
}

func NewUsingHoldableItem(
	player *Player,
	pokemon *pokemon.Pokemon,
	logger logger.Logger,
) *UsingHoldableItem {
	return &UsingHoldableItem{
		player:  player,
		pokemon: pokemon,
		logger:  logger,
	}
}

func (u *UsingHoldableItem) Name() StateName {
	return USING_HOLDABLE_ITEM
}

func (u *UsingHoldableItem) UseHoldableItem() {
	if u.pokemon.Item() == nil {
		panic(fmt.Sprintf("pokemon %s has no item", u.pokemon.Name()))
	}

	applied := u.pokemon.Item().ApplyHoldingEffect(u.pokemon)
	if applied {
		u.player.Nothing(10 * time.Second)
	}

	u.player.SetState(NewChallengingAPokemon(u.player, u.pokemon, u.logger))
}
