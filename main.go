package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/levabu/gator/internal/commands"
	"github.com/levabu/gator/internal/config"
	"github.com/levabu/gator/internal/database"
	"github.com/levabu/gator/internal/state"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config: %v\n", err)
		return
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Println("error: couldn't connect to the database")
		os.Exit(1)
	}
	defer db.Close()

	dbQueries := database.New(db)

	st := state.State{
		Config: &cfg,
		DB: dbQueries,
	}

	cmds := commands.Commands{
		Cmds: make(map[string]func(*state.State, commands.Command) error),
	}
	cmds.Register("login", commands.HandlerLogin)
	cmds.Register("register", commands.HandlerRegister)
	cmds.Register("reset", commands.HandlerReset)
	cmds.Register("users", commands.HandlerUsers)
	cmds.Register("agg", commands.HandlerAgg)
	cmds.Register("addfeed", middlewareLoggedIn(commands.HandlerAddFeed))
	cmds.Register("feeds", commands.HandlerFeeds)
	cmds.Register("follow", middlewareLoggedIn(commands.HandlerFollow))
	cmds.Register("following", middlewareLoggedIn(commands.HandlerFollowing))
	cmds.Register("unfollow", middlewareLoggedIn(commands.HandlerUnfollow))
	cmds.Register("browse", middlewareLoggedIn(commands.HandlerBrowse))

	args := os.Args
	if len(args) < 2 {
		log.Fatal("usage: <command> [args...]")
	}

	commandName := args[1]
	commandArgs := args[2:]
	cmd := commands.Command {
		Name: commandName,
		Args: commandArgs,
	}

	if err := cmds.Run(&st, cmd); err != nil {
		log.Fatal(err)
	}
}