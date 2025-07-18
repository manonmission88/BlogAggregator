package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	Config "github.com/manonmission88/BlogAggregator/internal/config"
	"github.com/manonmission88/BlogAggregator/internal/database"
	"github.com/manonmission88/BlogAggregator/internal/handler"
	"github.com/manonmission88/BlogAggregator/internal/state"
)

// read the config file and print on the terminal
func main() {

	cfg, err := Config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Default()
	}
	dbQueries := database.New(db)

	// update the new config
	appState := &state.State{
		Config:    &cfg,
		DbQueries: dbQueries,
	}
	cmds := handler.New()
	cmds.Register("login", handler.HandlerLogin)
	cmds.Register("register", handler.HandlerRegister)
	cmds.Register("reset", handler.HandlerReset)
	cmds.Register("users", handler.HandlerUsers)
	cmds.Register("agg", handler.HandlerAgg)

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("too few commands")
	}

	newCommands := &handler.Command{
		Name: args[1],
		Args: args[2:],
	}
	if err := cmds.Run(appState, *newCommands); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}
