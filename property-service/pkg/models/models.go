package models

import (
	"gorm.io/gorm"
)

type Property struct {
	gorm.Model
	Name     string
	Location string
	Rooms    []Room
}

type Room struct {
	gorm.Model
	ID         uint64
	PropertyID uint64 `gorm:"not null"`
	Type       string
	Capacity   int
	Price      float32 `gorm:"type:decimal(10,2)"`
}
