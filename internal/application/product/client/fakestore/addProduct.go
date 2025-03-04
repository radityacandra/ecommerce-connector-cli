package fakestore

import (
	"context"

	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/model"
)

func (f *Fakestore) AddProduct(ctx context.Context, product model.Product) error {
	price := float32(product.Price)
	_, err := f.client.AddProduct(ctx, Product{
		Category:    &product.Category,
		Description: &product.Description,
		Price:       &price,
		Title:       &product.Name,
	})

	return err
}
