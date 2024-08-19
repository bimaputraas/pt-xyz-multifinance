package logic

import (
	"errors"
	"xyz-multifinance/internal/config"
	"xyz-multifinance/internal/repository"
)

type (
	Logic struct {
		repo   *repository.Repository
		config *config.Config
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
	Illegal         = 4
)

func New(repo *repository.Repository, config *config.Config) (*Logic, error) {
	return &Logic{
		repo:   repo,
		config: config,
	}, nil
}

func ParseError(err error) (Error, bool) {
	var r *errLogic
	ok := errors.As(err, &r)
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
	return &errLogic{code: Illegal, msg: err.Error()}
}

func (e *errLogic) Code() int {
	return e.code
}
func (e *errLogic) Error() string {
	return e.msg
}
