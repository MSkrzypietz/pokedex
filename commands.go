package main

import (
	"errors"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func createCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays the names of all pokemons for the given location",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandMap(cfg *config, args ...string) error {
	locationAreasResp, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationAreasUrl)
	if err != nil {
		return err
	}

	fmt.Println()
	for _, locationArea := range locationAreasResp.Results {
		fmt.Println(locationArea.Name)
	}
	cfg.nextLocationAreasUrl = locationAreasResp.Next
	cfg.prevLocationAreasUrl = locationAreasResp.Previous
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationAreasUrl == nil {
		return errors.New("cannot go back before the start")
	}

	locationAreas, err := cfg.pokeapiClient.GetLocationAreas(cfg.prevLocationAreasUrl)
	if err != nil {
		return err
	}

	fmt.Println()
	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}
	cfg.nextLocationAreasUrl = locationAreas.Next
	cfg.prevLocationAreasUrl = locationAreas.Previous
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no location name supplied")
	}

	location, err := cfg.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, pokemon := range location.Pokemons {
		fmt.Printf("  - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range createCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
