package entity

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller/action"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller/button"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller/stick"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/value_object"
	"time"
)

type Player struct {
	controller   controller.Controller
	state        value_object.State
	actionBuffer []action.Action
	name         string
}

func NewPlayer(controller controller.Controller, name string) *Player {
	return &Player{
		controller:   controller,
		name:         name,
		state:        value_object.WAITING,
		actionBuffer: []action.Action{},
	}
}

func (player *Player) Nothing(duration time.Duration) *Player {
	nothingAction := action.NewNothing(duration)
	player.actionBuffer = append(player.actionBuffer, nothingAction)
	return player
}

func (player *Player) Talk(steps []value_object.TalkStep) *Player {
	talkAction := action.NewTalk(steps, player.controller)
	player.actionBuffer = append(player.actionBuffer, talkAction)
	return player
}

func (player *Player) Confirm() *Player {
	confirmAction := action.NewPressButton(button.A, 75*time.Millisecond, player.controller)
	player.actionBuffer = append(player.actionBuffer, confirmAction)
	return player
}

func (player *Player) Cancel() *Player {
	cancelAction := action.NewPressButton(button.B, 75*time.Millisecond, player.controller)
	player.actionBuffer = append(player.actionBuffer, cancelAction)
	return player
}

func (player *Player) MoveUp(duration time.Duration) *Player {
	moveUpAction := action.NewMoveStick(stick.LEFT, 0, 100, duration, player.controller)
	player.actionBuffer = append(player.actionBuffer, moveUpAction)
	return player
}

func (player *Player) MoveDown(duration time.Duration) *Player {
	moveDownAction := action.NewMoveStick(stick.LEFT, 0, -100, duration, player.controller)
	player.actionBuffer = append(player.actionBuffer, moveDownAction)
	return player
}

func (player *Player) Play() *Player {
	combinedAction := action.NewCombine(player.actionBuffer)
	combinedAction.Play()
	player.actionBuffer = []action.Action{}

	return player
}
