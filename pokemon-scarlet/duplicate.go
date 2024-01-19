package main

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/entity"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/infrastructure/service/virtual_joycon"
	"time"
)

func main() {
	joycon := virtual_joycon.NewVirtualJoyCon()
	joycon.Connect()

	player := entity.NewPlayer(joycon, "Zamba")

	//
	////load game
	player.
		PressHome().
		Nothing(75 * time.Millisecond).
		Play()

	loop(player, joycon)
}

func loop(player *entity.Player, joycon controller.Controller) {
	player.
		PressL().
		MoveUp(time.Second * 4).
		Nothing(7 * time.Second).
		Nothing(1 * time.Second).
		Confirm().
		PressDown().
		PressDown().
		PressDown().
		Confirm().
		Nothing(17 * time.Second).
		PressUp().
		Confirm().
		Nothing(6 * time.Second).
		Play()

	loop(player, joycon)
}
