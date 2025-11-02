package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {
	startRepl()
	pokeClient := pokeapi.NewClient(5 * time.Second)
		pokeAPIClient: pokeClient,
	}
}
