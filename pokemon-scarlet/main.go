package main

import (
	"fmt"
	bag2 "github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/bag"
	item2 "github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/item"
	"log/slog"
	"os"
	"time"

	nxbt_joycon "github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/nxbt-joycon"
	"github.com/germanozambelli/nintendo-switch-bot/switch-bot/pkg/player"
	pokemonPlayer "github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/player"
	pokemon "github.com/germanozambelli/nintendo-switch-bot/switch-bot/pokemon-scarlet/pokemon"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	joycon, err := nxbt_joycon.NewVirtualJoyCon("192.168.1.96", logger)

	if err != nil {
		panic(err)
	}

	rotom := pokemon.NewPokemon(
		"Rotom",
		nil,
		pokemon.NewSpell("Thunderbolt", 20, 20, 1),
		pokemon.NewSpell("Thunder Wave", 20, 20, 1),
		pokemon.NewSpell("Trick", 0, 16, 2),
		pokemon.NewSpell("Substitute", 10, 20, 1),
	)

	skeleridge := pokemon.NewPokemon(
		"Skeleridge",
		nil,
		pokemon.NewSpell("Thunderbolt", 20, 20, 1),
		pokemon.NewSpell("Thunder Wave", 20, 20, 1),
		pokemon.NewSpell("Trick", 8, 16, 1),
		pokemon.NewSpell("Substitute", 10, 20, 1),
	)

	team, err := pokemon.NewTeam(
		rotom,
		skeleridge,
	)

	elisirMax := item2.NewElisirMax(200)

	bag := bag2.NewBag().
		MustAddToRemedy(
			item2.NewEtereMax(1),
			elisirMax,
		).
		MustAddToBerry(
			item2.NewLeppaBerry(1),
		)

	if err != nil {
		panic(err)
	}

	p := pokemonPlayer.NewPlayer(
		logger,
		player.NewPlayer(joycon, "Ash", logger),
		team,
		bag,
	)

	p.PressHome()
	p.Nothing(75 * time.Millisecond)

	p.Forever(func() {
		logger.Info(fmt.Sprintf("elisir max quantity: %d", elisirMax.Quantity()))
		p.StartABattle(4 * time.Second)
		p.ChooseAPokemon(rotom)
		p.UseSpell("Trick")
		elisirMax.IncreaseQuantity()
		p.RunAway()
	})
}
