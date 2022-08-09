package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"

	"github.com/thanishsid/goserver/domain"
)

// Seed the database with default values.
func Seed(ctx context.Context, dbs DB) error {
	tx, err := dbs.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err := updateRoles(ctx, tx); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// Update Roles.
func updateRoles(ctx context.Context, q Querier) error {

	dbRoles, err := q.GetAllRoles(ctx)
	if err != nil {
		return err
	}

	// Check if roles exists in the database that do not exist in the application and remove them if possible.
	if len(dbRoles) > 0 {
		for _, dbRole := range dbRoles {

			role := domain.Role(dbRole.ID)

			if err := role.ValidateRole(); err != nil {

				if err := q.DeleteRole(ctx, dbRole.ID); err != nil {
					return fmt.Errorf("database role with id '%s' does not exist in the application and cannot be deleted, "+
						"details: %s", role, err.Error())
				}

			}

		}
	}

	for _, role := range domain.Roles {
		if err := q.InsertOrUpdateRoles(ctx, InsertOrUpdateRolesParams{
			ID:   role.ID.String(),
			Name: role.Name,
		}); err != nil {
			return err
		}
	}

	return nil
}
