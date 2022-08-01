-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION pg_trgm;
-- +goose StatementEnd


-- +goose Down
DROP EXTENSION pg_tgrm;
