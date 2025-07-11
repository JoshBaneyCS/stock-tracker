package models

type FavoriteStock struct {
	ID            uint   `gorm:"primaryKey"`
	UserID        uint   `gorm:"not null"`
	Symbol        string `gorm:"size:10;not null"`
	DisplayName   string `gorm:"size:100"`
	Color         string `gorm:"size:7"` // Hex code like #FF0000
	IsMarketIndex bool
}
