package internal

import (
	"fmt"
	"math/rand"
	"os"
)

func commandExit(con *Config, params []string) error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandPokedex(con *Config, params []string) error {
	fmt.Println("Your Pokedex:")
	if len(con.CaughtPokemons) < 0 {
		fmt.Println("You did not catch any pokemons !")
	}
	for _, pokemon := range con.CaughtPokemons {
		fmt.Printf("	- %s", pokemon.Name)
	}
	return nil
}

func commandCatch(con *Config, params []string) error {
	if len(params) == 0 {
		fmt.Println("Please provide a pokemon to catch.")
		return nil
	}

	pokemon, err := con.Client.Pokemon.GetPokemon(params[1])

	if err != nil {
		fmt.Println("Pokemon Not Found")
		return nil
	}

	if _, exists := con.CaughtPokemons[pokemon.Name]; exists {
		fmt.Println("You already captured this pokemon !")
		return nil
	}

	fmt.Printf("\nThrowing a Pokeball at %s...\n", pokemon.Name)

	if randomInt := rand.Intn(pokemon.BaseExperience); randomInt >= pokemon.BaseExperience/3 {
		fmt.Printf("%s was not caught ! \n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught !\n", pokemon.Name)

	fmt.Println("You may now inspect it with the inspect command.")

	con.CaughtPokemons[pokemon.Name] = *pokemon

	return nil
}

func commandExplore(con *Config, params []string) error {
	if len(params) == 0 {
		fmt.Println("Please provide a location to explore.")
		return nil
	}

	areaName := params[1]
	locationArea, err := con.Client.Locations.GetLocationArea(areaName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil
	}

	if locationArea.PokemonEncounters == nil || len(locationArea.PokemonEncounters) == 0 {
		fmt.Printf("No pokemon found in %s.\n", areaName)
		return nil
	}

	fmt.Printf("Exploring %s...\n", areaName)
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Println("- " + encounter.Pokemon.Name)
	}
	return nil
}

func commandMap(con *Config, params []string) error {
	locationList, _ := con.Client.Locations.GetLocationAreaList(20, con.Mapoffset)
	fmt.Println(" ")

	for _, value := range locationList.Results {
		fmt.Println(value.Name)
	}

	con.Mapoffset += 1

	return nil
}

func commandMapB(con *Config, params []string) error {
	locationList, _ := con.Client.Locations.GetLocationAreaList(20, con.Mapoffset)
	fmt.Println(" ")

	if con.Mapoffset >= 0 {
		con.Mapoffset -= 1
	}

	for _, value := range locationList.Results {
		fmt.Println(value.Name)
	}

	return nil
}

func commandHelp(con *Config, params []string) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:\n")
	for name, cmd := range con.Commands {
		fmt.Printf("%s: %s\n", name, cmd.Description)
	}
	return nil
}

func commandInspect(con *Config, params []string) error {
	if len(params) == 0 {
		fmt.Println("\nYou should pass a pokemon for inspection")
		return nil
	}
	if _, exists := con.CaughtPokemons[params[1]]; !exists {
		fmt.Println("You did not caught this pokemon !")
		return nil
	}

	pokemon := con.CaughtPokemons[params[1]]
	pokemonStr := fmt.Sprintf(
		"Name: %s\nLevel: %d\nHeight:",
		pokemon.Name,
		pokemon.BaseExperience,
		pokemon.Height,
	)

	println(pokemonStr)

	return nil
}
