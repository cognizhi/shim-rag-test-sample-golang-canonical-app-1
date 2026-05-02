package calculator

import (
	"errors"
	"testing"
)

func TestApply(t *testing.T) {
	tests := []struct {
		name      string
		operation Operation
		left      float64
		right     float64
		want      float64
	}{
		{name: "add", operation: Add, left: 7, right: 5, want: 12},
		{name: "subtract", operation: Subtract, left: 7, right: 5, want: 2},
		{name: "multiply", operation: Multiply, left: 7, right: 5, want: 35},
		{name: "divide", operation: Divide, left: 10, right: 2, want: 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Apply(test.operation, test.left, test.right)
			if err != nil {
				t.Fatalf("Apply() error = %v", err)
			}
			if got != test.want {
				t.Fatalf("Apply() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestApplyDivideByZero(t *testing.T) {
	_, err := Apply(Divide, 10, 0)
	if !errors.Is(err, ErrDivideByZero) {
		t.Fatalf("Apply() error = %v, want %v", err, ErrDivideByZero)
	}
}
