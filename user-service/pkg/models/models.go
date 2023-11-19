package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string
	Email            string
	BillingAddresses []BillingAddress `gorm:"foreignKey:UserID"`
}

type BillingAddress struct {
	gorm.Model
	UserID       uint `gorm:"index"`
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	Country      string
	PostalCode   string
}
