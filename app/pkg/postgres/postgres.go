package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Postgres struct {
	db *pgxpool.Pool
}

func New(ctx context.Context, connStr string) (*Postgres, error) {
	ctx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, err
	}
	defer pool.Close()
	logrus.Infof("Connected!")			////////

	return &Postgres{db: pool}, nil
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.db.Close()
}
