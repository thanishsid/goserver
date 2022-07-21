package repository

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/infrastructure/postgres"
	"github.com/thanishsid/goserver/internal/search"
)

// Initiate and return repositories.
func New(pg *pgxpool.Pool, searcher *search.Searcher) Repository {
	return &repo{
		pg:           pg,
		searcher:     searcher,
		TxRepository: NewTxRepository(postgres.New(pg), searcher),
	}
}

type TxFunc func(ctx context.Context, repo TxRepository) error

type Repository interface {
	ExecTx(ctx context.Context, txOpts pgx.TxOptions, txFunc TxFunc) error
	TxRepository
}

type repo struct {
	pg       *pgxpool.Pool
	searcher *search.Searcher
	TxRepository
}

// Execute transactions accross multiple repositories.
func (r *repo) ExecTx(ctx context.Context, txOpts pgx.TxOptions, txFunc TxFunc) error {
	tx, err := r.pg.BeginTx(ctx, txOpts)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	txquerier := postgres.New(tx)

	if err := txFunc(ctx, NewTxRepository(txquerier, r.searcher)); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

//*------------- TxRepository --------------

type TxRepository interface {
	UserRepository() domain.UserRepository
}

type txRepo struct {
	userRepository domain.UserRepository
}

// Get user repository.
func (t *txRepo) UserRepository() domain.UserRepository {
	return t.userRepository
}

// Get a new TxRepository (used within transactions).
func NewTxRepository(q postgres.Querier, searcher *search.Searcher) TxRepository {
	return &txRepo{
		userRepository: &userRepository{db: q, searchIndex: searcher.Users},
	}
}
