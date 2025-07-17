package handler

// command field

import (
	"fmt"

	"github.com/manonmission88/BlogAggregator/internal/state"
)

// Command field
type Command struct {
	Name string
	Args []string
}

// store all the handlers
type Commands struct {
	Handlers map[string]func(*state.State, Command) error
}

// instaniate
func New() *Commands {
	return &Commands{
		Handlers: map[string]func(*state.State, Command) error{},
	}
}

// Run executes a command
func (c *Commands) Run(s *state.State, cmd Command) error {
	handler, ok := c.Handlers[cmd.Name]
	if ok {
		return handler(s, cmd)
	}
	return fmt.Errorf("unknown command %s", cmd.Name)

}

// register a new handler
func (c *Commands) Register(name string, f func(*state.State, Command) error) {
	c.Handlers[name] = f // add the handlers
}
