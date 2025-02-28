package repository

import (
	"context"
	"errors"
)

type FindByUserIdAndTypeInput struct {
	UserId string
	Type   string
}

type FindByUserIdAndTypeOutput struct {
	Id     string
	UserId string `firestore:"user_id"`
	Type   string `firestore:"type"`
}

func (r *Repository) FindPlatformByUserIdAndType(ctx context.Context, input FindByUserIdAndTypeInput) (
	output FindByUserIdAndTypeOutput, err error) {
	docs := r.Db.Client.Collection("platforms").
		Where("user_id", "==", input.UserId).
		Where("type", "==", input.Type).Documents(ctx)

	defer docs.Stop()
	doc, err := docs.Next()
	if err != nil {
		err = errors.Join(err, ErrDataNotFound)
		return
	}

	doc.DataTo(&output)
	output.Id = doc.Ref.ID

	return
}
