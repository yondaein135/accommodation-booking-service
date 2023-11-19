package database

import (
	"accommodation-booking-system/property-service/pkg/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type DatabaseManager struct {
	DB *gorm.DB
}

func NewDatabaseManager(dsn string) (*DatabaseManager, error) {
	var db *gorm.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database. Retry in 5 seconds... (Attempt %d/5)", i+1)
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after retries: %v", err)
	}

	err = db.AutoMigrate(&models.Property{}, &models.Room{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto-migrate database: %v", err)
	}

	log.Println("Database auto-migration completed successfully")

	return &DatabaseManager{DB: db}, nil
}
