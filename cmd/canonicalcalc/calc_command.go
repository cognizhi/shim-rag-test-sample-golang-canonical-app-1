package main

import (
	"fmt"

	"github.com/cognizhi/ragtest/canonical-calculator/internal/calculator"
)

func init() {
	registerCommand(command{
		name:    "add",
		summary: "Add two numbers",
		usage:   "add <left> <right>",
		run: func(args []string) error {
			return runBinaryOperation(args, calculator.Add)
		},
	})
	registerCommand(command{
		name:    "sub",
		summary: "Subtract the second number from the first",
		usage:   "sub <left> <right>",
		run: func(args []string) error {
			return runBinaryOperation(args, calculator.Subtract)
		},
	})
	registerCommand(command{
		name:    "mul",
		summary: "Multiply two numbers",
		usage:   "mul <left> <right>",
		run: func(args []string) error {
			return runBinaryOperation(args, calculator.Multiply)
		},
	})
	registerCommand(command{
		name:    "div",
		summary: "Divide the first number by the second",
		usage:   "div <left> <right>",
		run: func(args []string) error {
			return runBinaryOperation(args, calculator.Divide)
		},
	})
}

func runBinaryOperation(args []string, operation calculator.Operation) error {
	if len(args) != 2 {
		return fmt.Errorf("%s expects exactly two operands", operation)
	}

	left, err := parseFloatArgument(args[0])
	if err != nil {
		return err
	}

	right, err := parseFloatArgument(args[1])
	if err != nil {
		return err
	}

	result, err := calculator.Apply(operation, left, right)
	if err != nil {
		return err
	}

	fmt.Println(formatNumber(result))
	return nil
}
