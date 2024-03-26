-- +goose Up
CREATE TYPE asset_classes AS ENUM ('FIXED_INCOME', 'VARIABLE_INCOME');

-- +goose Down
DROP TYPE asset_classes;
