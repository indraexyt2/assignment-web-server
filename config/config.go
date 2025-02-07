package config

import (
	"github.com/joho/godotenv"
	"golang-web-server/utils"
)

func SetupConfig() {
	// setup logger
	utils.SetupLogger()

	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		utils.Logger.Fatal("Error loading .env file: ", err)
		return
	}

	utils.Logger.Info("Loaded .env file")

	// setup database
	SetupDB()

}
