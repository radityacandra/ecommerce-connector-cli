package repository

import (
	"context"
	"errors"

	"github.com/radityacandra/ecommerce-connector-cli/internal/application/user/model"
	"github.com/radityacandra/ecommerce-connector-cli/pkg/database"
)

var (
	ErrDataNotFound = errors.New("data not found")
)

type IRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
}

type Repository struct {
	Db *database.Database
}

func NewRepository(db *database.Database) *Repository {
	return &Repository{
		Db: db,
	}
}
