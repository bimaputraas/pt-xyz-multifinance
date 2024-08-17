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
	InvalidArgument = 1
	NotFound        = 2
	Internal        = 3
	Ilegal          = 4
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

func ErrInvalidArgument(err error) error {
	return &errLogic{code: InvalidArgument, msg: err.Error()}
}

func ErrNotFound(err error) error {
	return &errLogic{code: NotFound, msg: err.Error()}
}

func ErrInternal(err error) error {
	return &errLogic{code: Internal, msg: err.Error()}
}

func ErrIllegal(err error) error {
	return &errLogic{code: Internal, msg: err.Error()}
}

func (e *errLogic) Code() int {
	return e.code
}
func (e *errLogic) Error() string {
	return e.msg
}
