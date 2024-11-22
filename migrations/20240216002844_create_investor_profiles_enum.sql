-- +goose Up
CREATE TYPE investor_profiles AS ENUM ('Conservador', 'Moderado', 'Arrojado');

-- +goose Down
DROP TYPE investor_profiles;
