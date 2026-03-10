package myerror

import (
	"fmt"
)

type MyError struct {
	Op string
}

func (e *MyError) Error() string {
	if e == nil {
		return "<nil MyError>"
	}
	return fmt.Sprintf("operation failed: %s", e.Op)
}

// BuggyDoThing returns a *MyError(nil) as error (typed nil pitfall)
func BuggyDoThing(fail bool) error {
	if fail {
		return &MyError{Op: "fail"}
	}
	var e *MyError = nil
	return e // BAD: interface is non-nil!
}

// FixedDoThing returns a true nil error when no error
func FixedDoThing(fail bool) error {
	if fail {
		return &MyError{Op: "fail"}
	}
	return nil // GOOD: interface is nil
}

// Wraps error for errors.As demonstration
func WrapError(err error) error {
	return fmt.Errorf("wrapped: %w", err)
}
