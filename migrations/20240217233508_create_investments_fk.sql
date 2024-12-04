-- +goose Up
ALTER TABLE investments
ADD CONSTRAINT users_fk FOREIGN KEY(user_id)
REFERENCES users (id)
ON UPDATE CASCADE
ON DELETE CASCADE;

ALTER TABLE investments
ADD CONSTRAINT assets_fk FOREIGN KEY(asset_id)
REFERENCES assets (id)
ON UPDATE CASCADE
ON DELETE CASCADE;

-- +goose Down
ALTER TABLE investments
DROP CONSTRAINT users_fk;

ALTER TABLE investments
DROP CONSTRAINT assets_fk;

