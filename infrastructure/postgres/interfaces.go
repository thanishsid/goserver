package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Transactioner interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}
