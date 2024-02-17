-- +goose Up
CREATE TYPE asset_types AS ENUM ('fixed', 'variable');

-- +goose Down
DROP TYPE asset_types;
