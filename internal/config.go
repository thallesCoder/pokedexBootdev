package internal

import (
	pokego "github.com/JoshGuarino/PokeGo/pkg"
	"github.com/JoshGuarino/PokeGo/pkg/models"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(con *Config, params []string) error
}

type Config struct {
	Commands       map[string]CliCommand
	Mapoffset      int
	CaughtPokemons map[string]models.Pokemon
	Client         pokego.PokeGo
}

var conf = Config{
	Commands: map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the map",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the mapb",
			Callback:    commandMapB,
		},
		"explore": {
			Name:        "explore",
			Description: "Displays the creatures in area.",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Try to catch a pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspects a pokemon catch",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Displays your pokemons",
			Callback:    commandPokedex,
		},
	},
	CaughtPokemons: map[string]models.Pokemon{},
	Mapoffset:      0,
	Client:         pokego.NewClient(),
}

func GetConfig() *Config {
	return &conf
}
