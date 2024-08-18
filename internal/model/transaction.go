package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	Transaction struct {
		ID            uint           `gorm:"primaryKey" json:"id,omitempty"`
		CreatedAt     time.Time      `json:"created_at,omitempty"`
		UpdatedAt     time.Time      `json:"updated_at,omitempty"`
		DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
		NomorKontrak  string         `gorm:"unique;not null" json:"nomor_kontrak" validate:"required"`
		OTR           float64        `gorm:"not null" json:"otr" validate:"required"`
		AdminFee      float64        `gorm:"not null" json:"admin_fee" validate:"required"`
		JumlahCicilan float64        `gorm:"not null" json:"jumlah_cicilan" validate:"required"`
		JumlahBunga   float64        `gorm:"not null" json:"jumlah_bunga" validate:"required"`
		NamaAsset     string         `gorm:"not null" json:"nama_asset" validate:"required"`
		UserDetail    UserDetail     `gorm:"references:ID"`
		UserID        uint           `gorm:"not null" json:"userDetail_id"`
	}
)
