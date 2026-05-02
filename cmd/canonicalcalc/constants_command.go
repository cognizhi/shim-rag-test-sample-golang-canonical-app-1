package main

import (
	"fmt"
	"strings"

	"github.com/cognizhi/ragtest/canonical-calculator/internal/constants"
)

func init() {
	registerCommand(command{
		name:    "constant",
		summary: "Print a named mathematical constant",
		usage:   "constant <name>",
		run:     runConstantCommand,
	})
}

func runConstantCommand(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("constant expects exactly one name; supported names: %s", strings.Join(constants.Names(), ", "))
	}

	value, exists := constants.Lookup(args[0])
	if !exists {
		return fmt.Errorf("unknown constant %q; supported names: %s", args[0], strings.Join(constants.Names(), ", "))
	}

	fmt.Printf("%s = %s\n", strings.ToLower(args[0]), formatNumber(value))
	return nil
}
