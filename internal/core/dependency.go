package core

import "github.com/radityacandra/ecommerce-connector-cli/pkg/database"

type Dependency struct {
	Db *database.Database
}

func NewDependency(Db *database.Database) *Dependency {
	return &Dependency{
		Db: Db,
	}
}
