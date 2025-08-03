package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world")
}

func cleanInput(text string) []string {
	return strings.Fields(text)
}