package main

import (
	"blogaggregator/internal/config"
	"fmt"
	"os"
)

// CONNECTION STRING
// postgres://postgres:postgres@localhost:5432/gator

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error:", err)
	}

	programState := &state{
		cfg: &cfg,
	}

	commands := commands{
		Commands: make(map[string]func(*state, command) error),
    }

	commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("You must provide a program and command name")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	argsSlice := os.Args[2:]

	command := command{
		cmdName,
		argsSlice,
    }

	err = commands.run(programState, command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
