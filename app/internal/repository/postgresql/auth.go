package postgresql

import "api/pkg/postgres"

type AuthPostgres struct {
	db *postgres.Postgres
}

func NewAuthPostgres(db *postgres.Postgres) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// func (r *AuthPostgres) Get() (int, error) {

// 	return orderId, nil
// }