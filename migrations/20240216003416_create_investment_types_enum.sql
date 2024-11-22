-- +goose Up
CREATE TYPE investment_types AS ENUM ('CDB', 'LCI', 'LCA', 'CRI', 'CRA', 'Debênture', 'Tesouro Direto', 'Ação', 'FII');

-- +goose Down
DROP TYPE investment_types;
