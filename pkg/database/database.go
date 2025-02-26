package database

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Database struct {
	Client *firestore.Client
}

var (
	ErrConnect = errors.New("failed to connect to firestore")
)

func NewDatabase(ctx context.Context, credentialFilePath string) (*Database, error) {
	opt := option.WithCredentialsFile(credentialFilePath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, errors.Join(err, ErrConnect)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, errors.Join(err, ErrConnect)
	}

	return &Database{
		Client: client,
	}, nil
}
