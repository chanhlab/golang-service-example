package migration

import (
	"context"
	"fmt"

	"github.com/chanhlab/go-utils/database/mysql"
	"github.com/chanhlab/go-utils/logger"
	"github.com/chanhlab/golang-service-example/config"
	"github.com/chanhlab/golang-service-example/internal/models"
)

// main ...
func RunMigration(ctx context.Context) error {
	fmt.Printf("Migrate \n")
	config.NewConfig()
	config := config.AppConfig
	logger.Init(config.Logger.LogLevel, config.Logger.LogTimeFormat)
	db := mysql.GetConnection(config.MySQL.Host, config.MySQL.Port, config.MySQL.DBName, config.MySQL.Username, config.MySQL.Password, config.MySQL.MaxIDLEConnection, config.MySQL.MaxOpenConnection)
	// Create Credential table
	return db.AutoMigrate(&models.Credential{})
}
