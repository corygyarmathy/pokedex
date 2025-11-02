package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeAPIClient: pokeClient,
	}
	startRepl(cfg)
}
