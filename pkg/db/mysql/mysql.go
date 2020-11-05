package mysql

import (
	"fmt"
	"strconv"
	"time"

	"github.com/chanhteam/golang-service-example/pkg"
	"github.com/chanhteam/golang-service-example/pkg/logger"

	"github.com/jinzhu/gorm"

	// Register some standard stuff
	_ "github.com/jinzhu/gorm/dialects/mysql"
	gormzap "github.com/wantedly/gorm-zap"
)

var connect *gorm.DB

// Mysql ...
type Mysql struct {
	*gorm.DB
}

// Init ...
func Init() error {
	host := pkg.GetEnv("MYSQL_HOST", "127.0.0.1")
	database := pkg.GetEnv("MYSQL_DATABASE", "db_name")
	username := pkg.GetEnv("MYSQL_USERNAME", "root")
	password := pkg.GetEnv("MYSQL_PASSWORD", "")

	maxIdleConnection, _ := strconv.ParseInt(pkg.GetEnv("MYSQL_MAX_IDLE_CONNECTION", "10"), 10, 64)
	maxOpenConnection, _ := strconv.ParseInt(pkg.GetEnv("MYSQL_MAX_OPEN_CONNECTION", "100"), 10, 64)

	var err error
	if connect == nil {
		strConnect := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, database)
		connect, err = gorm.Open("mysql", strConnect)

		if err == nil {
			connect.LogMode(false)
			connect.SetLogger(gormzap.New(logger.Log))
			connect.DB().SetConnMaxLifetime(time.Hour)
			connect.DB().SetMaxOpenConns(int(maxOpenConnection))
			connect.DB().SetMaxIdleConns(int(maxIdleConnection))
		} else {
			connect = nil
		}
		logger.Log.Info("First connection")
	}
	logger.Log.Info("Get MySQL Connection")
	return err
}

// GetConnection ...
func GetConnection() *gorm.DB {
	var err error
	if connect == nil {
		err = Init()
	}
	if err == nil {
		return connect
	}
	return nil
}
