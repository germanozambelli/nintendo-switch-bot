package action

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller/stick"
	"time"
)

type MoveStick struct {
	stick      stick.Stick
	yPosition  int
	xPosition  int
	duration   time.Duration
	controller controller.Controller
}

func NewMoveStick(
	stick stick.Stick,
	xPosition int,
	yPosition int,
	duration time.Duration,
	controller controller.Controller,
) MoveStick {
	return MoveStick{
		stick:      stick,
		yPosition:  yPosition,
		xPosition:  xPosition,
		duration:   duration,
		controller: controller,
	}
}

func (moveStick MoveStick) Play() {
	(moveStick.controller).MoveStick(moveStick.stick, moveStick.xPosition, moveStick.yPosition, moveStick.duration)
}
