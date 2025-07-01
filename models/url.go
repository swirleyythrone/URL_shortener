package models

import "time"

type URL struct {
	ID        uint      `gorm:"primaryKey"`
	LongURL   string    `gorm:"not null"`
	ShortID   string    `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	Clicks    int
}