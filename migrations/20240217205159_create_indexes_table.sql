-- +goose Up
CREATE TABLE indexes (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	name index_name NOT NULL,
	acronym index_acronym NOT NULL,
	PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE indexes;
