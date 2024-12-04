-- +goose Up
CREATE TYPE investment_types AS ENUM ('Buy', 'Sell');

-- +goose Down
DROP TYPE investment_types;
