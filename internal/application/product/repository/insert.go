package repository

import (
	"context"

	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/model"
)

func (r *Repository) Insert(ctx context.Context, product *model.Product) error {
	_, err := r.Db.Client.Collection("products").Doc(product.Id).Create(ctx, product)

	return err
}
