package main

import (
	"fmt"
)

func commandHelp(cfg *config, args []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()
	for _, c := range commands {
		fmt.Printf("%v: %v\n", c.name, c.description)
	}
	return nil
}
