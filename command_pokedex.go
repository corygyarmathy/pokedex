package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	// Print pokedex contents
	fmt.Println("Your pokedex: ")
	for _, p := range cfg.caughtPokemon {
		fmt.Printf(" - %v\n", p.Name)
	}
	return nil
}
