package controller

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller/button"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller/stick"
	"time"
)

type Controller interface {
	Connect() error
	PressButton(button button.Button, duration time.Duration)
	MoveStick(stick stick.Stick, xPosition int, yPosition int, duration time.Duration)
}
