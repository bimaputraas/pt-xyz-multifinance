package logic

import (
	"xyz-multifinance/internal/model"
	"xyz-multifinance/pkg"
)

func (logic Logic) Register(user model.User) error {
	if err := pkg.ValidateStruct(user); err != nil {
		return ErrInvalidArgument(err)
	}

	hashed, err := pkg.Hash(user.Password)
	if err != nil {
		return ErrInternal(err)
	}

	user.Password = hashed

	if err := logic.Repository.CreateUser(user); err != nil {
		return ErrInternal(err)
	}

	return nil
}
