package models

type StockAlert struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"not null"`
	Symbol      string  `gorm:"size:10;not null"`
	TargetPrice float64 `gorm:"type:decimal(10,2)"`
	Direction   string  `gorm:"type:enum('above','below')"`
	AlertSent   bool    `gorm:"default:false"`
}
