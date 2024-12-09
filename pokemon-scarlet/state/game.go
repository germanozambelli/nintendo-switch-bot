package state

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/joycon"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/logger"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/player"
	"time"
)

type Game struct {
	player     *player.Player
	controller *joycon.Controller
	logger     logger.Logger
}

func NewGame(
	player *player.Player,
	controller *joycon.Controller,
	logger logger.Logger,
) *Game {
	logger.Info("Game")

	return &Game{
		player:     player,
		controller: controller,
		logger:     logger,
	}
}

func (g *Game) Forever(do func(state *InFreeWorld)) {
	g.controller.PressHome()
	g.controller.Nothing(75 * time.Millisecond)

	inFreeWorld := NewInFreeWorld(g)

	for {
		do(inFreeWorld)
	}
}

func (g *Game) Once(do func(state *InFreeWorld)) {
	g.controller.PressHome()
	g.controller.Nothing(75 * time.Millisecond)

	inFreeWorld := NewInFreeWorld(g)

	do(inFreeWorld)
}
