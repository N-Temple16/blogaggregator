package main

import (
	"errors"
)

type command struct {
	Name        string
	StringSlice []string
}

type commands struct {
	Commands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.Commands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	val, ok := c.Commands[cmd.Name]
	if !ok {
		return errors.New("that command does not exist")
	}
	err := val(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
