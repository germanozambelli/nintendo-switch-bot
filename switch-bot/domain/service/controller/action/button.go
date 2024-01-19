package action

import (
	"fmt"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller/button"
	"time"
)

type PressButton struct {
	button     button.Button
	duration   time.Duration
	controller controller.Controller
}

func NewPressButton(button button.Button, duration time.Duration, controller controller.Controller) PressButton {
	return PressButton{button: button, duration: duration, controller: controller}
}

func (pressButton PressButton) Play() {
	(pressButton.controller).PressButton(pressButton.button, pressButton.duration)
	fmt.Println(fmt.Sprintf("pressed button %#v", pressButton.button))
}
