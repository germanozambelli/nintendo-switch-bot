package player

import (
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/logger"
	"time"

	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/player/controller"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/player/controller/button"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/player/controller/stick"
)

type Player struct {
	logger     logger.Logger
	controller controller.Controller
	name       string
}

func NewPlayer(
	controller controller.Controller,
	name string,
	logger logger.Logger,
) *Player {
	return &Player{
		controller: controller,
		name:       name,
		logger:     logger,
	}
}

func (p *Player) PressHome() {
	p.logger.Debug("pressing home button")
	p.controller.PressButton(button.HOME)
}
func (p *Player) Nothing(duration time.Duration) {
	p.logger.Debug(fmt.Sprintf("doing nothing for %s", duration))
	time.Sleep(duration)
}

func (p *Player) Confirm() {
	p.logger.Debug("confirming")
	p.controller.PressButton(button.A)
}

func (p *Player) Cancel() {
	p.logger.Debug("canceling")
	p.controller.PressButton(button.B)
}

func (p *Player) MoveUp(duration time.Duration) {
	p.logger.Debug(fmt.Sprintf("moving up for %s", duration))
	p.controller.MoveStick(stick.LEFT, 0, 100, duration)
}

func (p *Player) MoveDown(duration time.Duration) {
	p.logger.Debug(fmt.Sprintf("moving down for %s", duration))
	p.controller.MoveStick(stick.LEFT, 0, -100, duration)
}

func (p *Player) MoveCameraToRight(duration time.Duration) {
	p.logger.Debug(fmt.Sprintf("moving camera to right for %s", duration))
	p.controller.MoveStick(stick.RIGHT, -100, 0, duration)
}

func (p *Player) MoveCameraToLeft(duration time.Duration) {
	p.logger.Debug(fmt.Sprintf("moving camera to left for %s", duration))
	p.controller.MoveStick(stick.RIGHT, 100, 0, duration)
}

func (p *Player) PressDown() {
	p.logger.Debug("pressing down")
	p.controller.PressButton(button.DOWN)
}

func (p *Player) PressUp() {
	p.logger.Debug("pressing up")
	p.controller.PressButton(button.UP)
}

func (p *Player) PressLeft() {
	p.logger.Debug("pressing left")
	p.controller.PressButton(button.LEFT)
}

func (p *Player) PressRight() {
	p.logger.Debug("pressing right")
	p.controller.PressButton(button.RIGHT)
}

func (p *Player) PressL() {
	p.logger.Debug("pressing L")
	p.controller.PressButton(button.L)
}

func (p *Player) PressMinus() {
	p.logger.Debug("pressing minus")
	p.controller.PressButton(button.MINUS)
}

func (p *Player) HoldDown(duration time.Duration) {
	p.logger.Debug(fmt.Sprintf("holding down for %s", duration))
	p.controller.HoldButton(button.DOWN, duration)
}

func (p *Player) HoldUp(duration time.Duration) {
	p.logger.Debug(fmt.Sprintf("holding up for %s", duration))
	p.controller.HoldButton(button.UP, duration)
}

func (p *Player) HoldLeft(duration time.Duration) {
	p.logger.Debug(fmt.Sprintf("holding left for %s", duration))
	p.controller.HoldButton(button.LEFT, duration)
}

func (p *Player) HoldRight(duration time.Duration) {
	p.logger.Debug(fmt.Sprintf("holding right for %s", duration))
	p.controller.HoldButton(button.RIGHT, duration)
}

func (p *Player) TakeAScreenShot() {
	p.logger.Debug("taking a screenshot")
	p.controller.PressButton(button.CAPTURE)
}

func (p *Player) Forever(doForever func()) {
	for {
		doForever()
	}
}
