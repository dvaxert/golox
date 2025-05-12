package loxerr

import "fmt"

type LoxError struct {
	line    int
	where   string
	message string
}

//----------------------------------------------------------------------------------------------------------------------

type LoxErrorOpt func(*LoxError)

//----------------------------------------------------------------------------------------------------------------------

func New(opts ...LoxErrorOpt) *LoxError {
	err := &LoxError{}

	for _, opt := range opts {
		opt(err)
	}

	return err
}

//----------------------------------------------------------------------------------------------------------------------

func WithMessage(message string) LoxErrorOpt {
	return func(e *LoxError) {
		e.message = message
	}
}

//----------------------------------------------------------------------------------------------------------------------

func WithLine(line int) LoxErrorOpt {
	return func(e *LoxError) {
		e.line = line
	}
}

//----------------------------------------------------------------------------------------------------------------------

func WithWhere(where string) LoxErrorOpt {
	return func(e *LoxError) {
		e.where = where
	}
}

//----------------------------------------------------------------------------------------------------------------------

func (e *LoxError) Error() string {
	return fmt.Sprintf("[line %d] Error%s: %s", e.line, e.where, e.message)
}
