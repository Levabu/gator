package commands

import (
	"fmt"

	"github.com/levabu/gator/internal/state"
)

type Commands struct {
	Cmds map[string]func(*state.State, Command) error
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) {
	c.Cmds[name] = f
}

func (c *Commands) Run(s *state.State, cmd Command) error {
	handler, ok := c.Cmds[cmd.Name]
	if !ok {
		return fmt.Errorf("command %s not registered", cmd.Name)
	}

	if err := handler(s, cmd); err != nil {
		return err
	}

	return nil
}
