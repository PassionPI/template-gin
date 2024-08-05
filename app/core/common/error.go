package common

import "fmt"

type Error struct {
	Test error
}

func NewError() *Error {
	return &Error{
		Test: fmt.Errorf("test"),
	}
}
