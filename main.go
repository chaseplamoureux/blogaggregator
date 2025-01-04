package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chaseplamoureux/blogaggregator/internal/config"
)

type state struct {
	ConfPointer *config.Config
}

func main() {

	argsWithoutProg := os.Args[1:]
	
	if len(argsWithoutProg) <= 1 {
		fmt.Println("Not enough arguments provided")
		os.Exit(1)
	}

	userCommand := command {
		commandName: argsWithoutProg[0],
		commandArgs: argsWithoutProg[1:],
	}


	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	currentState := state{ConfPointer: &conf}

	registeredCommands := commands{commandsMap: make(map[string]func(*state, command) error) }

	registeredCommands.register("login", handlerLogin)

	err = registeredCommands.run(&currentState, userCommand)
	if err != nil {
		fmt.Println(err)
	}

}
