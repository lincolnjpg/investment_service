-- +goose Up
CREATE TYPE index_acronym AS ENUM ('CDI', 'IPCA');

-- +goose Down
DROP TYPE index_acronym;
