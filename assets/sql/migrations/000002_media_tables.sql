-- +goose Up
-- +goose StatementBegin
CREATE TABLE images (
    id UUID PRIMARY KEY,
    file_name TEXT UNIQUE NOT NULL,
    file_size BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);


CREATE TABLE videos (
    id UUID PRIMARY KEY,
    file_name TEXT UNIQUE NOT NULL,
    thumbnail_id UUID NOT NULL REFERENCES images (id),
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd


-- +goose Down
DROP TABLE images;
DROP TABLE videos;
