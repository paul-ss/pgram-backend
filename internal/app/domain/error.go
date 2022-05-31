package domain

import (
	"fmt"
)

const (
	ErrTypeDefault = "default"
	ErrType
)

func NewBaseError(err error, errType string) ErrorBase {
	return ErrorBase{
		Type:    errType,
		Message: err.Error(),
		Err:     err,
	}
}

type ErrorBase struct {
	Type    string
	Message string
	Err     error `json:"-"`
}

func (e *ErrorBase) Error() string {
	return fmt.Sprintf("type=%s: %s", e.Type, e.Err.Error())
}

type ErrorNotFound struct {
	ErrorBase
}

type ErrorNotAuthorised struct {
	ErrorBase
}

type ErrorBadRequest struct {
	ErrorBase
}
