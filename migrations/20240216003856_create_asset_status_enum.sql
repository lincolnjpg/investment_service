-- +goose Up
CREATE TYPE investment_status AS ENUM ('Pending', 'Done', 'Canceled');

-- +goose Down
DROP TYPE investment_status;
