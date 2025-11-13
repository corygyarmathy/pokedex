package main

import (
	"fmt"
)

func commandMapFwrd(cfg *config, args ...string) error {
	result, err := cfg.pokeAPIClient.FetchLocationAreas(cfg.nextLocURL)
	if err != nil {
		return fmt.Errorf("fetching locations: %w", err)
	}

	cfg.nextLocURL = result.Next
	cfg.prevLocURL = result.Previous

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(cfg *config, args ...string) error {
	if cfg.prevLocURL == nil {
		fmt.Println("You're on the first page")
		return nil
	}

	result, err := cfg.pokeAPIClient.FetchLocationAreas(cfg.prevLocURL)
	if err != nil {
		return fmt.Errorf("fetching previous locations: %w", err)
	}

	cfg.nextLocURL = result.Next
	cfg.prevLocURL = result.Previous

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}
	return nil
}
