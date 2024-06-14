package main

import (
	"BakeMeets/api/models"
	"BakeMeets/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.GetConfig()
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to database: %v\n", err)
		return
	}

	// Automigrate will create or update the tables based on the models
	err = db.AutoMigrate(&models.Bread{}, &models.User{}, &models.Comment{})
	if err != nil {
		fmt.Printf("Error during migration: %v\n", err)
		return
	}

	fmt.Println("Migration completed successfully")
}
