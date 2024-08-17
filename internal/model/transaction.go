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
		NomorKontrak  string         `gorm:"unique;not null" json:"nomor_kontrak"`
		OTR           float64        `gorm:"not null" json:"otr"`
		AdminFee      float64        `gorm:"not null" json:"admin_fee"`
		JumlahCicilan float64        `gorm:"not null" json:"jumlah_cicilan"`
		JumlahBunga   float64        `gorm:"not null" json:"jumlah_bunga"`
		NamaAsset     string         `gorm:"not null" json:"nama_asset"`
		Customer      Customer       `gorm:"references:ID"`
		CustomerID    uint           `gorm:"not null" json:"customer_id"`
	}
)
