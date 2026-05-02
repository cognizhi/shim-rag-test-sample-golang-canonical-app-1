package main

import (
	"fmt"

	"github.com/cognizhi/ragtest/canonical-calculator/internal/api"
)

func init() {
	registerCommand(command{
		name:    "serve",
		summary: "Start the HTTP calculator API",
		usage:   "serve [addr]",
		run:     runServeCommand,
	})
}

func runServeCommand(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("serve expects zero or one address argument")
	}

	address := ":8080"
	if len(args) == 1 {
		address = args[0]
	}

	fmt.Println("serving calculator API on", address)
	return api.Run(address)
}
