package repository

import (
	"xyz-multifinance/internal/model"
)

type (
	Repository struct {
		CustomerRepository
		UserRepository
		TransactionRepository
		CustomerLimitRepository
	}

	CustomerRepository interface {
		FindCustomerById(id int) (model.Customer, error, bool)
		FindAllCustomers() ([]model.Customer, error)
		CreateCustomer(model.Customer) error
		UpdateCustomerById(id int, update model.Customer) (error, bool)
	}
	UserRepository interface {
		FindUserById(id int) (model.User, error, bool)
		FindAllUsers() ([]model.User, error)
		CreateUser(model.User) error
		UpdateUserById(id int, update model.User) (error, bool)
	}
	TransactionRepository interface {
		FindTransactionById(id int) (model.Transaction, error, bool)
		FindAllTransactions() ([]model.Transaction, error)
		CreateTransaction(model.Transaction) error
		UpdateTransactionById(id int, update model.Transaction) (error, bool)
	}
	CustomerLimitRepository interface {
		FindCustomerLimitById(id int) (model.CustomerLimit, error, bool)
		FindAllCustomerLimits() ([]model.CustomerLimit, error)
		CreateCustomerLimit(model.CustomerLimit) error
		UpdateCustomerLimitById(id int, update model.CustomerLimit) (error, bool)
	}
)

func New(UserRepository, CustomerRepository, TransactionRepository, CustomerLimitRepository) (*Repository, error) {
	return &Repository{}, nil
}
