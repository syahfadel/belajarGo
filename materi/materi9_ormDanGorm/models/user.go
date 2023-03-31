package models

import "time"

/*
jika ingin membuat tabel menggunakan GORM kita perlu membuat modelsnya terlebih dahulu
GORM menyediakan tags yang dapat digunakan seperti halnya pada property ID dan Email
*/

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"not null; unique;type:varchar(191)"`
	Products  []Product //merupakan assosiasi dengan table product sehingga akan memiliki foreignKey dari ID User (one to many)
	CreatedAt time.Time
	UpdatedAt time.Time
}
