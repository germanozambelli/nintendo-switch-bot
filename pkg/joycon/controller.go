package joycon

import (
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/joycon/button"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/joycon/stick"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/logger"
	"time"
)

type JoyCon interface {
	PressButton(button button.Button)
	HoldButton(
		button button.Button,
		duration time.Duration,
	)
	MoveStick(
		stick stick.Stick,
		xPosition int,
		yPosition int,
		duration time.Duration,
	)
}

type Controller struct {
	logger logger.Logger
	joyCon JoyCon
}

func NewController(
	joyCon JoyCon,
	logger logger.Logger,
) *Controller {
	return &Controller{
		joyCon: joyCon,
		logger: logger,
	}
}

func (c *Controller) PressHome() {
	c.logger.Debug("pressing home button")
	c.joyCon.PressButton(button.HOME)
}

func (c *Controller) Nothing(duration time.Duration) {
	c.logger.Debug(fmt.Sprintf("doing nothing for %s", duration))
	time.Sleep(duration)
}

func (c *Controller) Confirm() {
	c.logger.Debug("confirming")
	c.joyCon.PressButton(button.A)
}

func (c *Controller) Cancel() {
	c.logger.Debug("canceling")
	c.joyCon.PressButton(button.B)
}

func (c *Controller) MoveUp(duration time.Duration) {
	c.logger.Debug(fmt.Sprintf("moving up for %s", duration))
	c.joyCon.MoveStick(stick.LEFT, 0, 100, duration)
}

func (c *Controller) MoveDown(duration time.Duration) {
	c.logger.Debug(fmt.Sprintf("moving down for %s", duration))
	c.joyCon.MoveStick(stick.LEFT, 0, -100, duration)
}

func (c *Controller) MoveCameraToRight(duration time.Duration) {
	c.logger.Debug(fmt.Sprintf("moving camera to right for %s", duration))
	c.joyCon.MoveStick(stick.RIGHT, -100, 0, duration)
}

func (c *Controller) MoveCameraToLeft(duration time.Duration) {
	c.logger.Debug(fmt.Sprintf("moving camera to left for %s", duration))
	c.joyCon.MoveStick(stick.RIGHT, 100, 0, duration)
}

func (c *Controller) PressDown() {
	c.logger.Debug("pressing down")
	c.joyCon.PressButton(button.DOWN)
}

func (c *Controller) PressUp() {
	c.logger.Debug("pressing up")
	c.joyCon.PressButton(button.UP)
}

func (c *Controller) PressLeft() {
	c.logger.Debug("pressing left")
	c.joyCon.PressButton(button.LEFT)
}

func (c *Controller) PressRight() {
	c.logger.Debug("pressing right")
	c.joyCon.PressButton(button.RIGHT)
}

func (c *Controller) PressL() {
	c.logger.Debug("pressing L")
	c.joyCon.PressButton(button.L)
}

func (c *Controller) PressMinus() {
	c.logger.Debug("pressing minus")
	c.joyCon.PressButton(button.MINUS)
}

func (c *Controller) HoldDown(duration time.Duration) {
	c.logger.Debug(fmt.Sprintf("holding down for %s", duration))
	c.joyCon.HoldButton(button.DOWN, duration)
}

func (c *Controller) HoldUp(duration time.Duration) {
	c.logger.Debug(fmt.Sprintf("holding up for %s", duration))
	c.joyCon.HoldButton(button.UP, duration)
}

func (c *Controller) HoldLeft(duration time.Duration) {
	c.logger.Debug(fmt.Sprintf("holding left for %s", duration))
	c.joyCon.HoldButton(button.LEFT, duration)
}

func (c *Controller) HoldRight(duration time.Duration) {
	c.logger.Debug(fmt.Sprintf("holding right for %s", duration))
	c.joyCon.HoldButton(button.RIGHT, duration)
}

func (c *Controller) TakeAScreenShot() {
	c.logger.Debug("taking a screenshot")
	c.joyCon.PressButton(button.CAPTURE)
}
