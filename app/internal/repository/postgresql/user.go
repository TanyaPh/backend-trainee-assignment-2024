package postgresql

import "api/pkg/postgres"

type UserPostgres struct {
	db *postgres.Postgres
}

func NewUserPostgres(db *postgres.Postgres) *UserPostgres {
	return &UserPostgres{db: db}
}

// func (r *UserPostgres) Get(tagId int, featureId int, lastRevision bool) (int, error) {

// 	return orderId, nil
// }
