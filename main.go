package main

import (
	"fmt"
	"log"
	Config "main/internal/config"
)

// read the config file and print on the terminal
func main() {

	cfg, err := Config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	err = cfg.SetUser("Manish")
	if err != nil {
		log.Fatal("couldnot set the field")
	}
	fmt.Printf("db: %s\n", cfg.DbUrl)
	fmt.Printf("User: %s\n", cfg.CurrentUser)

}
