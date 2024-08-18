package logic

import (
	"errors"
	"xyz-multifinance/internal/model"
	"xyz-multifinance/pkg"
)

func (logic *Logic) RegisterInstallment(userId int, userDetail model.UserDetail) error {
	if err := pkg.ValidateStruct(userDetail); err != nil {
		return ErrInvalidArgument(err)
	}

	if userDetail.Gaji < 500000 {
		return ErrInvalidArgument(errors.New("gaji below required amount"))
	}
	userDetail.UserID = uint(userId)
	userDetail.IsVerified = true
	
	tx, err := logic.repo.NewTx()
	if err != nil {
		return err
	}

	tx, err = logic.repo.UserDetailRepository.CreateWithTx(tx, userDetail)
	if err != nil {
		tx.Rollback()
		return ErrInternal(err)
	}

	userLimit := model.UserLimit{}
	userLimit.DefaultTenors(userDetail.Gaji)
	tx, err = logic.repo.UserLimitRepository.CreateWithTx(tx, userLimit)
	if err != nil {
		tx.Rollback()
		return ErrInternal(err)
	}

	tx.Commit()

	return nil
}
