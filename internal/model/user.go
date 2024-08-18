package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	User struct {
		ID        uint           `gorm:"primaryKey" json:"id,omitempty"`
		CreatedAt time.Time      `json:"created_at,omitempty"`
		UpdatedAt time.Time      `json:"updated_at,omitempty"`
		DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
		Email     string         `gorm:"unique;not null" json:"email,omitempty" validate:"required,email"`
		Password  string         `gorm:"not null" json:"password,omitempty" validate:"required"`
	}

	UserDetail struct {
		ID           uint           `gorm:"primaryKey" json:"id,omitempty"`
		CreatedAt    time.Time      `json:"created_at,omitempty"`
		UpdatedAt    time.Time      `json:"updated_at,omitempty"`
		DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
		NIK          string         `gorm:"unique;not null" json:"nik" validate:"required"`
		FullName     string         `gorm:"not null" json:"full_name" validate:"required"`
		LegalName    string         `gorm:"not null" json:"legal_name" validate:"required"`
		TempatLahir  string         `gorm:"not null" json:"tempat_lahir" validate:"required"`
		TanggalLahir time.Time      `gorm:"not null" json:"tanggal_lahir" validate:"required"`
		Gaji         float64        `gorm:"not null" json:"gaji" validate:"required"`
		FotoKTP      string         `gorm:"not null" json:"foto_ktp" validate:"required"`
		FotoSelfie   string         `gorm:"not null" json:"foto_selfie" validate:"required"`
		User         User           `gorm:"references:ID"`
		UserID       uint           `gorm:"not null" json:"user_id"`
	}

	UserLimit struct {
		ID        uint           `gorm:"primaryKey" json:"id,omitempty"`
		UserID    uint           `gorm:"not null" json:"konsumen_id"`
		Tenor1    float64        `json:"tenor_1" validate:"required"`
		Tenor2    float64        `json:"tenor_2" validate:"required"`
		Tenor3    float64        `json:"tenor_3" validate:"required"`
		Tenor4    float64        `json:"tenor_4" validate:"required"`
		CreatedAt time.Time      `json:"created_at,omitempty"`
		UpdatedAt time.Time      `json:"updated_at,omitempty"`
		DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

		UserDetail UserDetail `gorm:"foreignKey:UserID" json:"konsumen,omitempty"`
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
