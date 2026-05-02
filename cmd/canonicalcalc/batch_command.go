package main

import (
	"fmt"
	"os"

	"github.com/cognizhi/ragtest/canonical-calculator/internal/batch"
)

func init() {
	registerCommand(command{
		name:    "batch",
		summary: "Run calculator operations from a CSV file",
		usage:   "batch <csv-path>",
		run:     runBatchCommand,
	})
}

func runBatchCommand(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("batch expects exactly one CSV path")
	}

	file, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer file.Close()

	results, err := batch.RunCSV(file)
	if err != nil {
		return err
	}

	for _, result := range results {
		fmt.Printf("line %d: %s(%s, %s) = %s\n", result.Line, result.Operation, formatNumber(result.Left), formatNumber(result.Right), formatNumber(result.Result))
	}
	return nil
}
