package util

import "fmt"

type Error struct {
	Message string
}

func (e Error) Error() string {
	return e.Message
}

func EmailUsed() Error {
	return Error{
		Message: "email has been userd",
	}
}

func NotFound() Error {
	return Error{
		Message: "Data not found",
	}
}

func Errors(err error) Error {
	return Error{
		Message: fmt.Sprintf("something error: %s", err.Error()),
	}
}