package repository

import (
	"fmt"
	"xyz-multifinance/internal/model"

	"gorm.io/gorm"
)

func UserMySQL(db *gorm.DB) UserRepository {
	return &mySqlRepo{
		DB: db,
	}
}
func UserDetailMySQL(db *gorm.DB) UserDetailRepository {
	return &mySqlRepo{
		DB: db,
	}
}
func TransactionMySQL(db *gorm.DB) TransactionRepository {
	return &mySqlRepo{
		DB: db,
	}
}
func UserLimitMySQL(db *gorm.DB) UserLimitRepository {
	return &mySqlRepo{
		DB: db,
	}
}

type mySqlRepo struct {
	*gorm.DB
}

// user

func (r *mySqlRepo) CreateUser(user model.User) error {
	if err := r.DB.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *mySqlRepo) FindAllUsers() ([]model.User, error) {
	var users []model.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to find users: %w", err)
	}
	return users, nil
}

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
func (r *mySqlRepo) FindUserByEmail(email string) (model.User, error, bool) {
	var user model.User
	if err := r.DB.First(&user, email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, nil, false
		}
		return user, fmt.Errorf("failed to find user by id: %w", err), false
	}
	return user, nil, true
}

// user detail

func (r *mySqlRepo) FindUserDetailByUId(userId int) (model.UserDetail, error, bool) {
	var user model.UserDetail
	if err := r.DB.First(&user, model.UserDetail{UserID: uint(userId)}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, nil, false
		}
		return user, fmt.Errorf("failed to find user detail: %w", err), false
	}
	return user, nil, true
}

func (r *mySqlRepo) CreateUserDetail(userDetail model.UserDetail) error {
	if err := r.DB.Create(&userDetail).Error; err != nil {
		return fmt.Errorf("failed to create userDetail: %w", err)
	}
	return nil
}

// transaction

func (r *mySqlRepo) CreateTransaction(transaction model.Transaction) error {
	if err := r.DB.Create(&transaction).Error; err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}
	return nil
}

func (r *mySqlRepo) FindAllTransactions() ([]model.Transaction, error) {
	var transactions []model.Transaction
	if err := r.DB.Find(&transactions).Error; err != nil {
		return nil, fmt.Errorf("failed to find transactions: %w", err)
	}
	return transactions, nil
}

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

func (r *mySqlRepo) UpdateTransactionById(id int, update model.Transaction) (error, bool) {
	if err := r.DB.Model(&model.Transaction{ID: uint(id)}).Updates(update).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false
		}
		return fmt.Errorf("failed to update transaction: %w", err), false
	}
	return nil, true
}

// user limit

func (r *mySqlRepo) CreateUserLimit(userLimit model.UserLimit) error {
	if err := r.DB.Create(&userLimit).Error; err != nil {
		return fmt.Errorf("failed to create userLimit: %w", err)
	}
	return nil
}

func (r *mySqlRepo) FindUserLimitByUId(userId int) (model.UserLimit, error, bool) {
	var userLimit model.UserLimit
	if err := r.DB.First(&userLimit, model.UserLimit{
		UserID: uint(userId),
	}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return userLimit, nil, false
		}
		return userLimit, fmt.Errorf("failed to find userLimit by id: %w", err), false
	}
	return userLimit, nil, true
}

func (r *mySqlRepo) UpdateUserLimitByUId(userId int, update model.UserLimit) (error, bool) {
	if err := r.DB.Model(&model.UserLimit{ID: uint(userId)}).Updates(update).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false
		}
		return fmt.Errorf("failed to update userLimit: %w", err), false
	}
	return nil, true
}
