-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE users (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    full_name TEXT NOT NULL,
    role TEXT NOT NULL REFERENCES roles (id),
    password_hash TEXT,
    picture_id UUID UNIQUE REFERENCES images (id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,

    search_index TSVECTOR GENERATED ALWAYS AS (
       setweight(to_tsvector('simple', email), 'A')
        || setweight(to_tsvector('english', full_name), 'B')
            || setweight(to_tsvector('english', username), 'C')
    ) STORED
);
CREATE INDEX idx_user_created_at ON users (created_at);
CREATE INDEX idx_user_updated_at ON users (updated_at);
CREATE INDEX idx_user_email ON users USING HASH (email);
CREATE INDEX idx_user_role ON users USING HASH (role);
CREATE INDEX idx_user_deleted_at ON users (deleted_at) WHERE deleted_at IS NOT NULL;
CREATE INDEX idx_user_search_index ON users USING GIN(search_index);
-- +goose StatementEnd


-- +goose Down
DROP TABLE users;
DROP TABLE roles;
