package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null; type:varchar(191)"`
	Brand     string `gorm:"not null; type:varchar(191)"`
	UserID    uint   //otomatis menjadi foreign key karena property ini memliki nama berupa gabungan dara nama struct lain dendgan primarykey struct tersebut
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GROM neyediakan hooks yang dapat digunakan salah satunya adalah method BeforeCreate yang akan dieksekusi sebelum melakukan create data

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Product Before Create()")

	if len(p.Name) < 4 {
		err = errors.New("Product name is too short")
	}

	return
}
