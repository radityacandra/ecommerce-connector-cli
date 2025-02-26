package repository

import (
	"context"
	"errors"

	"github.com/radityacandra/ecommerce-connector-cli/internal/application/user/model"
)

func (r *Repository) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	docs := r.Db.Client.Collection("users").Where("email", "==", email).Documents(ctx)

	defer docs.Stop()
	doc, err := docs.Next()
	if err != nil {
		return nil, errors.Join(err, ErrDataNotFound)
	}

	user := model.User{}
	doc.DataTo(&user)
	user.Id = doc.Ref.ID

	return &user, nil
}
