package repository

import (
	"fmt"
	"log"

	"backend/internal/config"
	"backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connection established")

	// AutoMigrate
	err = db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Account{},
		&models.Transaction{},
		&models.Budget{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate: %v", err)
	}
	
	log.Println("Database tables migrated successfully")

	return db, nil
}
