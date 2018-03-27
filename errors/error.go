package errors

import "fmt"

type Error struct {
	message string
	context *Context
}

func (m *Error) Error() string {
	return fmt.Sprintf("%s\n%s", m.message, m.context)
}

func New(format string, v ...interface{}) error {

	return &Error{
		context: NewContext(),
		message: fmt.Sprintf(
			format,
			v...,
		),
	}
}

func By(err error) error {

	switch err.(type) {
	case *Error:
		return err
	default:
		return &Error{
			message: err.Error(),
			context: NewContext(),
		}
	}
}
