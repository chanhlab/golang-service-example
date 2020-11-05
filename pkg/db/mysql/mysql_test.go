package mysql

import (
	"testing"

	"github.com/chanhteam/golang-service-example/pkg/logger"

	"github.com/stretchr/testify/assert"
)

func TestGetConnectionShouldReturnNil(t *testing.T) {
	logger.NewDefault()
	err := Init()
	db := GetConnection()

	assert.NotNil(t, err)
	assert.Nil(t, db)
}

// func TestGetConnectionShouldReturnConnection(t *testing.T) {
// 	logger.NewDefault()
// 	os.Setenv("MYSQL_HOST", "10.30.3.105")
// 	os.Setenv("MYSQL_DATABASE", "shared_services")
// 	os.Setenv("MYSQL_USERNAME", "admin")
// 	os.Setenv("MYSQL_PASSWORD", "Hie8oox9ahhohsh")
// 	logger.NewDefault()
// 	Init()
// 	db := GetConnection()

// 	assert.NotNil(t, db)
// }

func TestInitShouldReturnError(t *testing.T) {
	logger.NewDefault()
	err := Init()

	assert.NotNil(t, err)
}

// func TestInitShouldReturnNil(t *testing.T) {
// 	os.Setenv("MYSQL_HOST", "10.30.3.105")
// 	os.Setenv("MYSQL_DATABASE", "shared_services")
// 	os.Setenv("MYSQL_USERNAME", "admin")
// 	os.Setenv("MYSQL_PASSWORD", "Hie8oox9ahhohsh")
// 	logger.NewDefault()
// 	err := Init()

// 	assert.Nil(t, err)
// }
