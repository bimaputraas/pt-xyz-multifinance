package logic

import (
	"xyz-multifinance/internal/repository"
)

type (
	Logic struct {
		repository.Repository
	}
	Error interface {
		Code() int
		error
	}

	errLogic struct {
		code int
		msg  string
	}
)

const (
	ErrInvalidArgument = 1
	ErrNotFound        = 2
	ErrInternal        = 3
)

func New() (*Logic, error) {
	return &Logic{}, nil
}

func ParseError(err error) (Error, bool) {
	r, ok := err.(*errLogic)
	if !ok {
		return &errLogic{}, false
	}

	return r, true
}

func InvalidArgument(msg string) error {
	return &errLogic{code: ErrInvalidArgument, msg: msg}
}

func NotFound(msg string) error {
	return &errLogic{code: ErrNotFound, msg: msg}
}

func Internal(msg string) error {
	return &errLogic{code: ErrInternal, msg: msg}
}

func (e *errLogic) Code() int {
	return e.code
}
func (e *errLogic) Error() string {
	return e.msg
}
