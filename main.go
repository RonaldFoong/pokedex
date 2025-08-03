package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	fmt.Print("Pokedex > ")
	scnr := bufio.NewScanner(os.Stdin)

	for scnr.Scan() {
		text := cleanInput(scnr.Text())
		fmt.Println("Your command was: " + text[0])
		fmt.Print("Pokedex > ")
	}

}

func cleanInput(text string) []string {
	return strings.Fields(text)
}