package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"net/http"
	"encoding/json"
)

type cliCommand struct {
	name 		string
	description string
	callback 	func(*Config) error
}

type Config struct {
	Next 		*string `json:"next"`
	Previous 	*string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}

var commands map[string]cliCommand

func main() {
	commands = map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Successively displays the names of the next 20 areas in the Pokemon world",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description:  "Successively displays the names of the previous 20 areas in the Pokemon world",
			callback: commandMapb,
		},
	}

	cfg := &Config{}

	fmt.Print("Pokedex > ")
	scnr := bufio.NewScanner(os.Stdin)
	
	for scnr.Scan() {
		text := cleanInput(scnr.Text())
		if len(text) == 0 {
			fmt.Print("Pokedex > ")
			continue
		}

		cmd, ok := commands[text[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
				fmt.Print("Pokedex > ")
				continue
			}
			if cmd.name == "exit" {
				os.Exit(0)
			}
		}
		fmt.Print("Pokedex > ")
	}
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	return nil
}

func commandHelp(cfg *Config) error {
	var cmds string
	for _, cmd := range commands {
		cmds += "\n" + cmd.name + ": " + cmd.description
	}
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n%s\n", cmds)
	return nil
}

func commandMap(cfg *Config) error {
	var url string
	if cfg.Next != nil {
		url = *cfg.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area?limit=20"
	}
	
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error getting response: %w", err)
	}
	defer response.Body.Close()

	var apiResponse Config
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		return fmt.Errorf("error decoding response body: %w", err)
	}

	for _, area := range apiResponse.Results {
		fmt.Println(area.Name)
	}

	cfg.Next = apiResponse.Next
	cfg.Previous = apiResponse.Previous

	return nil
}

func commandMapb(cfg *Config) error {
	var url string
	if cfg.Previous != nil {
		url = *cfg.Previous
	} else {
		url = "https://pokeapi.co/api/v2/location-area?limit=20"
	}

	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error getting response: %w", err)
	}
	defer response.Body.Close()

	var apiResponse Config
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	if err != nil {
		return fmt.Errorf("error decoding response body: %w", err)
	}

	for _, area := range apiResponse.Results {
		fmt.Println(area.Name)
	}

	cfg.Next = apiResponse.Next
	cfg.Previous = apiResponse.Previous

	return nil
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}