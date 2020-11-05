package main

import (
	"fmt"

	"github.com/chanhteam/golang-service-example/config"
	"github.com/chanhteam/golang-service-example/internal/models"
	"github.com/chanhteam/golang-service-example/pkg/db/mysql"
	"github.com/chanhteam/golang-service-example/pkg/logger"
)

// main ...
func main() {
	fmt.Printf("Migrate \n")
	config.NewConfig()
	config := config.AppConfig
	err := logger.Init(config.Logger.LogLevel, config.Logger.LogTimeFormat)
	if err != nil {
		panic(err)
	}

	db := mysql.GetConnection()

	// Create Credential table
	db.AutoMigrate(&models.Credential{})
}
