package logic

import (
	"errors"
	"xyz-multifinance/internal/model"
	"xyz-multifinance/pkg"
)

func (logic Logic) RegisterInstallment(userId int, userDetail model.UserDetail) error {
	if err := pkg.ValidateStruct(userDetail); err != nil {
		return ErrInvalidArgument(err)
	}

	userDetail.UserID = uint(userId)
	userDetail.IsVerified = true
	err := logic.Repository.CreateUserDetail(userDetail)
	if err != nil {
		return ErrInternal(errors.New("failed update userDetail"))
	}

	return nil
}
