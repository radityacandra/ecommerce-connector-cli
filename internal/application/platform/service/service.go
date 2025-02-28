package service

import (
	"context"

	"github.com/radityacandra/ecommerce-connector-cli/internal/application/platform/repository"
)

type Service struct {
	Repository repository.IRepository
}

type IService interface {
	Set(context.Context, SetPlatformInput) error
}

func NewService(repo repository.IRepository) *Service {
	return &Service{
		Repository: repo,
	}
}
