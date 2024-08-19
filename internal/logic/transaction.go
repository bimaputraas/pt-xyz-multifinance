package logic

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	"sync"
	"time"
	"xyz-multifinance/internal/model"
	"xyz-multifinance/pkg"
)

const (
	BUNGA     float64 = 0.08
	ADMIN_FEE float64 = 30000
)

func (logic *Logic) NewTransaction(userId int, data model.Transaction) error {
	var (
		errs  = []error{}
		wg    = sync.WaitGroup{}
		mutex = sync.Mutex{}
	)
	if err := pkg.ValidateStruct(data); err != nil {
		return ErrInvalidArgument(err)
	}

	userLimit, err := logic.repo.UserLimitRepository.FindByUserId(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound(err)
	}

	if err != nil {
		return ErrInternal(err)
	}

	tx, err := logic.repo.NewTx()
	if err != nil {
		return ErrInternal(err)
	}

	defer func() {
		p := recover()
		if p != nil {
			tx.Rollback()
		}
	}()

	transaction := model.Transaction{
		NomorKontrak:  strconv.Itoa(int(time.Now().UnixNano())),
		OTR:           data.OTR,
		AdminFee:      ADMIN_FEE,
		Tenor:         data.Tenor,
		JumlahCicilan: calcJumlahCicilan(data.Tenor, data.OTR),
		JumlahBunga:   calcJumlahBunga(data.Tenor, data.OTR),
		NamaAsset:     data.NamaAsset,
	}

	wg.Add(2)

	go func() {
		defer wg.Done()
		transaction.UserID = uint(userId)
		transaction.Datetime = time.Now().Format("2006-01-02 15:04:05")
		tx, err = logic.repo.TransactionRepository.CreateWithTx(tx, transaction)
		if err != nil {
			mutex.Lock()
			errs = append(errs, ErrInternal(err))
			mutex.Unlock()
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := updateUserLimitAmount(&userLimit, transaction.JumlahCicilan); err != nil {
			mutex.Lock()
			errs = append(errs, ErrInvalidArgument(err))
			mutex.Unlock()
			return
		}
		userLimit.UserID = uint(userId)
		tx, err = logic.repo.UserLimitRepository.UpdateWithTx(tx, userLimit)
		if err != nil {
			mutex.Lock()
			errs = append(errs, ErrInternal(err))
			mutex.Unlock()
		}
	}()

	wg.Wait()
	if len(errs) > 0 {
		tx.Rollback()
		return ErrInternal(errs[0])
	}

	tx.Commit()
	return nil
}

func calcJumlahBunga(tenor int, otr float64) float64 {
	monthlyBunga := otr / float64(tenor) * (BUNGA)
	return monthlyBunga * float64(tenor)
}

func calcJumlahCicilan(tenor int, otr float64) float64 {
	monthly := otr / float64(tenor) * (1 + BUNGA)
	return monthly * float64(tenor)
}

func updateUserLimitAmount(user *model.UserLimit, jumlahCicilan float64) error {
	limitAmount := user.Tenor4 - jumlahCicilan
	if limitAmount < 0 {
		return errors.New("not enough limit")
	}
	user.Tenor1 = limitAmount * 0.25
	user.Tenor2 = limitAmount * 0.5
	user.Tenor3 = limitAmount * 0.75
	user.Tenor4 = limitAmount
	return nil
}
