package main

import (
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/bag/item"
	"log/slog"
	"os"
	"time"

	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/joycon"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/nxbt-joycon"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/bag"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/player"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/state"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	virtualJoyCon, err := nxbt_joycon.NewVirtualJoyCon("192.168.1.96", logger)
	if err != nil {
		panic(err)
	}

	controller := joycon.NewController(virtualJoyCon, logger)

	rotom := pokemon.NewPokemon(
		"Rotom",
		nil,
		pokemon.NewSpell("Trick", 2, 16, 2),
	)

	team, err := pokemon.NewTeam(
		rotom,
	)

	bag := bag.NewBag().
		MustAddToRemedy(
			item.NewElisirMax(646),
		)

	if err != nil {
		panic(err)
	}

	p := player.NewPlayer(
		team,
		bag,
	)

	state.
		NewGame(p, controller, logger).
		Forever(
			func(state *state.InFreeWorld) {
				state.
					StartABattle(4 * time.Second).
					ChooseAPokemon(rotom).
					UseSpell("Trick").
					RunAway()
			},
		)
}
