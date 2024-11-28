-- +goose Up
CREATE TYPE asset_status AS ENUM ('Pending', 'Done', 'Canceled');

-- +goose Down
DROP TYPE asset_status;
