-- +goose Up
CREATE TYPE index_acronyms AS ENUM ('CDI', 'IPCA');

-- +goose Down
DROP TYPE index_acronyms;
