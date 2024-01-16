-- +goose Up
CREATE TYPE investor_profiles AS ENUM ('conservador', 'moderado', 'arrojado');

-- +goose Down
DROP TYPE investor_profiles;
