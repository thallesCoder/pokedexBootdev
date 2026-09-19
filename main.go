package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal"
	"strings"
)

var commandList map[string]internal.CliCommand

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commandList = internal.GetConfig().Commands
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		if len(input) <= 0 {
			continue
		}

		cleanedInput := cleanInput(input)

		commandName := cleanedInput[0]

		if command, exists := commandList[commandName]; exists {
			err := command.Callback(internal.GetConfig(), cleanedInput)
			if err != nil {
				fmt.Printf("Error: %v \n", err)
			}
		} else {
			err := unknow()
			if err != nil {
				return
			}
		}
	}
}

func unknow() error {
	fmt.Println("Unknow command.")
	return nil
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
