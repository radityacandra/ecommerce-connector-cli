package service

import (
	"context"
	"errors"

	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/client"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/model"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/repository"
)

type Service struct {
	Repository    repository.IRepository
	ProductClient client.ProductClient
}

type IService interface {
	Add(context.Context, model.Product) error
}

var (
	ErrProductConflicted = errors.New("product already exist")
)

func NewService(repository repository.IRepository, productClient client.ProductClient) *Service {
	return &Service{
		Repository:    repository,
		ProductClient: productClient,
	}
}
