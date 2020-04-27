package es

import "fmt"

type ErrCommandExecution struct {
	Command CommandMessage
	Reason  string
}

func (e ErrCommandExecution) Error() string {
	return fmt.Sprintf("Invalid operation. Command: %s Reason: %s", e.Command.Type(), e.Reason)
}

type ErrConcurrency struct {
	Expected int
	Given    int
}

func (e ErrConcurrency) Error() string {
	return fmt.Sprintf("Concurrency issue, expected event version of %d, %d given", e.Expected, e.Given)
}
