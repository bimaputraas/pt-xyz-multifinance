package repository

import (
	"errors"
	"fmt"
	"xyz-multifinance/internal/model"

	"gorm.io/gorm"
)

type (
	userMySql struct {
		*gorm.DB
	}
	userDetailMySql struct {
		*gorm.DB
	}
	transactionMySql struct {
		*gorm.DB
	}
	userLimitMySql struct {
		*gorm.DB
	}
)

// user

func (r *userMySql) Create(data model.User) error {
	if err := r.DB.Create(&data).Error; err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *userMySql) FindById(id int) (model.User, error) {
	var user model.User
	if err := r.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, ErrNotFound(err)
		}
		return user, err
	}
	return user, nil
}
func (r *userMySql) FindByEmail(email string) (model.User, error) {
	var user model.User
	if err := r.DB.First(&user, email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, ErrNotFound(err)
		}
		return user, err
	}
	return user, nil
}

// user detail

func (r *userDetailMySql) FindByUserId(userId int) (model.UserDetail, error) {
	var user model.UserDetail
	if err := r.DB.First(&user, model.UserDetail{UserID: uint(userId)}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, ErrNotFound(err)
		}
		return user, err
	}
	return user, nil
}

func (r *userDetailMySql) CreateWithTx(dbTx DBTx, data model.UserDetail) (DBTx, error) {
	err := dbTx.gorm().Create(&data).Error
	return dbTx, err
}

// transaction

func (r *transactionMySql) CreateWithTx(dbTx DBTx, data model.Transaction) (DBTx, error) {
	err := dbTx.gorm().Create(&data).Error
	return dbTx, err
}

// user limit

func (r *userLimitMySql) CreateWithTx(dbTx DBTx, data model.UserLimit) (DBTx, error) {
	err := dbTx.gorm().Create(&data).Error
	return dbTx, err
}

func (r *userLimitMySql) FindByUserId(userId int) (model.UserLimit, error) {
	var userLimit model.UserLimit
	if err := r.DB.First(&userLimit, model.UserLimit{
		UserID: uint(userId),
	}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userLimit, ErrNotFound(err)
		}
		return userLimit, err
	}
	return userLimit, nil
}

func (r *userLimitMySql) UpdateWithTx(dbTx DBTx, data model.UserLimit) (DBTx, error) {
	if err := dbTx.gorm().Model(&model.UserLimit{ID: data.UserID}).Updates(data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dbTx, ErrNotFound(err)
		}
		return dbTx, err
	}
	return dbTx, nil
}
