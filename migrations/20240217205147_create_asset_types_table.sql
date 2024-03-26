-- +goose Up
CREATE TABLE asset_types (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	name investment_types NOT NULL,
	description VARCHAR NOT NULL,
	class asset_classes NOT NULL,
	index_id UUID NULL,
	PRIMARY KEY (id),
	UNIQUE (name, index_id)
);

-- +goose Down
DROP TABLE asset_types;
