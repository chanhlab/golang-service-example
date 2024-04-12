package models

import (
	"database/sql"
	"testing"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CredentialTestSuite struct {
	suite.Suite
	DB                   *gorm.DB
	mock                 sqlmock.Sqlmock
	credentialRepository CredentialRepository
	credential           *Credential
}

func (c *CredentialTestSuite) SetupSuite() {
	var db *sql.DB
	var err error

	db, c.mock, _ = sqlmock.New()
	c.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	require.NoError(c.T(), err)
	c.credentialRepository = NewCredentialRepository(c.DB)

	c.credential = &Credential{
		ID:        uuid.New().String(),
		Key:       faker.Name(),
		Value:     faker.Name(),
		Status:    CredentialActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c *CredentialTestSuite) AfterTest(_, _ string) {
	c.Assert().NoError(c.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(CredentialTestSuite))
}

func (c *CredentialTestSuite) TestBeforeCreate() {
	credential := &Credential{}
	err := credential.BeforeCreate()
	c.Assert().Nil(err)
	c.Assert().NotEmpty(credential.ID)
}

func (c *CredentialTestSuite) TestGetCredentials() {
	offset := 0
	limit := 10
	credential := c.credential

	mockCredentials := sqlmock.NewRows([]string{"id", "key", "value", "status", "created_at", "updated_at"}).
		AddRow(credential.ID, credential.Key, credential.Value, credential.Status, credential.CreatedAt, credential.UpdatedAt)

	c.mock.ExpectQuery("SELECT (.+) FROM `credentials`").WillReturnRows(mockCredentials)

	credentials, err := c.credentialRepository.GetCredentials(offset, limit)
	c.Assert().NoError(err)
	c.Assert().Equal(credential.Key, credentials[0].Key)
	c.Assert().Equal(credential.Status.String(), credentials[0].Status.String())
}

func (c *CredentialTestSuite) TestGetCredential() {
	credential := c.credential

	mockCredentials := sqlmock.NewRows([]string{"id", "key", "value", "status", "created_at", "updated_at"}).
		AddRow(credential.ID, credential.Key, credential.Value, credential.Status, credential.CreatedAt, credential.UpdatedAt)

	c.mock.ExpectQuery("SELECT (.+) FROM `credentials`").WillReturnRows(mockCredentials)

	credential, err := c.credentialRepository.GetCredential(credential.ID)

	c.Assert().NoError(err)
	c.Assert().Equal(credential.Key, credential.Key)
	c.Assert().Equal(credential.Status.String(), credential.Status.String())
}

func (c *CredentialTestSuite) TestCreateCredential() {
	credential := c.credential

	var err error
	c.mock.ExpectBegin()
	c.mock.ExpectExec("INSERT INTO `credentials`").
		WithArgs(credential.ID, credential.Key, credential.Value, credential.Status, credential.CreatedAt, credential.UpdatedAt).WillReturnResult(sqlmock.NewErrorResult(err))
	c.mock.ExpectCommit()

	err = c.credentialRepository.Create(credential)
	c.Assert().NoError(err)
}

func (c *CredentialTestSuite) TestActivateCredential() {
	credential := c.credential

	var err error
	c.mock.ExpectBegin()
	c.mock.ExpectExec("UPDATE `credentials`").
		WithArgs(credential.Status, credential.ID).
		WillReturnResult(sqlmock.NewErrorResult(err))
	c.mock.ExpectCommit()

	err = c.credentialRepository.Activate(credential)
	c.Assert().NoError(err)
}

func (c *CredentialTestSuite) TestDeactivateCredential() {
	credential := c.credential
	credential.Status = CredentialInactive

	var err error
	c.mock.ExpectBegin()
	c.mock.ExpectExec("UPDATE `credentials`").
		WithArgs(credential.Status, credential.ID).
		WillReturnResult(sqlmock.NewErrorResult(err))
	c.mock.ExpectCommit()

	err = c.credentialRepository.Deactivate(credential)
	c.Assert().NoError(err)
}

func (c *CredentialTestSuite) TestDeleteCredential() {
	credential := c.credential

	var err error
	c.mock.ExpectBegin()
	c.mock.ExpectExec("DELETE FROM `credentials`").
		WithArgs(credential.ID).
		WillReturnResult(sqlmock.NewErrorResult(err))
	c.mock.ExpectCommit()

	err = c.credentialRepository.Delete(credential)
	c.Assert().NoError(err)
}
