package model

import "time"

type User struct {
	Id uint `gorm:"primaryKey"`
	Username string `gorm:"column:username"`
	Email string `gorm:"column:email"`
	Phone string `gorm:"column:phone"`
	Password string `gorm:"column:password"`
    CreatedAt time.Time `gorm:"column:created_at"`
}