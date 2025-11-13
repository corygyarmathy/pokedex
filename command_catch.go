package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("provide a single word (pokemon-name): %v", args)
	}

	name := args[0]
	fmt.Println("Throwing a Pokeball at " + name + "...")
	result, err := cfg.pokeAPIClient.FetchPokemon(name)
	if err != nil {
		return fmt.Errorf("fetching pokemon: %w", err)
	}

	catchChance := (0.5 * float32(result.BaseExperience)) / 100
	random := rand.Float32()
	if random > catchChance {
		fmt.Println(name + " was caught!")
	} else {
		fmt.Println(name + " escaped!")
	}

	cfg.caughtPokemon[result.Name] = *result
	return nil
}
