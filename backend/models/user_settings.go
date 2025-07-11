package models

type UserSettings struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"uniqueIndex"`
	BaseCurrency string `gorm:"size:10;default:USD"`
}
