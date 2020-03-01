package http_lib

import "fmt"

// errorString is a trivial implementation of error.
type ErrorString struct {
	Reason	string
	Code	int
}

func (e *ErrorString) Error() string {
	return fmt.Sprintf("%s: %d", e.Reason, e.Code)
}
