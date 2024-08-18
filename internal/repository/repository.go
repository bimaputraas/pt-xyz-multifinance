package repository

import (
	"xyz-multifinance/internal/model"
)

type (
	Repository struct {
		UserRepository
		UserDetailRepository
		TransactionRepository
		UserLimitRepository
	}

	UserRepository interface {
		FindUserById(id int) (model.User, error, bool)
		FindAllUsers() ([]model.User, error)
		CreateUser(model.User) error
		FindUserByEmail(email string) (model.User, error, bool)
	}
	UserDetailRepository interface {
		FindUserDetailByUId(userId int) (model.UserDetail, error, bool)
		CreateUserDetail(model.UserDetail) error
	}

	UserLimitRepository interface {
		FindUserLimitByUId(userId int) (model.UserLimit, error, bool)
		CreateUserLimit(model.UserLimit) error
		UpdateUserLimitByUId(userId int, update model.UserLimit) (error, bool)
	}
	TransactionRepository interface {
		FindTransactionById(id int) (model.Transaction, error, bool)
		FindAllTransactions() ([]model.Transaction, error)
		CreateTransaction(model.UserLimit, model.Transaction) error
		UpdateTransactionById(id int, update model.Transaction) (error, bool)
	}
)

func New(UserRepository, UserDetailRepository, TransactionRepository, UserLimitRepository) (*Repository, error) {
	return &Repository{}, nil
}
