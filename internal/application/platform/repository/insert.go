package repository

import (
	"context"

	"github.com/radityacandra/ecommerce-connector-cli/internal/application/platform/model"
)

func (r *Repository) Insert(ctx context.Context, input *model.Platform) error {
	_, err := r.Db.Client.Collection("platforms").Doc(input.Id).Create(ctx, input)
	return err
}
