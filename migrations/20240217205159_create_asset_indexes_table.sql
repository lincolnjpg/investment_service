-- +goose Up
CREATE TABLE asset_indexes (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	name VARCHAR NOT NULL,
	acronym VARCHAR NOT NULL,
	PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE asset_indexes;
