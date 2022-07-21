-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    full_name TEXT NOT NULL,
    role_id INTEGER NOT NULL,
    password_hash TEXT NOT NULL,
    picture_id UUID REFERENCES images (id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);
CREATE INDEX idx_user_created_at ON users (created_at);
CREATE INDEX idx_user_updated_at ON users (updated_at);
CREATE INDEX idx_user_role ON users USING HASH (user_role);
CREATE INDEX idx_user_username ON users USING HASH (username);
CREATE INDEX idx_user_deleted_at ON users (deleted_at) WHERE deleted_at IS NOT NULL;
-- +goose StatementEnd


-- +goose Down
DROP TABLE users;
