-- name: InsertOrUpdateRoles :exec
INSERT INTO roles (
    id, name
) VALUES (
    @id::TEXT,
    @name::TEXT
) 
ON CONFLICT (id)
DO
UPDATE SET
    name = EXCLUDED.name;


-- name: GetAllRoles :many
SELECT id, name FROM roles;


-- name: DeleteRole :exec
DELETE FROM roles WHERE id = @id::TEXT;
