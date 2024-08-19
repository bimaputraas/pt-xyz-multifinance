package model

import (
	"time"
)

type (
	User struct {
		ID       uint   `gorm:"column:id;primaryKey" json:"id,omitempty"`
		Datetime string `gorm:"column:datetime" json:"datetime,omitempty"`
		Email    string `gorm:"column:email;unique;not null" json:"email,omitempty" validate:"required,email"`
		Password string `gorm:"column:password;not null" json:"password,omitempty" validate:"required"`
	}

	UserDetail struct {
		ID           uint      `gorm:"column:id;primaryKey" json:"id,omitempty"`
		Datetime     string    `gorm:"column:datetime" json:"datetime,omitempty"`
		NIK          string    `gorm:"column:nik;unique;not null" json:"nik" validate:"required"`
		FullName     string    `gorm:"column:full_name;not null" json:"full_name" validate:"required"`
		LegalName    string    `gorm:"column:legal_name;not null" json:"legal_name" validate:"required"`
		TempatLahir  string    `gorm:"column:tempat_lahir;not null" json:"tempat_lahir" validate:"required"`
		TanggalLahir time.Time `gorm:"column:tanggal_lahir;not null" json:"tanggal_lahir" validate:"required"`
		Gaji         float64   `gorm:"column:gaji;not null" json:"gaji" validate:"required"`
		FotoKTP      string    `gorm:"column:foto_ktp;not null" json:"foto_ktp" validate:"required"`
		FotoSelfie   string    `gorm:"column:foto_selfie;not null" json:"foto_selfie" validate:"required"`
		UserID       uint      `gorm:"column:user_id;not null" json:"user_id"`
	}

	UserLimit struct {
		ID       uint    `gorm:"column:id;primaryKey" json:"id,omitempty"`
		UserID   uint    `gorm:"column:user_id;not null" json:"user_id"`
		Tenor1   float64 `gorm:"column:tenor_1" json:"tenor_1" validate:"required"`
		Tenor2   float64 `gorm:"column:tenor_2" json:"tenor_2" validate:"required"`
		Tenor3   float64 `gorm:"column:tenor_3" json:"tenor_3" validate:"required"`
		Tenor4   float64 `gorm:"column:tenor_4" json:"tenor_4" validate:"required"`
		Datetime string  `gorm:"column:datetime" json:"datetime,omitempty"`
	}
)

func (u *UserLimit) DefaultTenors(gaji float64) {
	base := gaji * 0.4

	u.Tenor1 = base
	u.Tenor2 = base * 2
	u.Tenor3 = base * 3
	u.Tenor4 = base * 4
}

func (u *User) Reset() {
	u = &User{}
}
