-- name: InsertOrUpdateUser :exec
INSERT INTO users (
  id, 
  username, 
  email, 
  full_name, 
  user_role,
  password_hash,
  picture_id,
  created_at, 
  updated_at
) VALUES (
  @id::UUID,
  @username::TEXT,
  @email::TEXT,
  @full_name::TEXT,
  @user_role::TEXT,
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
  user_role = EXCLUDED.role_id,
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
user_role,
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
user_role,
password_hash,
picture_id,
created_at,
updated_at,
deleted_at
FROM users 
WHERE
email = @email::TEXT
LIMIT 1;



-- name: GetAllUsers :many
SELECT 
id, 
username, 
email, 
full_name,
user_role,
picture_id,
created_at,
updated_at,
deleted_at
FROM users
WHERE sqlc.narg('user_ids')::UUID[] IS NULL OR id = ANY(sqlc.narg('user_ids')::UUID[]);



-- name: DeleteAllUsers :exec
DELETE FROM users;

