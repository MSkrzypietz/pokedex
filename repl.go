package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	commands := createCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputs := cleanInput(scanner.Text())
		if len(inputs) == 0 {
			continue
		}

		if command, ok := commands[inputs[0]]; ok {
			if err := command.callback(cfg, inputs[1:]...); err != nil {
				fmt.Printf("Error encountered: %s\n", err.Error())
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(input string) []string {
	cleaned := strings.ToLower(strings.TrimSpace(input))
	return strings.Fields(cleaned)
}
