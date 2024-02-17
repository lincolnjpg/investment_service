-- +goose Up
CREATE TYPE investor_profiles AS ENUM ('conservative', 'moderate', 'aggressive');

-- +goose Down
DROP TYPE investor_profiles;
