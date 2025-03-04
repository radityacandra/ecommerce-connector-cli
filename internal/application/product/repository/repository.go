package repository

import (
	"context"
	"errors"

	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/model"
	"github.com/radityacandra/ecommerce-connector-cli/pkg/database"
)

type Repository struct {
	Db *database.Database
}

type IRepository interface {
	FindProductByEanCode(ctx context.Context, eanCode, userId string) (*model.Product, error)
	Insert(ctx context.Context, product *model.Product) error
}

var (
	ErrProductNotFound = errors.New("product not found")
)

func NewRepository(db *database.Database) *Repository {
	return &Repository{
		Db: db,
	}
}
