package services

import (
	"context"
	"testing"

	"github.com/chanhlab/golang-service-example/internal/models/mocks"
	pb "github.com/chanhlab/golang-service-example/protobuf/v1/credential"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type CredentialServiceTestSuite struct {
	suite.Suite
	CredentialService *CredentialService
}

func (c *CredentialServiceTestSuite) SetupTest() {
	credentialRepository := mocks.NewCredentialRepository()
	c.CredentialService = NewCredentialService(
		credentialRepository,
	)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(CredentialServiceTestSuite))
}

func (c *CredentialServiceTestSuite) TestListCredentialsNormalCase() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Normal case
	request := &pb.ListCredentialRequest{
		Offset: 0,
		Limit:  10,
	}
	_, err := c.CredentialService.List(ctx, request)
	c.Assert().Nil(err)
}

func (c *CredentialServiceTestSuite) TestListCredentialsLimitIsZero() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Limit is zero
	request := &pb.ListCredentialRequest{
		Offset: 0,
		Limit:  0,
	}
	_, err := c.CredentialService.List(ctx, request)
	c.Assert().Nil(err)
}

func (c *CredentialServiceTestSuite) TestListCredentialsShouldReturnError() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Should return an error
	request := &pb.ListCredentialRequest{
		Offset: 101,
		Limit:  10,
	}
	credentials, err := c.CredentialService.List(ctx, request)
	c.Assert().NotNil(err)
	c.Assert().Nil(credentials)
}

func (c *CredentialServiceTestSuite) TestGetShouldReturnSingleRecord() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Normal Case
	request := &pb.GetCredentialRequest{
		Id: uuid.New().String(),
	}
	_, err := c.CredentialService.Get(ctx, request)
	c.Assert().Nil(err)
}

func (c *CredentialServiceTestSuite) TestGetShouldInvalidID() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Normal Case
	request := &pb.GetCredentialRequest{
		Id: "",
	}
	_, err := c.CredentialService.Get(ctx, request)
	c.Assert().NotNil(err)
}

func (c *CredentialServiceTestSuite) TestCreateCredential() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	request := &pb.CreateCredentialRequest{
		Key:   faker.Name(),
		Value: faker.Name(),
	}

	credential, err := c.CredentialService.Create(ctx, request)
	c.Assert().Nil(err)
	c.Assert().NotNil(credential)
}

func (c *CredentialServiceTestSuite) TestCreateCredentailShouldReturnError() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	request := &pb.CreateCredentialRequest{
		Key:   "",
		Value: "",
	}

	credential, err := c.CredentialService.Create(ctx, request)
	c.Assert().NotNil(err)
	c.Assert().Nil(credential)
}

func (c *CredentialServiceTestSuite) TestUpdateCredential() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	request := &pb.UpdateCredentialRequest{
		Id:    uuid.New().String(),
		Value: faker.Name(),
	}

	credential, err := c.CredentialService.Update(ctx, request)
	c.Assert().Nil(err)
	c.Assert().NotNil(credential)
}

func (c *CredentialServiceTestSuite) TestUpdateCredentialShouldReturnError() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	request := &pb.UpdateCredentialRequest{
		Id:    "",
		Value: faker.Name(),
	}

	credential, err := c.CredentialService.Update(ctx, request)
	c.Assert().NotNil(err)
	c.Assert().Nil(credential)
}

func (c *CredentialServiceTestSuite) TestActivateCredential() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	request := &pb.GetCredentialRequest{
		Id: uuid.New().String(),
	}

	credential, err := c.CredentialService.Activate(ctx, request)
	c.Assert().Nil(err)
	c.Assert().NotNil(credential)
}

func (c *CredentialServiceTestSuite) TestActivateCredentialShouldReturnError() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	request := &pb.GetCredentialRequest{
		Id: "",
	}

	credential, err := c.CredentialService.Activate(ctx, request)
	c.Assert().NotNil(err)
	c.Assert().Nil(credential)

	request = &pb.GetCredentialRequest{
		Id: "empty",
	}

	credential, err = c.CredentialService.Activate(ctx, request)
	c.Assert().NotNil(err)
	c.Assert().Nil(credential)
}

func (c *CredentialServiceTestSuite) TestDeactivateCredential() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	request := &pb.GetCredentialRequest{
		Id: uuid.New().String(),
	}

	credential, err := c.CredentialService.Deactivate(ctx, request)
	c.Assert().Nil(err)
	c.Assert().NotNil(credential)
}

func (c *CredentialServiceTestSuite) TestDeactivateCredentialShouldReturnError() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	request := &pb.GetCredentialRequest{
		Id: "",
	}

	credential, err := c.CredentialService.Deactivate(ctx, request)
	c.Assert().NotNil(err)
	c.Assert().Nil(credential)

	request = &pb.GetCredentialRequest{
		Id: "empty",
	}

	credential, err = c.CredentialService.Deactivate(ctx, request)
	c.Assert().NotNil(err)
	c.Assert().Nil(credential)
}

func (c *CredentialServiceTestSuite) TestDeleteCredential() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	request := &pb.DeleteCredentialRequest{
		Id: uuid.New().String(),
	}

	credential, err := c.CredentialService.Delete(ctx, request)
	c.Assert().Nil(err)
	c.Assert().NotNil(credential)
}

func (c *CredentialServiceTestSuite) TestDeleteCredentialShouldReturnError() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	request := &pb.DeleteCredentialRequest{
		Id: "",
	}

	credential, err := c.CredentialService.Delete(ctx, request)
	c.Assert().NotNil(err)
	c.Assert().Nil(credential)
}
