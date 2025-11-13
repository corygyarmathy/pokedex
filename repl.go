package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

type config struct {
	pokeAPIClient *pokeapi.Client
	nextLocURL    *string
	prevLocURL    *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := getCommands()

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break // EOF or error: exit REPL
		}

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmd, ok := cmds[words[0]]
		if !ok {
			fmt.Printf("Unknown command: %v\n", words[0])
			continue
		}

		if err := cmd.callback(cfg, words[1:]...); err != nil {
			fmt.Fprintf(os.Stderr, "Command error: %v\n", err)
		}
	}

	// Check for scanner errors after loop
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Scanner error: %v\n", err)
	}
}

// Command names
const (
	cmdExit = "exit"
	cmdHelp = "help"
	cmdMap  = "map"
	cmdMapB = "mapb"
	cmdExp  = "explore"
)

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		cmdExit: {
			name:        cmdExit,
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		cmdHelp: {
			name:        cmdHelp,
			description: "Displays a help message",
			callback:    commandHelp,
		},
		cmdMap: {
			name:        cmdMap,
			description: "Displays the next page of locations",
			callback:    commandMapFwrd,
		},
		cmdMapB: {
			name:        cmdMapB,
			description: "Displays the previous page of locations",
			callback:    commandMapBack,
		},
		cmdExp: {
			name:        cmdExp + "<location-name>",
			description: "Displays the Pokemon found in a given location",
			callback:    commandExplore,
		},
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)

	return words
}
