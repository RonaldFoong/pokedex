package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type cliCommand struct {
	name 		string
	description string
	callback 	func() error
}

func main() {
	commands := map[string]cliCommand{
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
	}
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
			fmt.Println(cmd.callback())
			if cmd.name == "exit" {
				os.Exit(0)
			}
		}

		fmt.Print("Pokedex > ")
	}

}

func commandExit() error {
	return fmt.Errorf("Closing the Pokedex... Goodbye!")
}

func commandHelp() error {
	return fmt.Errorf(
`Welcome to the Pokedex!
Usage: 

help: Displays a help message
exit: Exit the Pokedex`)
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}