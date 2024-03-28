package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func createCommands() map[string]cliCommand {
	return map[string]cliCommand{
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

func commandHelp() error {
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

func commandExit() error {
	os.Exit(0)
	return nil
}