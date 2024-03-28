-- +goose Up
CREATE TABLE asset_types (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	name investment_types NOT NULL,
	description VARCHAR NOT NULL,
	class asset_classes NOT NULL,
	PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE asset_types;
