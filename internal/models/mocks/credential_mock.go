package mocks

import (
	"errors"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/chanhlab/golang-service-example/internal/models"
	"github.com/google/uuid"
)

// CredentialRepository interface
type CredentialRepository interface {
	GetCredentials(int, int) ([]*models.Credential, error)
	GetCredential(string) (*models.Credential, error)
	Create(*models.Credential) error
	Update(*models.Credential) error
	Activate(*models.Credential) error
	Deactivate(*models.Credential) error
	Delete(*models.Credential) error
}

const (
	MaxOffset = 100
)

// CredentialDBMock structure
type CredentialDBMock struct {
}

// NewCredentialRepository creates a new CredentialRepository
func NewCredentialRepository() CredentialRepository {
	return &CredentialDBMock{}
}

// GetCredentials returns a list of credentials
func (db *CredentialDBMock) GetCredentials(offset int, _ int) ([]*models.Credential, error) {
	var err error

	if offset > MaxOffset {
		return nil, errors.New("can not query database")
	}

	credentials := []*models.Credential{}
	credential := &models.Credential{
		ID:        uuid.New().String(),
		Key:       faker.Name(),
		Status:    models.CredentialActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	credentials = append(credentials, credential)
	return credentials, err
}

// GetCredential returns a single credential
func (db *CredentialDBMock) GetCredential(id string) (*models.Credential, error) {
	var err error
	if id == "" {
		return nil, errors.New("id can not be empty")
	}
	credential := &models.Credential{
		ID:        id,
		Key:       faker.Name(),
		Status:    models.CredentialActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return credential, err
}

// Create a new Credential
func (db *CredentialDBMock) Create(credential *models.Credential) error {
	var err error
	if credential.Key == "" {
		return errors.New("name can not be empty")
	}
	err = credential.BeforeCreate()
	credential.CreatedAt = time.Now()
	credential.UpdatedAt = time.Now()
	return err
}

// Update ...
func (db *CredentialDBMock) Update(credential *models.Credential) error {
	var err error
	if credential.ID == "" {
		return errors.New("id can not be empty")
	}
	credential.UpdatedAt = time.Now()
	return err
}

// Activate updates status to active
func (db *CredentialDBMock) Activate(credential *models.Credential) error {
	var err error
	if credential.ID == "empty" {
		return errors.New("id can not be empty")
	}
	credential.Status = models.CredentialActive
	return err
}

// Deactivate updates status to inactive
func (db *CredentialDBMock) Deactivate(credential *models.Credential) error {
	var err error
	if credential.ID == "empty" {
		return errors.New("id can not be empty")
	}
	credential.Status = models.CredentialInactive
	return err
}

// Delete ...
func (db *CredentialDBMock) Delete(credential *models.Credential) error {
	var err error
	if credential.ID == "" {
		return errors.New("id can not be empty")
	}
	return err
}
