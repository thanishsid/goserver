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
    @file_hash::TEXT,
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
SELECT EXISTS(SELECT 1 FROM images WHERE file_hash = @file_hash::TEXT);



-- name: GetImageById :one
SELECT id, title, file_hash, created_at, updated_at FROM images WHERE id = @id::UUID;



-- name: GetAllImages :many
SELECT 
    id, 
    title, 
    created_at, 
    updated_at
FROM images
WHERE sqlc.narg('image_ids')::UUID[] IS NULL OR id = ANY(sqlc.narg('image_ids')::UUID[]);



-- name: GetManyImages :many
SELECT 
    i.id, 
    i.title, 
    i.created_at, 
    i.updated_at
FROM images i
LEFT JOIN users u ON u.picture_id = i.id
WHERE 
@view_unused::BOOLEAN = FALSE OR (u.picture_id IS NULL) AND
sqlc.narg('updated_after')::TIMESTAMPTZ IS NULL OR i.updated_at > sqlc.narg('updated_after')::TIMESTAMPTZ
ORDER BY i.updated_at DESC
LIMIT @image_limit::BIGINT;



