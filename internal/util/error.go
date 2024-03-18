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
		Message: "email has been used",
	}
}

func DataHasBeenUsed() Error {
	return Error{
		Message: "data by this name has been used",
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

func Validation(msg string) Error {
	return Error{
		Message: msg,
	}
}
