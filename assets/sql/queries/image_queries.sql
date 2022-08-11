-- name: InsertOrUpdateImage :exec
INSERT INTO images (
    id,
    file_name,
    created_at,
    updated_at
) VALUES (
    @id::UUID,
    @file_name::TEXT,
    @created_at::TIMESTAMPTZ,
    @updated_at::TIMESTAMPTZ
) ON CONFLICT (id)
DO UPDATE
SET
    file_name = EXCLUDED.file_name,
    updated_at = EXCLUDED.updated_at;



-- name: DeleteImage :exec
DELETE FROM images WHERE id = @id::UUID;



-- name: CheckImageExists :one
SELECT EXISTS(SELECT 1 FROM images WHERE id = @id::UUID);



-- name: CheckImageFileExists :one
SELECT EXISTS(SELECT 1 FROM images WHERE file_name = @file_name::TEXT);



-- name: GetImageById :one
SELECT 
id, 
file_name,
created_at, 
updated_at
FROM images
WHERE id = @id::UUID;



-- name: GetManyImages :many
SELECT 
    id,
    file_name,
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
    file_name,
    created_at, 
    updated_at
FROM images
WHERE sqlc.narg('image_ids')::UUID[] IS NULL OR id = ANY(sqlc.narg('image_ids')::UUID[]);