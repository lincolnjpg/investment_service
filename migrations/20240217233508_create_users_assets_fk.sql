-- +goose Up
ALTER TABLE users_assets ADD CONSTRAINT users_fk FOREIGN KEY(user_id) REFERENCES users (id);
ALTER TABLE users_assets ADD CONSTRAINT assets_fk FOREIGN KEY(asset_id) REFERENCES assets (id);

-- +goose Down
ALTER TABLE users_assets DROP CONSTRAINT users_fk;
ALTER TABLE users_assets DROP CONSTRAINT assets_fk;

