package repository

import (
	"context"
	"errors"

	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/model"
)

func (r *Repository) FindProductByEanCode(ctx context.Context, eanCode, userId string) (*model.Product, error) {
	docs := r.Db.Client.
		Collection("products").
		Where("ean_code", "==", eanCode).
		Where("user_id", "==", userId).
		Documents(ctx)

	defer docs.Stop()
	doc, err := docs.Next()
	if err != nil {
		return nil, errors.Join(err, ErrProductNotFound)
	}

	var product model.Product
	doc.DataTo(&product)
	product.Id = doc.Ref.ID

	return nil, nil
}
