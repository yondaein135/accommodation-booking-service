package models

import (
	"gorm.io/gorm"
	"time"
)

type Booking struct {
	gorm.Model
	UserID      uint64    `gorm:"not null"`
	RoomID      uint64    `gorm:"not null"`
	StartDate   time.Time `gorm:"type:date"`
	EndDate     time.Time `gorm:"type:date"`
	Price       float32   `gorm:"type:decimal(10,2)"`
	Occupants   uint64    `gorm:"not null"`
	IsPaid      bool      `gorm:"default:false"`
	IsCancelled bool      `gorm:"default:false"`
}
