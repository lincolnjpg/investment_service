-- +goose Up
CREATE TYPE asset_types AS ENUM ('CDB', 'LCI', 'LCA', 'CRI', 'CRA', 'Debênture', 'Tesouro Direto', 'Ação', 'FII');

-- +goose Down
DROP TYPE asset_types;
