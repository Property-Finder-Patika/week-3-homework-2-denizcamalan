package myerror

import "fmt"

type NegativeArgumentError struct {
	message string
	arg     float64
}

func (e *NegativeArgumentError) Error() string {
	s := fmt.Sprintf("%s: %f", "argument can't be negative", e.arg)
	return s
}

func NewNegativeArgumentError(text string, arg float64) *NegativeArgumentError {
	return &NegativeArgumentError{text, arg}
}
