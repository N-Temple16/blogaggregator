package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.StringSlice) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	
	name := cmd.StringSlice[0]

	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}