package repository

import (
	"context"
	"errors"

	"github.com/radityacandra/ecommerce-connector-cli/internal/application/platform/model"
	"github.com/radityacandra/ecommerce-connector-cli/pkg/database"
)

type Repository struct {
	Db *database.Database
}

type IRepository interface {
	FindPlatformByUserIdAndType(context.Context, FindByUserIdAndTypeInput) (FindByUserIdAndTypeOutput, error)
	Insert(context.Context, *model.Platform) error
	GetPlatformByUserId(ctx context.Context, userId string) ([]FindByUserIdAndTypeOutput, error)
}

var (
	ErrDataNotFound = errors.New("data not found")
)

func NewRepository(db *database.Database) *Repository {
	return &Repository{
		Db: db,
	}
}
