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
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if input == "" {
			continue
		}

		if command, ok := commands[input]; ok {
			if err := command.callback(cfg); err != nil {
				fmt.Printf("Error encountered: %s\n", err.Error())
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
