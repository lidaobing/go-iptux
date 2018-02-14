package errors

import (
	"fmt"
)

type ErrorCode int32

const (
	NotImplementd ErrorCode = iota
	Wrap
)

type Error struct {
	code ErrorCode
	message string
	causedBy error
}

func (e *Error) Error() string {
	if e.causedBy == nil {
		return fmt.Sprintf("libiptux.Error: [%d] %s", e.code, e.message);
	} else {
		return fmt.Sprintf("libiptux.Error: [%d] %s\n  Caused By: %s", e.code, e.message, e.causedBy.Error())
	}
}

func NewError(code ErrorCode, message string) *Error{
	return &Error{code, message, nil}
}

func NewErrorWithCausedBy(code ErrorCode, message string, causedBy error) *Error {
	return &Error{code, message, causedBy}
}

func NewNotImplementedError() *Error {
	return NewError(NotImplementd, "not implemented");
}

func WrapError(err error) *Error {
	_, ok := err.(*Error)
	if ok {
		return err.(*Error)
	}

	return NewErrorWithCausedBy(Wrap, "wrap error", err)
}


