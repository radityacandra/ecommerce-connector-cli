package repository

import (
	"context"

	"github.com/radityacandra/ecommerce-connector-cli/internal/application/user/model"
)

func (r *Repository) Insert(ctx context.Context, user *model.User) error {
	_, err := r.Db.Client.Collection("users").Doc(user.Id).Create(ctx, user)

	return err
}
