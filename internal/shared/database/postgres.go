package database

import (
	"log"

	"github.com/seyamibrahim/movie-reservation-system/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *configs.AppConfig) (*gorm.DB, error) {
	// Connect to Database
	db, err := gorm.Open(postgres.Open(cfg.DBURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Set Connection Pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// Connection pool configuration
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(10)
	// Verify Database Connection
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	log.Printf("Successfully Connected to Database")

	return db, nil
}
