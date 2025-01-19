package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/chaseplamoureux/blogaggregator/internal/config"
	"github.com/chaseplamoureux/blogaggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	ConfPointer *config.Config
	dbConn      *database.Queries
}

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Println("Not enough arguments provided")
		os.Exit(1)
	}

	userCommand := command{
		commandName: argsWithoutProg[0],
		commandArgs: argsWithoutProg[1:],
	}

	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	currentState := state{ConfPointer: &conf}

	db, err := sql.Open("postgres", currentState.ConfPointer.DB_URL)

	dbQueries := database.New(db)

	currentState.dbConn = dbQueries

	registeredCommands := commands{commandsMap: make(map[string]func(*state, command) error)}

	registeredCommands.register("login", handlerLogin)

	registeredCommands.register("register", handlerRegister)

	registeredCommands.register("reset", handlerReset)

	registeredCommands.register("users", handlerGetUsers)

	registeredCommands.register("agg", handlerAgg)

	registeredCommands.register("addfeed", middlewareLoggedIn(handlerAddFeed))

	registeredCommands.register("feeds", handlerFeeds)

	registeredCommands.register("follow", middlewareLoggedIn(handlerFollow))

	registeredCommands.register("following", middlewareLoggedIn(handlerFollowing))

	registeredCommands.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	err = registeredCommands.run(&currentState, userCommand)
	if err != nil {
		log.Fatal(err)
	}

}
