package controller

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/player/controller/button"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/player/controller/stick"
	"time"
)

type Controller interface {
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
