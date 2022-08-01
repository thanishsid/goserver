package db

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDB(pg *pgxpool.Pool) DB {
	if pg == nil {
		panic("postgres connection pool is null")
	}

	return &db{
		pg:      pg,
		Querier: New(pg),
	}
}

type TxFunc func(ctx context.Context, q Querier) error

type DB interface {
	Querier
	BeginTx(ctx context.Context, txOpts pgx.TxOptions) (Transactioner, error)
}

type db struct {
	Querier
	pg *pgxpool.Pool
}

type Transactioner interface {
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
	Querier
}

type txdb struct {
	pgx.Tx
	Querier
}

func (d *db) BeginTx(ctx context.Context, txOpts pgx.TxOptions) (Transactioner, error) {
	tx, err := d.pg.BeginTx(ctx, txOpts)
	if err != nil {
		return nil, err
	}

	txquerier := New(tx)

	return &txdb{
		Tx:      tx,
		Querier: txquerier,
	}, nil
}
