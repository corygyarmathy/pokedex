package main

import (
	"pokedex/internal/pokeapi"
	"pokedex/internal/pokecache"
	"time"
)

func main() {
	c := pokecache.NewCache(5 * time.Minute)
	defer c.Close()

	pokeClient := pokeapi.NewClient(5*time.Second, c)
	cfg := &config{
		pokeAPIClient: pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}
	startRepl(cfg)
}
