package setup

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/thanishsid/goserver/domain"
	"github.com/thanishsid/goserver/infrastructure/db"
)

func NewSeeder(dbs db.DB) Seeder {
	return &seeder{dbs}
}

type Seeder interface {
	// Update the database with the roles specified in the security package.
	UpdateRoles(ctx context.Context) error

	// Check if a admin account exists in the database, if not will guide through creating a new account.
	CreateAdmin(ctx context.Context) error
}

type seeder struct {
	db.DB
}

// Update the database with the roles specified in the security package.
func (s *seeder) UpdateRoles(ctx context.Context) error {

	tx, err := s.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, role := range domain.AllRoles {
		if err := tx.InsertOrUpdateRoles(ctx, db.InsertOrUpdateRolesParams{
			ID:   string(role),
			Name: strings.Title(string(role)),
		}); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

// Check if a admin account exists in the database, if not will guide through creating a new account.
func (s *seeder) CreateAdmin(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}
