package main

import (
	"fmt"
	"strings"

	"github.com/cognizhi/ragtest/canonical-calculator/internal/units"
)

func init() {
	registerCommand(command{
		name:    "convert-length",
		summary: "Convert length values between supported units",
		usage:   "convert-length <value> <from> <to>",
		run:     runConvertLengthCommand,
	})
}

func runConvertLengthCommand(args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("convert-length expects value, source unit, and target unit; supported units: %s", strings.Join(units.SupportedLengthUnits(), ", "))
	}

	value, err := parseFloatArgument(args[0])
	if err != nil {
		return err
	}

	converted, err := units.ConvertLength(value, args[1], args[2])
	if err != nil {
		return err
	}

	fmt.Printf("%s %s = %s %s\n", formatNumber(value), strings.ToLower(args[1]), formatNumber(converted), strings.ToLower(args[2]))
	return nil
}
