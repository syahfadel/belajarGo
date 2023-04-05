package entity

import (
	"time"
)

type Book struct {
	ID        uint   `gorm:"primaryKey"`
	NameBook  string `gorm:"not null"`
	Author    string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
