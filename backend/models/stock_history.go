package models

import "time"

type StockHistory struct {
	ID          uint      `gorm:"primaryKey"`
	Symbol      string    `gorm:"size:10;not null"`
	JSONData    string    `gorm:"type:text"`
	LastUpdated time.Time `gorm:"autoUpdateTime"`
}
