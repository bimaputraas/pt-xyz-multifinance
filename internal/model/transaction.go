package model

type (
	Transaction struct {
		ID            uint    `gorm:"column:id;primaryKey" json:"id,omitempty"`
		Datetime      string  `gorm:"column:datetime" json:"datetime,omitempty"`
		NomorKontrak  string  `gorm:"column:nomor_kontrak;unique;not null" json:"nomor_kontrak"`
		OTR           float64 `gorm:"column:otr;not null" json:"otr" validate:"required"`
		AdminFee      float64 `gorm:"column:admin_fee;not null" json:"admin_fee"`
		Tenor         int     `gorm:"column:tenor;not null" json:"tenor" validate:"required"`
		JumlahCicilan float64 `gorm:"column:jumlah_cicilan;not null" json:"jumlah_cicilan"`
		JumlahBunga   float64 `gorm:"column:jumlah_bunga;not null" json:"jumlah_bunga"`
		NamaAsset     string  `gorm:"column:nama_asset;not null" json:"nama_asset" validate:"required"`
		UserID        uint    `gorm:"column:user_id;not null" json:"user_id"`
	}
)
