-- name: InsertOrUpdateVideo :exec
INSERT INTO videos (
    id,
    file_name,
    thumbnail_id,
    created_at,
    updated_at
) VALUES (
    @id::UUID,
    @file_name::TEXT,
    @thumbnail_id::UUID,
    @created_at::TIMESTAMPTZ,
    @updated_at::TIMESTAMPTZ
) ON CONFLICT (id)
DO UPDATE
SET
    file_name = EXCLUDED.file_name,
    thumbnail_id = EXCLUDED.thumbnail_id,
    updated_at = EXCLUDED.updated_at;


-- name: DeleteVideo :exec
DELETE FROM videos WHERE id = @id::UUID;



-- name: CheckVideoExists :one
SELECT EXISTS(SELECT 1 FROM videos WHERE id = @id::UUID);



-- name: CheckVideoFileExists :one
SELECT EXISTS(SELECT 1 FROM videos WHERE file_name = @file_name::TEXT);


-- name: GetVideoById :one
SELECT 
id, 
file_name, 
thumbnail_id,
created_at, 
updated_at
FROM videos
WHERE id = @id::UUID;



-- name: GetManyVideos :many
SELECT 
    id, 
    file_name,
    thumbnail_id,
    created_at, 
    updated_at
FROM videos
WHERE
sqlc.narg('updated_before')::TIMESTAMPTZ IS NULL OR updated_at < sqlc.narg('updated_before')::TIMESTAMPTZ
ORDER BY updated_at DESC
LIMIT @image_limit::BIGINT;



-- name: GetAllVideosInIDS :many
SELECT 
    id, 
    file_name, 
    thumbnail_id,
    created_at, 
    updated_at
FROM videos
WHERE sqlc.narg('video_ids')::UUID[] IS NULL OR id = ANY(sqlc.narg('video_ids')::UUID[]);