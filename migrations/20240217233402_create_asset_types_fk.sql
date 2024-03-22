-- +goose Up
ALTER TABLE asset_types
ADD CONSTRAINT indexes_fk FOREIGN KEY(index_id)
REFERENCES indexes (id)
ON UPDATE CASCADE
ON DELETE CASCADE;

-- +goose Down
ALTER TABLE asset_types
DROP CONSTRAINT indexes_fk;

