// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: image_queries.sql

package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	null "gopkg.in/guregu/null.v4"
)

const checkImageExists = `-- name: CheckImageExists :one
SELECT EXISTS(SELECT 1 FROM images WHERE id = $1::UUID)
`

func (q *Queries) CheckImageExists(ctx context.Context, id uuid.UUID) (bool, error) {
	row := q.db.QueryRow(ctx, checkImageExists, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const checkImageHashExists = `-- name: CheckImageHashExists :one
SELECT EXISTS(SELECT 1 FROM images WHERE file_hash = $1::TEXT)
`

func (q *Queries) CheckImageHashExists(ctx context.Context, fileHash string) (bool, error) {
	row := q.db.QueryRow(ctx, checkImageHashExists, fileHash)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const deleteImage = `-- name: DeleteImage :exec
DELETE FROM images WHERE id = $1::UUID
`

func (q *Queries) DeleteImage(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteImage, id)
	return err
}

const getAllImages = `-- name: GetAllImages :many
SELECT 
    id, 
    title, 
    created_at, 
    updated_at
FROM images
WHERE $1::UUID[] IS NULL OR id = ANY($1::UUID[])
`

type GetAllImagesRow struct {
	ID        uuid.UUID
	Title     null.String
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) GetAllImages(ctx context.Context, imageIds []uuid.UUID) ([]GetAllImagesRow, error) {
	rows, err := q.db.Query(ctx, getAllImages, imageIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllImagesRow
	for rows.Next() {
		var i GetAllImagesRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertOrUpdateImage = `-- name: InsertOrUpdateImage :exec
INSERT INTO images (
    id,
    title,
    file_hash,
    created_at,
    updated_at
) VALUES (
    $1::UUID,
    $2::TEXT,
    $3::TEXT,
    $4::TIMESTAMPTZ,
    $5::TIMESTAMPTZ
) ON CONFLICT (id)
DO UPDATE
SET
    title = EXCLUDED.title,
    file_hash = EXCLUDED.file_hash,
    updated_at = EXCLUDED.updated_at
`

type InsertOrUpdateImageParams struct {
	ID        uuid.UUID
	Title     string
	FileHash  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) InsertOrUpdateImage(ctx context.Context, arg InsertOrUpdateImageParams) error {
	_, err := q.db.Exec(ctx, insertOrUpdateImage,
		arg.ID,
		arg.Title,
		arg.FileHash,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}