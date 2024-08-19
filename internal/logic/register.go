package logic

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"xyz-multifinance/internal/model"
	"xyz-multifinance/pkg"
)

func (logic *Logic) Register(user model.User) error {
	if err := pkg.ValidateStruct(user); err != nil {
		return ErrInvalidArgument(err)
	}

	hashed, err := pkg.Hash(user.Password)
	if err != nil {
		return ErrInternal(err)
	}

	user.Password = hashed

	check, err := logic.repo.UserRepository.FindByEmail(user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return ErrInternal(err)
	}

	if check.Email == user.Email {
		return ErrInvalidArgument(errors.New("email already registered"))
	}

	user.Datetime = time.Now().Format("2006-01-02 15:04:05")
	if err := logic.repo.UserRepository.Create(user); err != nil {
		return ErrInternal(err)
	}

	return nil
}
