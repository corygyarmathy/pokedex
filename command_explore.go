package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("provide a single word (location-name): %v", args)
	}

	location := args[0]
	fmt.Println("Exploring " + location + "...")
	result, err := cfg.pokeAPIClient.FetchLocationArea(location)
	if err != nil {
		return fmt.Errorf("fetching locations: %w", err)
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range result.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}
	return nil
}
