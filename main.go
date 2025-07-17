package main

import (
	"fmt"
	"log"
	Config "main/internal/config"
	"main/internal/handler"
	"main/internal/state"
	"os"
)

// read the config file and print on the terminal
func main() {

	cfg, err := Config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	// update the new config
	appState := &state.State{
		Config: &cfg,
	}
	cmds := handler.New()
	cmds.Register("login", handler.HandlerLogin)

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
