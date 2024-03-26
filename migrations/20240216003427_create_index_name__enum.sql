-- +goose Up
CREATE TYPE index_name AS ENUM ('Certificado de Depósito Interbancário', 'Índice Nacional de Preços ao Consumidor Amplo');

-- +goose Down
DROP TYPE index_name;
