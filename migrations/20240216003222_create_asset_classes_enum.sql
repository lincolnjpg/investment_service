-- +goose Up
CREATE TYPE asset_classes AS ENUM ('Renda Fixa', 'Renda Variável');

-- +goose Down
DROP TYPE asset_classes;
