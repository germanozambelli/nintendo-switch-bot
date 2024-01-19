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

func (player *Player) StartAGame(gamePosition int) *Player {
	a := action.NewStartGame(gamePosition, player.controller)
	player.actionBuffer = append(player.actionBuffer, a)
	return player
}

func (player *Player) PressHome() *Player {
	a := action.NewPressButton(button.HOME, 75*time.Millisecond, player.controller)
	player.actionBuffer = append(player.actionBuffer, a)
	return player
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

func (player *Player) MoveCameraToRight(duration time.Duration) *Player {
	a := action.NewMoveStick(stick.RIGHT, -100, 0, duration, player.controller)
	player.actionBuffer = append(player.actionBuffer, a)
	return player
}

func (player *Player) MoveCameraToLeft(duration time.Duration) *Player {
	a := action.NewMoveStick(stick.RIGHT, 100, 0, duration, player.controller)
	player.actionBuffer = append(player.actionBuffer, a)
	return player
}

func (player *Player) PressDown() *Player {
	pressDownAction := action.NewPressButton(button.DOWN, 75*time.Millisecond, player.controller)
	player.actionBuffer = append(player.actionBuffer, pressDownAction)
	return player
}

func (player *Player) PressUp() *Player {
	pressUpAction := action.NewPressButton(button.UP, 75*time.Millisecond, player.controller)
	player.actionBuffer = append(player.actionBuffer, pressUpAction)
	return player
}

func (player *Player) HoldDown(duration time.Duration) *Player {
	holdDownAction := action.NewPressButton(button.DOWN, duration, player.controller)
	player.actionBuffer = append(player.actionBuffer, holdDownAction)
	return player
}

func (player *Player) HoldUp(duration time.Duration) *Player {
	holdUpAction := action.NewPressButton(button.UP, duration, player.controller)
	player.actionBuffer = append(player.actionBuffer, holdUpAction)
	return player
}

func (player *Player) PressLeft() *Player {
	pressLeftAction := action.NewPressButton(button.LEFT, 75*time.Millisecond, player.controller)
	player.actionBuffer = append(player.actionBuffer, pressLeftAction)
	return player
}

func (player *Player) PressRight() *Player {
	pressRightAction := action.NewPressButton(button.RIGHT, 75*time.Millisecond, player.controller)
	player.actionBuffer = append(player.actionBuffer, pressRightAction)
	return player
}

func (player *Player) PressL() *Player {
	a := action.NewPressButton(button.L, 75*time.Millisecond, player.controller)
	player.actionBuffer = append(player.actionBuffer, a)
	return player
}

func (player *Player) HoldLeft(duration time.Duration) *Player {
	holdLeftAction := action.NewPressButton(button.LEFT, duration, player.controller)
	player.actionBuffer = append(player.actionBuffer, holdLeftAction)
	return player
}

func (player *Player) HoldRight(duration time.Duration) *Player {
	holdRightAction := action.NewPressButton(button.RIGHT, duration, player.controller)
	player.actionBuffer = append(player.actionBuffer, holdRightAction)
	return player
}

func (player *Player) Play() *Player {
	combinedAction := action.NewCombine(player.actionBuffer)
	combinedAction.Play()
	player.actionBuffer = []action.Action{}

	return player
}
