-- +goose Up
CREATE TYPE index_names AS ENUM ('Certificado de Depósito Interbancário', 'Índice Nacional de Preços ao Consumidor Amplo');

-- +goose Down
DROP TYPE index_names;
