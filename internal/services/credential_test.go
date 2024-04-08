package services

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	credentialv1 "github.com/chanhlab/golang-service-example/generated/go/credential/v1"
	"github.com/chanhlab/golang-service-example/internal/models/mocks"
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
	request := &credentialv1.ListRequest{
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
	request := &credentialv1.ListRequest{
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
	request := &credentialv1.ListRequest{
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
	request := &credentialv1.GetRequest{
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
	request := &credentialv1.GetRequest{
		Id: "",
	}
	_, err := c.CredentialService.Get(ctx, request)
	c.Assert().NotNil(err)
}

func (c *CredentialServiceTestSuite) TestCreateCredential() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	request := &credentialv1.CreateRequest{
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

	request := &credentialv1.CreateRequest{
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

	request := &credentialv1.UpdateRequest{
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

	request := &credentialv1.UpdateRequest{
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

	request := &credentialv1.ActivateRequest{
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

	request := &credentialv1.ActivateRequest{
		Id: "",
	}

	credential, err := c.CredentialService.Activate(ctx, request)
	c.Assert().NotNil(err)
	c.Assert().Nil(credential)

	request = &credentialv1.ActivateRequest{
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

	request := &credentialv1.DeactivateRequest{
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

	request := &credentialv1.DeactivateRequest{
		Id: "",
	}

	credential, err := c.CredentialService.Deactivate(ctx, request)
	c.Assert().NotNil(err)
	c.Assert().Nil(credential)

	request = &credentialv1.DeactivateRequest{
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

	request := &credentialv1.DeleteRequest{
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

	request := &credentialv1.DeleteRequest{
		Id: "",
	}

	credential, err := c.CredentialService.Delete(ctx, request)
	c.Assert().NotNil(err)
	c.Assert().Nil(credential)
}
