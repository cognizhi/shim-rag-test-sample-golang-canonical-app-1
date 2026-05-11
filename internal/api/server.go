package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/cognizhi/ragtest/canonical-calculator/internal/calculator"
)


type calculateResponse struct {
	Operation string  `json:"operation"`
	Left      float64 `json:"left"`
	Right     float64 `json:"right"`
	Result    float64 `json:"result"`
}


func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", handleHealth)
	mux.HandleFunc("/calculate", handleCalculate)
	return mux
}

func Run(address string) error {
	return http.ListenAndServe(address, NewHandler())
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte("ok\n"))
}

func handleCalculate(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(writer, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	operation, err := parseOperation(request.URL.Query().Get("op"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	left, err := parseQueryNumber(request, "left")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	right, err := parseQueryNumber(request, "right")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := calculator.Apply(operation, left, right)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(calculateResponse{
		Operation: string(operation),
		Left:      left,
		Right:     right,
		Result:    result,
	})
}

func parseQueryNumber(request *http.Request, name string) (float64, error) {
	value := request.URL.Query().Get(name)
	if value == "" {
		return 0, fmt.Errorf("missing %q query parameter", name)
	}

	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("parse %q query parameter: %w", name, err)
	}

	return parsed, nil
}

func parseOperation(value string) (calculator.Operation, error) {
	switch strings.ToLower(value) {
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
