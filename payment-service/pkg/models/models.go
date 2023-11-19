package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	BookingID uint64
	Amount    float32
	Success   bool
}
