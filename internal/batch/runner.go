package batch

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/cognizhi/ragtest/canonical-calculator/internal/calculator"
)

type Result struct {
	Line      int
	Operation calculator.Operation
	Left      float64
	Right     float64
	Result    float64
}

func RunCSV(reader io.Reader) ([]Result, error) {
	csvReader := csv.NewReader(reader)
	csvReader.TrimLeadingSpace = true

	header, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("read header: %w", err)
	}
	if err := validateHeader(header); err != nil {
		return nil, err
	}

	var results []Result
	lineNumber := 1
	for {
		lineNumber++
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("read line %d: %w", lineNumber, err)
		}

		result, err := runRecord(lineNumber, record)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func validateHeader(header []string) error {
	expected := []string{"operation", "left", "right"}
	if len(header) != len(expected) {
		return fmt.Errorf("header must be operation,left,right")
	}

	for index, expectedColumn := range expected {
		if strings.ToLower(strings.TrimSpace(header[index])) != expectedColumn {
			return fmt.Errorf("header column %d must be %q", index+1, expectedColumn)
		}
	}

	return nil
}

func runRecord(lineNumber int, record []string) (Result, error) {
	if len(record) != 3 {
		return Result{}, fmt.Errorf("line %d must contain operation,left,right", lineNumber)
	}

	operation, err := parseOperation(record[0])
	if err != nil {
		return Result{}, fmt.Errorf("line %d: %w", lineNumber, err)
	}

	left, err := strconv.ParseFloat(strings.TrimSpace(record[1]), 64)
	if err != nil {
		return Result{}, fmt.Errorf("line %d: parse left operand: %w", lineNumber, err)
	}

	right, err := strconv.ParseFloat(strings.TrimSpace(record[2]), 64)
	if err != nil {
		return Result{}, fmt.Errorf("line %d: parse right operand: %w", lineNumber, err)
	}

	computed, err := calculator.Apply(operation, left, right)
	if err != nil {
		return Result{}, fmt.Errorf("line %d: %w", lineNumber, err)
	}

	return Result{
		Line:      lineNumber,
		Operation: operation,
		Left:      left,
		Right:     right,
		Result:    computed,
	}, nil
}

func parseOperation(value string) (calculator.Operation, error) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "add":
		return calculator.Add, nil
	case "sub":
		return calculator.Subtract, nil
	case "mul":
		return calculator.Multiply, nil
	case "div":
		return calculator.Divide, nil
	default:
		return "", fmt.Errorf("unsupported operation %q", value)
	}
}
