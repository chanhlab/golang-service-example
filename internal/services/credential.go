package services

import (
	"context"
	"fmt"

	"github.com/chanhlab/go-utils/timestamp"
	"github.com/chanhlab/golang-service-example/internal/models"

	pb "github.com/chanhlab/golang-service-example/protobuf/v1/credential"
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
func (s *CredentialService) List(ctx context.Context, request *pb.ListCredentialRequest) (*pb.ListCredentialResponse, error) {
	offset := request.GetOffset()
	limit := request.GetLimit()
	if limit == 0 {
		limit = 10
	}
	var credentials []*models.Credential
	var err error

	credentials, err = s.CredentialRepository.GetCredentials(offset, limit)

	if err != nil {
		return nil, fmt.Errorf("Can not get Credentials: %+v", err)
	}
	pbCredentials := []*pb.Credential{}
	for _, credential := range credentials {
		pbCredentials = append(pbCredentials, s.CredentialToProto(credential))
	}
	return &pb.ListCredentialResponse{Credentials: pbCredentials}, nil
}

// Get ...
func (s *CredentialService) Get(ctx context.Context, request *pb.GetCredentialRequest) (*pb.GetCredentialResponse, error) {
	id := request.GetId()
	var credential *models.Credential
	credential, err := s.CredentialRepository.GetCredential(id)
	if err != nil {
		return nil, fmt.Errorf("Credential not found with ID %s, Error: %+v", id, err)
	}
	return &pb.GetCredentialResponse{Credential: s.CredentialToProto(credential)}, nil
}

// Create ...
func (s *CredentialService) Create(ctx context.Context, request *pb.CreateCredentialRequest) (*pb.CreateCredentialResponse, error) {
	key := request.GetKey()
	value := request.GetValue()

	credential := &models.Credential{Key: key, Value: value}
	err := s.CredentialRepository.Create(credential)
	if err != nil {
		return nil, fmt.Errorf("Can not create Credential: %+v", err)
	}
	credential, _ = s.CredentialRepository.GetCredential(credential.ID)
	return &pb.CreateCredentialResponse{Credential: s.CredentialToProto(credential)}, nil
}

// Update ...
func (s *CredentialService) Update(ctx context.Context, request *pb.UpdateCredentialRequest) (*pb.UpdateCredentialResponse, error) {
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
	return &pb.UpdateCredentialResponse{Credential: s.CredentialToProto(credential)}, nil
}

// Activate updates status of credential to active
func (s *CredentialService) Activate(ctx context.Context, request *pb.GetCredentialRequest) (*pb.UpdateCredentialResponse, error) {
	id := request.GetId()

	credential, err := s.CredentialRepository.GetCredential(id)

	if err != nil {
		return nil, fmt.Errorf("Credential not found with ID %s", id)
	}

	err = s.CredentialRepository.Activate(credential)
	if err != nil {
		return nil, fmt.Errorf("Can not Activate Credential with ID %s, %+v", id, err)
	}

	return &pb.UpdateCredentialResponse{Credential: s.CredentialToProto(credential)}, nil
}

// Deactivate updates status of credential to Inactive
func (s *CredentialService) Deactivate(ctx context.Context, request *pb.GetCredentialRequest) (*pb.UpdateCredentialResponse, error) {
	id := request.GetId()

	credential, err := s.CredentialRepository.GetCredential(id)

	if err != nil {
		return nil, fmt.Errorf("Credential not found with ID %s", id)
	}

	err = s.CredentialRepository.Deactivate(credential)
	if err != nil {
		return nil, fmt.Errorf("Can not Deactivate Credential with ID %s, %+v", id, err)
	}

	return &pb.UpdateCredentialResponse{Credential: s.CredentialToProto(credential)}, nil
}

// Delete ...
func (s *CredentialService) Delete(ctx context.Context, request *pb.DeleteCredentialRequest) (*pb.DeleteCredentialResponse, error) {
	id := request.GetId()
	credential := &models.Credential{ID: id}
	err := s.CredentialRepository.Delete(credential)
	if err != nil {
		return nil, fmt.Errorf("Can not Remove Credential with ID %s, %+v", id, err)
	}
	return &pb.DeleteCredentialResponse{DeletedAt: timestamp.TimestampProtoNow()}, nil
}

// CredentialToProto converts Credential to Proto Credential
func (s *CredentialService) CredentialToProto(credential *models.Credential) *pb.Credential {
	return &pb.Credential{
		Id:        credential.ID,
		Key:       credential.Key,
		Value:     credential.Value,
		Status:    credential.Status.String(),
		CreatedAt: timestamp.TimeToTimestampProto(credential.CreatedAt),
		UpdatedAt: timestamp.TimeToTimestampProto(credential.UpdatedAt),
	}
}
