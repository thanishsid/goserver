-- name: InsertOrUpdateUser :exec
INSERT INTO users (
  id, 
  username, 
  email, 
  full_name, 
  role,
  password_hash,
  picture_id,
  created_at, 
  updated_at
) VALUES (
  @id::UUID,
  @username::TEXT,
  @email::TEXT,
  @full_name::TEXT,
  @role::TEXT,
  @password_hash::TEXT,
  sqlc.narg('picture_id')::UUID,
  @created_at::TIMESTAMPTZ,
  @updated_at::TIMESTAMPTZ
)
ON CONFLICT (id)
DO 
UPDATE SET 
  username = EXCLUDED.username,
  full_name = EXCLUDED.full_name,
  role = EXCLUDED.role,
  password_hash = EXCLUDED.password_hash,
  picture_id = EXCLUDED.picture_id,
  updated_at = EXCLUDED.updated_at;



-- name: SoftDeleteUser :exec
UPDATE users
SET
deleted_at = NOW()
WHERE id = @user_id::UUID;



-- name: HardDeleteUser :exec
DELETE FROM users WHERE id = @user_id::UUID;



-- name: GetUserById :one
SELECT 
id, 
username, 
email, 
full_name,
role,
password_hash,
picture_id,
created_at,
updated_at,
deleted_at
FROM users
WHERE
id = @user_id::UUID
LIMIT 1;



-- name: GetUserByEmail :one
SELECT 
id, 
username, 
email, 
full_name,
role,
password_hash,
picture_id,
created_at,
updated_at,
deleted_at
FROM users
WHERE
email = @email::TEXT
LIMIT 1;



-- name: GetManyUsers :many
SELECT
id,
username,
email,
full_name,
role,
picture_id,
created_at,
updated_at,
deleted_at
FROM 
  users, 
  to_tsquery('simple', array_to_string(string_to_array(sqlc.narg('search')::TEXT, ' '), '&')) query,
  SIMILARITY(sqlc.narg('search')::TEXT, email || full_name || username) sm,
  GREATEST(ts_rank(search_index, query), sm) ranking
WHERE
  (query IS NULL OR (query @@ search_index) OR sm > 0.1) AND
  (sqlc.narg('role')::TEXT IS NULL OR role = sqlc.narg('role')::TEXT) AND
  (sqlc.narg('updated_before')::TIMESTAMPTZ IS NULL OR updated_at < sqlc.narg('updated_before')::TIMESTAMPTZ) AND
  (@show_deleted::BOOLEAN OR deleted_at IS NULL)
ORDER BY 
ranking DESC,
updated_at DESC
LIMIT @users_limit::BIGINT;



-- name: GetAllUsersInIDS :many
SELECT 
id, 
username, 
email, 
full_name,
role,
picture_id,
created_at,
updated_at,
deleted_at
FROM users
WHERE sqlc.narg('user_ids')::UUID[] IS NULL OR id = ANY(sqlc.narg('user_ids')::UUID[]);
