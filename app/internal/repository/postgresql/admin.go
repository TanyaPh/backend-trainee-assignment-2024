package postgresql

import "api/pkg/postgres"

type AdminPostgres struct {
	db *postgres.Postgres
}

func NewAdminPostgres(db *postgres.Postgres) *AdminPostgres {
	return &AdminPostgres{db: db}
}

// func (r *AdminPostgres) Get() (int, error) {

// 	return orderId, nil
// }