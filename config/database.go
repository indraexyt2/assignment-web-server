package config

import (
	"fmt"
	"golang-web-server/models"
	"golang-web-server/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func SetupDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Logger.Fatal("Failed to connect to database: ", err)
		return
	}

	err = db.AutoMigrate(&models.User{}, &models.Product{}, &models.Inventory{}, &models.Order{}, &models.OrderItems{})
	if err != nil {
		utils.Logger.Fatal("Failed to migrate database: ", err)
		return
	}

	DB = db
	utils.Logger.Info("Connected to database")
}
