package services

import (
	"context"
	"fmt"

	"github.com/chanhlab/go-utils/timestamp"
	credentialv1 "github.com/chanhlab/golang-service-example/generated/go/credential/v1"
	"github.com/chanhlab/golang-service-example/internal/models"
)

// CredentialService ...
type CredentialService struct {
	CredentialRepository models.CredentialRepository
}

// NewCredentialService ...
func NewCredentialService(credentialRepository models.CredentialRepository) *CredentialService {
	return &CredentialService{
		CredentialRepository: credentialRepository,
	}
}

// List ...
func (s *CredentialService) List(_ context.Context, request *credentialv1.ListRequest) (*credentialv1.ListResponse, error) {
	offset := int(request.GetOffset())
	limit := int(request.GetLimit())
	if limit == 0 {
		limit = 10
	}
	var credentials []*models.Credential
	var err error

	credentials, err = s.CredentialRepository.GetCredentials(offset, limit)

	if err != nil {
		return nil, fmt.Errorf("Can not get Credentials: %+v", err)
	}
	pbCredentials := []*credentialv1.Credential{}
	for _, credential := range credentials {
		pbCredentials = append(pbCredentials, s.CredentialToProto(credential))
	}
	return &credentialv1.ListResponse{Credentials: pbCredentials}, nil
}

// Get ...
func (s *CredentialService) Get(_ context.Context, request *credentialv1.GetRequest) (*credentialv1.GetResponse, error) {
	id := request.GetId()
	var credential *models.Credential
	credential, err := s.CredentialRepository.GetCredential(id)
	if err != nil {
		return nil, fmt.Errorf("Credential not found with ID %s, Error: %+v", id, err)
	}
	return &credentialv1.GetResponse{Credential: s.CredentialToProto(credential)}, nil
}

// Create ...
func (s *CredentialService) Create(_ context.Context, request *credentialv1.CreateRequest) (*credentialv1.CreateResponse, error) {
	key := request.GetKey()
	value := request.GetValue()

	credential := &models.Credential{Key: key, Value: value}
	err := s.CredentialRepository.Create(credential)
	if err != nil {
		return nil, fmt.Errorf("Can not create Credential: %+v", err)
	}
	credential, _ = s.CredentialRepository.GetCredential(credential.ID)
	return &credentialv1.CreateResponse{Credential: s.CredentialToProto(credential)}, nil
}

// Update ...
func (s *CredentialService) Update(_ context.Context, request *credentialv1.UpdateRequest) (*credentialv1.UpdateResponse, error) {
	id := request.GetId()
	value := request.GetValue()

	credential, err := s.CredentialRepository.GetCredential(id)
	if err != nil {
		return nil, fmt.Errorf("Credential not found with ID %s", id)
	}

	credential.Value = value
	err = s.CredentialRepository.Update(credential)
	if err != nil {
		return nil, fmt.Errorf("Can not update Credential with ID %s, Errors: %v", id, err)
	}

	credential, _ = s.CredentialRepository.GetCredential(credential.ID)
	return &credentialv1.UpdateResponse{Credential: s.CredentialToProto(credential)}, nil
}

// Activate updates status of credential to active
func (s *CredentialService) Activate(_ context.Context, request *credentialv1.ActivateRequest) (*credentialv1.ActivateResponse, error) {
	id := request.GetId()

	credential, err := s.CredentialRepository.GetCredential(id)

	if err != nil {
		return nil, fmt.Errorf("Credential not found with ID %s", id)
	}

	err = s.CredentialRepository.Activate(credential)
	if err != nil {
		return nil, fmt.Errorf("Can not Activate Credential with ID %s, %+v", id, err)
	}

	return &credentialv1.ActivateResponse{Credential: s.CredentialToProto(credential)}, nil
}

// Deactivate updates status of credential to Inactive
func (s *CredentialService) Deactivate(_ context.Context, request *credentialv1.DeactivateRequest) (*credentialv1.DeactivateResponse, error) {
	id := request.GetId()

	credential, err := s.CredentialRepository.GetCredential(id)

	if err != nil {
		return nil, fmt.Errorf("Credential not found with ID %s", id)
	}

	err = s.CredentialRepository.Deactivate(credential)
	if err != nil {
		return nil, fmt.Errorf("Can not Deactivate Credential with ID %s, %+v", id, err)
	}

	return &credentialv1.DeactivateResponse{Credential: s.CredentialToProto(credential)}, nil
}

// Delete ...
func (s *CredentialService) Delete(_ context.Context, request *credentialv1.DeleteRequest) (*credentialv1.DeleteResponse, error) {
	id := request.GetId()
	credential := &models.Credential{ID: id}
	err := s.CredentialRepository.Delete(credential)
	if err != nil {
		return nil, fmt.Errorf("Can not Remove Credential with ID %s, %+v", id, err)
	}
	return &credentialv1.DeleteResponse{DeletedAt: timestamp.ProtoTimestampNow()}, nil
}

// CredentialToProto converts Credential to Proto Credential
func (s *CredentialService) CredentialToProto(credential *models.Credential) *credentialv1.Credential {
	return &credentialv1.Credential{
		Id:        credential.ID,
		Key:       credential.Key,
		Value:     credential.Value,
		Status:    credential.Status.String(),
		CreatedAt: timestamp.TimeToProtoTimestamp(credential.CreatedAt),
		UpdatedAt: timestamp.TimeToProtoTimestamp(credential.UpdatedAt),
	}
}
