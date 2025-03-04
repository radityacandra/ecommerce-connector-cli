package client

import (
	"context"

	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/model"
)

type ProductClient interface {
	AddProduct(ctx context.Context, product model.Product) error
}
