-- +goose Up
CREATE TABLE indexes (
	id UUID NOT NULL DEFAULT uuid_generate_v4(),
	name VARCHAR(100) NOT NULL,
	acronym VARCHAR(10) NOT NULL,
	PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE indexes;
