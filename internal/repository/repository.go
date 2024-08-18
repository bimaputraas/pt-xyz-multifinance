package repository

import (
	"errors"
	"gorm.io/gorm"
	"xyz-multifinance/internal/model"
)

func NewMYSQL(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:        &userMySql{db},
		UserDetailRepository:  &userDetailMySql{db},
		UserLimitRepository:   &userLimitMySql{db},
		TransactionRepository: &transactionMySql{},
		db:                    db,
	}
}

const (
	Internal = 1
	NotFound = 2
)

type (
	Error interface {
		Code() int
		error
	}
	errRepo struct {
		code int
		msg  string
	}

	Repository struct {
		UserRepository
		UserDetailRepository
		UserLimitRepository
		TransactionRepository
		db any
	}
	TransactionRepository interface {
		CreateWithTx(DBTx, model.Transaction) (DBTx, error)
	}

	UserLimitRepository interface {
		CreateWithTx(DBTx, model.UserLimit) (DBTx, error)
		FindByUserId(int) (model.UserLimit, error)
		UpdateWithTx(DBTx, model.UserLimit) (DBTx, error)
	}

	UserDetailRepository interface {
		CreateWithTx(DBTx, model.UserDetail) (DBTx, error)
		FindByUserId(int) (model.UserDetail, error)
	}

	UserRepository interface {
		Create(model.User) error
		FindById(id int) (model.User, error)
		FindByEmail(email string) (model.User, error)
	}
)

func (r *Repository) NewTx() (DBTx, error) {
	db, ok := r.db.(*gorm.DB)
	if !ok {
		return nil, errors.New("failed assert")
	}
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dbTx{tx: tx}, nil
}

func ParseError(err error) (Error, bool) {
	r, ok := err.(*errRepo)
	if !ok {
		return &errRepo{}, false
	}

	return r, true
}
func ErrNotFound(err error) error {
	return &errRepo{code: NotFound, msg: err.Error()}
}
func (e *errRepo) Code() int {
	return e.code
}
func (e *errRepo) Error() string {
	return e.msg
}

type DBTx interface {
	Rollback()
	Commit()
	gorm() *gorm.DB
}

func (d *dbTx) Rollback() {
	d.tx.Rollback()
}

func (d *dbTx) Commit() {
	d.tx.Commit()
}

func (d *dbTx) gorm() *gorm.DB {
	return d.tx
}

type dbTx struct {
	tx *gorm.DB
}
