package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	SerialNumber string    `json:"serial_number"`
	UserRefer    int       `json:"user_id" gorm:"user_id"`
	User         User      `gorm:"foreignKey:UserRefer"`
	CreatedAt    time.Time `json:"created_at"`
}
