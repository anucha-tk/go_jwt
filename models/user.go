package models

import "time"

type User struct {
	ID        uint
	Name      string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
