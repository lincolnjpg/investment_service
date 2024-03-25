-- +goose Up
CREATE TYPE investment_types AS ENUM ('CDB', 'LCI', 'LCA', 'CRI', 'CRA', 'TESOURO DIRETO', 'AÇÃO', 'FII');

-- +goose Down
DROP TYPE investment_types;
