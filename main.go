package main

import (
	"database/sql"
	"fmt"
	"os"

	"blogaggregator/internal/config"
	"blogaggregator/internal/database"
	_ "github.com/lib/pq"
)

// CONNECTION STRING
// postgres://postgres:postgres@localhost:5432/gator

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}
	defer db.Close()

	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	commands := commands{
		Commands: make(map[string]func(*state, command) error),
    }

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)

	if len(os.Args) < 2 {
		fmt.Println("You must provide a program and command name")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	argsSlice := os.Args[2:]

	command := command{
		Name:        cmdName,
		StringSlice: argsSlice,
    }

	err = commands.run(programState, command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
