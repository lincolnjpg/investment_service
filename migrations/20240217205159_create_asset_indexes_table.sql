-- +goose Up
CREATE TABLE asset_indexes (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	name index_names NOT NULL,
	acronym index_acronyms NOT NULL,
	PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE asset_indexes;
