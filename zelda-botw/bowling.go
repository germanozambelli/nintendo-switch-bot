package main

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/entity"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/service/controller/action"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/domain/value_object"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/infrastructure/service/screenshot_monitor"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/infrastructure/service/virtual_joycon"
	"time"
)

func main() {
	joycon := virtual_joycon.NewVirtualJoyCon()
	joycon.Connect()

	player := entity.NewPlayer(joycon, "Link")

	//
	////load game
	action.NewStartGame(1, joycon).Play()
	//////
	player.
		Nothing(15 * time.Second).
		Confirm().
		Nothing(1 * time.Second).
		Confirm().
		Nothing(2 * time.Second).
		Confirm().
		Nothing(25 * time.Second).
		Play()
	//////

	loop(player, joycon)
}

func loop(player *entity.Player, joycon controller.Controller) {
	//////first talk
	player.
		Confirm().
		Talk([]value_object.TalkStep{
			value_object.NewTalkStep(true, 3*time.Second, 75*time.Millisecond),
			value_object.NewTalkStep(true, 3*time.Second, 75*time.Millisecond),
			value_object.NewTalkStep(true, 2*time.Second, 6*time.Second),
			value_object.NewTalkStep(true, 4*time.Second, 2*time.Second),
			value_object.NewTalkStep(true, 3*time.Second, 75*time.Millisecond),
			value_object.NewTalkStep(true, 3*time.Second, 75*time.Millisecond),
			value_object.NewTalkStep(true, 2*time.Second, 75*time.Millisecond),
		}).
		Nothing(4 * time.Second).
		Play()

	//throw ball
	player.
		Confirm().
		Nothing(1 * time.Second).
		MoveUp(1 * time.Second).
		Confirm().
		Nothing(17 * time.Second).
		Play()

	//check state
	virtualmonitor := screenshot_monitor.NewScreenshotMonitor(joycon)
	strike := virtualmonitor.CurrentDialogContains([]string{"UAAAU!", "Uno strike!"})

	if strike {
		//get prize
		player.
			Talk([]value_object.TalkStep{
				value_object.NewTalkStep(true, 0*time.Second, 9*time.Second),
				value_object.NewTalkStep(true, 3*time.Second, 2*time.Second),
				value_object.NewTalkStep(true, 3*time.Second, 75*time.Millisecond),
				value_object.NewTalkStep(false, 3*time.Second, 1*time.Second),
				value_object.NewTalkStep(true, 3*time.Second, 75*time.Millisecond),
			}).
			Nothing(5 * time.Second).
			Play()
		loop(player, joycon)
	} else {
		player.
			Talk([]value_object.TalkStep{
				value_object.NewTalkStep(true, 0*time.Second, 9*time.Second),
				value_object.NewTalkStep(true, 3*time.Second, 2*time.Second),
			}).
			MoveDown(200 * time.Millisecond).
			Nothing(1 * time.Second).
			Confirm().
			Nothing(1 * time.Second).
			MoveUp(1 * time.Second).
			Confirm().
			Nothing(20 * time.Second).
			Talk([]value_object.TalkStep{
				value_object.NewTalkStep(true, 5*time.Second, 2*time.Second),
				value_object.NewTalkStep(true, 5*time.Second, 5*time.Second),
				value_object.NewTalkStep(true, 5*time.Second, 2*time.Second),
				value_object.NewTalkStep(false, 5*time.Second, 2*time.Second),
				value_object.NewTalkStep(false, 5*time.Second, 2*time.Second),
				value_object.NewTalkStep(false, 5*time.Second, 2*time.Second),
			}).
			Nothing(5 * time.Second).
			Play()
		loop(player, joycon)
	}
}
