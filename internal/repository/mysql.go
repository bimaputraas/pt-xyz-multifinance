package repository

import (
	"fmt"
	"xyz-multifinance/internal/model"

	"gorm.io/gorm"
)

func CustomerMySQL(db *gorm.DB) CustomerRepository {
	return &mySqlRepo{
		DB: db,
	}
}
func UserMySQL(db *gorm.DB) UserRepository {
	return &mySqlRepo{
		DB: db,
	}
}
func TransactionMySQL(db *gorm.DB) TransactionRepository {
	return &mySqlRepo{
		DB: db,
	}
}
func CustomerLimitMySQL(db *gorm.DB) CustomerLimitRepository {
	return &mySqlRepo{
		DB: db,
	}
}

type mySqlRepo struct {
	*gorm.DB
}

// CreateUser creates a new user in the database using GORM.
func (r *mySqlRepo) CreateUser(user model.User) error {
	if err := r.DB.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// FindAllUsers retrieves all users from the database using GORM.
func (r *mySqlRepo) FindAllUsers() ([]model.User, error) {
	var users []model.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to find users: %w", err)
	}
	return users, nil
}

// FindUserById retrieves a user by ID from the database using GORM.
func (r *mySqlRepo) FindUserById(id int) (model.User, error, bool) {
	var user model.User
	if err := r.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, nil, false
		}
		return user, fmt.Errorf("failed to find user by id: %w", err), false
	}
	return user, nil, true
}

// UpdateUserById updates a user by ID in the database using GORM.
func (r *mySqlRepo) UpdateUserById(id int, update model.User) (error, bool) {
	// Ensure we only update the fields that should be updated
	if err := r.DB.Model(&model.User{ID: uint(id)}).Updates(update).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false
		}
		return fmt.Errorf("failed to update user: %w", err), false
	}
	return nil, true
}

// CreateCustomer creates a new customer in the database using GORM.
func (r *mySqlRepo) CreateCustomer(customer model.Customer) error {
	if err := r.DB.Create(&customer).Error; err != nil {
		return fmt.Errorf("failed to create customer: %w", err)
	}
	return nil
}

// FindAllCustomers retrieves all customers from the database using GORM.
func (r *mySqlRepo) FindAllCustomers() ([]model.Customer, error) {
	var customers []model.Customer
	if err := r.DB.Find(&customers).Error; err != nil {
		return nil, fmt.Errorf("failed to find customers: %w", err)
	}
	return customers, nil
}

// FindCustomerById retrieves a customer by ID from the database using GORM.
func (r *mySqlRepo) FindCustomerById(id int) (model.Customer, error, bool) {
	var customer model.Customer
	if err := r.DB.First(&customer, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return customer, nil, false
		}
		return customer, fmt.Errorf("failed to find customer by id: %w", err), false
	}
	return customer, nil, true
}

// UpdateCustomerById updates a customer by ID in the database using GORM.
func (r *mySqlRepo) UpdateCustomerById(id int, update model.Customer) (error, bool) {
	// Ensure we only update the fields that should be updated
	if err := r.DB.Model(&model.Customer{ID: uint(id)}).Updates(update).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false
		}
		return fmt.Errorf("failed to update customer: %w", err), false
	}
	return nil, true
}

// CreateTransaction creates a new transaction in the database using GORM.
func (r *mySqlRepo) CreateTransaction(transaction model.Transaction) error {
	if err := r.DB.Create(&transaction).Error; err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}
	return nil
}

// FindAllTransactions retrieves all transactions from the database using GORM.
func (r *mySqlRepo) FindAllTransactions() ([]model.Transaction, error) {
	var transactions []model.Transaction
	if err := r.DB.Find(&transactions).Error; err != nil {
		return nil, fmt.Errorf("failed to find transactions: %w", err)
	}
	return transactions, nil
}

// FindTransactionById retrieves a transaction by ID from the database using GORM.
func (r *mySqlRepo) FindTransactionById(id int) (model.Transaction, error, bool) {
	var transaction model.Transaction
	if err := r.DB.First(&transaction, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return transaction, nil, false
		}
		return transaction, fmt.Errorf("failed to find transaction by id: %w", err), false
	}
	return transaction, nil, true
}

// UpdateTransactionById updates a transaction by ID in the database using GORM.
func (r *mySqlRepo) UpdateTransactionById(id int, update model.Transaction) (error, bool) {
	// Ensure we only update the fields that should be updated
	if err := r.DB.Model(&model.Transaction{ID: uint(id)}).Updates(update).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false
		}
		return fmt.Errorf("failed to update transaction: %w", err), false
	}
	return nil, true
}

// CreateCustomerLimit creates a new customerLimit in the database using GORM.
func (r *mySqlRepo) CreateCustomerLimit(customerLimit model.CustomerLimit) error {
	if err := r.DB.Create(&customerLimit).Error; err != nil {
		return fmt.Errorf("failed to create customerLimit: %w", err)
	}
	return nil
}

// FindAllCustomerLimits retrieves all customerLimits from the database using GORM.
func (r *mySqlRepo) FindAllCustomerLimits() ([]model.CustomerLimit, error) {
	var customerLimits []model.CustomerLimit
	if err := r.DB.Find(&customerLimits).Error; err != nil {
		return nil, fmt.Errorf("failed to find customerLimits: %w", err)
	}
	return customerLimits, nil
}

// FindCustomerLimitById retrieves a customerLimit by ID from the database using GORM.
func (r *mySqlRepo) FindCustomerLimitById(id int) (model.CustomerLimit, error, bool) {
	var customerLimit model.CustomerLimit
	if err := r.DB.First(&customerLimit, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return customerLimit, nil, false
		}
		return customerLimit, fmt.Errorf("failed to find customerLimit by id: %w", err), false
	}
	return customerLimit, nil, true
}

// UpdateCustomerLimitById updates a customerLimit by ID in the database using GORM.
func (r *mySqlRepo) UpdateCustomerLimitById(id int, update model.CustomerLimit) (error, bool) {
	// Ensure we only update the fields that should be updated
	if err := r.DB.Model(&model.CustomerLimit{ID: uint(id)}).Updates(update).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false
		}
		return fmt.Errorf("failed to update customerLimit: %w", err), false
	}
	return nil, true
}
