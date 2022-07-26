// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: role_queries.sql

package db

import (
	"context"
)

const deleteRole = `-- name: DeleteRole :exec
DELETE FROM roles WHERE id = $1::TEXT
`

func (q *Queries) DeleteRole(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteRole, id)
	return err
}

const getAllRoles = `-- name: GetAllRoles :many
SELECT id, name FROM roles
`

func (q *Queries) GetAllRoles(ctx context.Context) ([]Role, error) {
	rows, err := q.db.Query(ctx, getAllRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertOrUpdateRoles = `-- name: InsertOrUpdateRoles :exec
INSERT INTO roles (
    id, name
) VALUES (
    $1::TEXT,
    $2::TEXT
) 
ON CONFLICT (id)
DO
UPDATE SET
    name = EXCLUDED.name
`

type InsertOrUpdateRolesParams struct {
	ID   string
	Name string
}

func (q *Queries) InsertOrUpdateRoles(ctx context.Context, arg InsertOrUpdateRolesParams) error {
	_, err := q.db.Exec(ctx, insertOrUpdateRoles, arg.ID, arg.Name)
	return err
}
