package calculator

import (
	"errors"
	"fmt"
)

type Operation string

const (
	Add      Operation = "add"
	Subtract Operation = "sub"
	Multiply Operation = "mul"
	Divide   Operation = "div"
)

var ErrDivideByZero = errors.New("division by zero")

func Apply(operation Operation, left float64, right float64) (float64, error) {
	switch operation {
	case Add:
		return left + right, nil
	case Subtract:
		return left - right, nil
	case Multiply:
		return left * right, nil
	case Divide:
		if right == 0 {
			return 0, ErrDivideByZero
		}
		return left / right, nil
	default:
		return 0, fmt.Errorf("unsupported operation %q", operation)
	}
}
