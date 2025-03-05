package core

import (
	"github.com/go-playground/validator/v10"
	"github.com/radityacandra/ecommerce-connector-cli/pkg/database"
)

type Dependency struct {
	Db        *database.Database
	Validator *validator.Validate
}

func NewDependency(Db *database.Database, validator *validator.Validate) *Dependency {
	return &Dependency{
		Db:        Db,
		Validator: validator,
	}
}
