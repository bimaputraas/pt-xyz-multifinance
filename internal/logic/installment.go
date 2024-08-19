package logic

import (
	"errors"
	"time"
	"xyz-multifinance/internal/model"
	"xyz-multifinance/pkg"
)

func (logic *Logic) RegisterInstallment(userId int, userDetail model.UserDetail) error {
	if err := pkg.ValidateStruct(userDetail); err != nil {
		return ErrInvalidArgument(err)
	}
	now := time.Now().Format("2006-01-02 15:04:05")

	if userDetail.Gaji < 500000 {
		return ErrInvalidArgument(errors.New("gaji below required amount"))
	}
	tx, err := logic.repo.NewTx()
	if err != nil {
		return err
	}

	userDetail.UserID = uint(userId)
	userDetail.Datetime = now
	tx, err = logic.repo.UserDetailRepository.CreateWithTx(tx, userDetail)
	if err != nil {
		tx.Rollback()
		return ErrInternal(err)
	}

	userLimit := model.UserLimit{}
	userLimit.UserID = uint(userId)
	userLimit.DefaultTenors(userDetail.Gaji)
	userLimit.Datetime = now
	tx, err = logic.repo.UserLimitRepository.CreateWithTx(tx, userLimit)
	if err != nil {
		tx.Rollback()
		return ErrInternal(err)
	}

	tx.Commit()

	return nil
}
