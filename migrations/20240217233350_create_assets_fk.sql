-- +goose Up
ALTER TABLE assets
ADD CONSTRAINT asset_types_fk FOREIGN KEY(asset_type_id)
REFERENCES asset_types (id)
ON UPDATE CASCADE
ON DELETE CASCADE;

-- +goose Down
ALTER TABLE assets
DROP CONSTRAINT asset_types_fk;

