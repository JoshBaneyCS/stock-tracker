package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey"`
	FirstName    string    `gorm:"size:50"`
	LastName     string    `gorm:"size:50"`
	Email        string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	IPAddress    string    `gorm:"size:45"`
	CreatedAt    time.Time
	Favorites    []FavoriteStock
	Settings     UserSettings
	Alerts       []StockAlert
}
