-- +goose Up
ALTER TABLE assets
ADD CONSTRAINT asset_indexes_fk FOREIGN KEY(asset_index_id)
REFERENCES asset_indexes (id)
ON UPDATE CASCADE
ON DELETE CASCADE;

-- +goose Down
ALTER TABLE assets
DROP CONSTRAINT asset_indexes_fk;
