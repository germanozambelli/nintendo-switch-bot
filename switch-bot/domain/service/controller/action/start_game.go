package action

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller/button"
	"time"
)

type StartGame struct {
	index      int
	controller controller.Controller
}

func NewStartGame(index int, controller controller.Controller) StartGame {
	return StartGame{index: index, controller: controller}
}

func (startGame StartGame) Play() {
	NewCombine([]Action{
		NewPressButton(button.HOME, 75*time.Millisecond, startGame.controller),
		NewRepeat(
			NewPressButton(button.DOWN, 75*time.Millisecond, startGame.controller),
			2,
		),
		NewRepeat(
			NewPressButton(button.LEFT, 75*time.Millisecond, startGame.controller),
			6,
		),
		NewPressButton(button.UP, 75*time.Millisecond, startGame.controller),
		NewRepeat(
			NewPressButton(button.RIGHT, 75*time.Millisecond, startGame.controller),
			startGame.index-1,
		),
		NewPressButton(button.A, 75*time.Millisecond, startGame.controller),
		NewNothing(1 * time.Second),
		NewPressButton(button.A, 75*time.Millisecond, startGame.controller),
	}).Play()
}
