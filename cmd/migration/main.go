package main

import (
	"fmt"

	"github.com/chanhteam/go-utils/database/mysql"
	"github.com/chanhteam/go-utils/logger"
	"github.com/chanhteam/golang-service-example/config"
	"github.com/chanhteam/golang-service-example/internal/models"
)

// main ...
func main() {
	fmt.Printf("Migrate \n")
	config.NewConfig()
	config := config.AppConfig
	logger.Init(config.Logger.LogLevel, config.Logger.LogTimeFormat)
	db := mysql.GetConnection(config.MySQL.Host, config.MySQL.DBName, config.MySQL.Username, config.MySQL.Password, config.MySQL.MaxIDLEConnection, config.MySQL.MaxOpenConnection)
	// Create Credential table
	db.AutoMigrate(&models.Credential{})
}
