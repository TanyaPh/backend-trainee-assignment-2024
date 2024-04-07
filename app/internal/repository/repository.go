package repository

import (
	"api/internal/repository/postgresql"
	"api/pkg/postgres"
)

type Auth interface {

}

type Admin interface {
	
}

type User interface {
	Get(tagId int, featureId int, lastRevision bool)
}

type Repository struct {
	Authorization Auth
	Admin Admin
	User User
}

func NewRepository(db *postgres.Postgres) *Repository {
	return &Repository{
		Authorization: postgresql.NewAuthPostgres(db),
		Admin: postgresql.NewAdminPostgres(db),
		User: postgresql.NewUserPostgres(db),
	}
}
