package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// CredentialStatus defines the credential status
type CredentialStatus int

const (
	// CredentialDefaultStatus defines an default status
	CredentialDefaultStatus CredentialStatus = iota

	// CredentialActive defines an active credential
	CredentialActive

	// CredentialInactive defines a inactive credential
	CredentialInactive
)

var types = [...]string{
	"default",
	"active",
	"inactive",
}

func (credentialStatus CredentialStatus) String() string {
	return types[credentialStatus]
}

// Credential ...
type Credential struct {
	ID        string           `gorm:"primary_key"`
	Key       string           `gorm:"type:varchar(255) not null;unique"`
	Value     string           `gorm:"type:varchar(255) not null"`
	Status    CredentialStatus `gorm:"type:int(11) not null default 1"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CredentialRepository interface
type CredentialRepository interface {
	GetCredentials(offset int64, limit int64) ([]*Credential, error)
	GetCredential(id string) (*Credential, error)
	Create(*Credential) error
	Update(*Credential) error
	Activate(*Credential) error
	Deactivate(*Credential) error
	Delete(*Credential) error
}

// CredentialDB structure
type CredentialDB struct {
	DB *gorm.DB
}

// NewCredentialRepository creates a new CredentialRepository
func NewCredentialRepository(db *gorm.DB) CredentialRepository {
	return &CredentialDB{
		DB: db,
	}
}

// BeforeCreate generates UUID
func (m *Credential) BeforeCreate() error {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.Status = CredentialActive
	return nil
}

// GetCredentials returns a list of credentials
func (db *CredentialDB) GetCredentials(offset int64, limit int64) ([]*Credential, error) {
	credentials := []*Credential{}
	err := db.DB.Offset(offset).Limit(limit).
		Find(&credentials).Error
	return credentials, err
}

// GetCredential returns a single credential
func (db *CredentialDB) GetCredential(id string) (*Credential, error) {
	credential := &Credential{}
	err := db.DB.Where(&Credential{ID: id}).First(credential).Error
	return credential, err
}

// Create a new Credential
func (db *CredentialDB) Create(credential *Credential) error {
	err := db.DB.Create(&credential).Error
	return err
}

// Update a credential
func (db *CredentialDB) Update(credential *Credential) error {
	err := db.DB.Save(&credential).Error
	return err
}

// Activate updates status to active
func (db *CredentialDB) Activate(credential *Credential) error {
	err := db.DB.Model(&credential).UpdateColumns(&Credential{Status: CredentialActive}).Error
	return err
}

// Deactivate updates status to inactive
func (db *CredentialDB) Deactivate(credential *Credential) error {
	err := db.DB.Model(&credential).
		UpdateColumns(&Credential{Status: CredentialInactive}).Error
	return err
}

// Delete ...
func (db *CredentialDB) Delete(credential *Credential) error {
	err := db.DB.Delete(&credential).Error
	return err
}
