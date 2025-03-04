package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/radityacandra/ecommerce-connector-cli/internal/application/product/model"
)

func (s *Service) Add(ctx context.Context, input model.Product) error {
	// find by ean_code
	_, err := s.Repository.FindProductByEanCode(ctx, input.EanCode, input.UserId)
	if err == nil {
		err = ErrProductConflicted
		return err
	}
	err = nil

	product := model.NewProduct(uuid.NewString(), input.Name, input.Description, input.Category, input.Price, input.EanCode, input.UserId)
	err = s.Repository.Insert(ctx, product)

	// TODO: product integration

	return err
}
