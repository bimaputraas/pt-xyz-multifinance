package logic

import (
	"errors"
	"xyz-multifinance/internal/model"
)

func (logic Logic) NewTransaction(userId int, transaction model.Transaction) error {
	userLimit, err, ok := logic.Repository.FindUserLimitByUId(userId)
	if err != nil {
		return ErrInternal(err)
	}

	if !ok {
		return ErrNotFound(errors.New("not found"))
	}

	logic.Repository.CreateTransaction(userLimit, transaction)
	return nil
}
