package main

import (
	"fmt"
	"math"
	"strconv"
)

func parseFloatArgument(value string) (float64, error) {
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("parse %q as number: %w", value, err)
	}
	return parsed, nil
}

func formatNumber(value float64) string {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return strconv.FormatFloat(value, 'g', -1, 64)
	}

	if math.Abs(value-math.Round(value)) < 0.000000001 {
		return strconv.FormatFloat(math.Round(value), 'f', 0, 64)
	}

	return strconv.FormatFloat(value, 'f', -1, 64)
}
