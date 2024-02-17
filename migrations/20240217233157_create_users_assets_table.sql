-- +goose Up
CREATE TABLE users_assets (
	id UUID NOT NULL DEFAULT uuid_generate_v4(),
	user_id UUID NOT NULL,
	asset_id UUID NOT NULL,
	quantity SMALLINT NOT NULL,
	purchase_date DATE NOT NULL,
	PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE users_assets;
