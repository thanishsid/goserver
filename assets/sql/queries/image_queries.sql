-- name: InsertOrUpdateImage :exec
INSERT INTO images (
    id,
    title,
    file_hash,
    created_at,
    updated_at
) VALUES (
    @id::UUID,
    sqlc.narg('title')::TEXT,
    @file_hash::BYTEA,
    @created_at::TIMESTAMPTZ,
    @updated_at::TIMESTAMPTZ
) ON CONFLICT (id)
DO UPDATE
SET
    title = EXCLUDED.title,
    file_hash = EXCLUDED.file_hash,
    updated_at = EXCLUDED.updated_at;



-- name: DeleteImage :exec
DELETE FROM images WHERE id = @id::UUID;



-- name: CheckImageExists :one
SELECT EXISTS(SELECT 1 FROM images WHERE id = @id::UUID);



-- name: CheckImageHashExists :one
SELECT EXISTS(SELECT 1 FROM images WHERE file_hash = @file_hash::BYTEA);



-- name: GetImageById :one
SELECT 
id, 
title, 
file_hash, 
created_at, 
updated_at
FROM images
WHERE id = @id::UUID;



-- name: GetManyImages :many
SELECT 
    id, 
    title, 
    created_at, 
    updated_at
FROM images
WHERE
sqlc.narg('updated_before')::TIMESTAMPTZ IS NULL OR updated_at < sqlc.narg('updated_before')::TIMESTAMPTZ
ORDER BY updated_at DESC
LIMIT @image_limit::BIGINT;



-- name: GetAllImagesInIDS :many
SELECT 
    id, 
    title, 
    created_at, 
    updated_at
FROM images
WHERE sqlc.narg('image_ids')::UUID[] IS NULL OR id = ANY(sqlc.narg('image_ids')::UUID[]);