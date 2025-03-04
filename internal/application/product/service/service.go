package service

import (
	"context"
	"errors"

	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/client/fakestore"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/model"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/repository"
)

type Service struct {
	Repository      repository.IRepository
	FakestoreClient fakestore.ClientInterface
}

type IService interface {
	Add(context.Context, model.Product) error
}

var (
	ErrProductConflicted = errors.New("product already exist")
)

func NewService(repository repository.IRepository) *Service {
	return &Service{
		Repository: repository,
	}
}
