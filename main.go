package main

import (
	"fmt"
	"blogaggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = cfg.SetUser("nigel")
	if err != nil {
		fmt.Println("Error:", err)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(cfg)
}
