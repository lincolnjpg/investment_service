-- +goose Up
CREATE TABLE types (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	name VARCHAR(100) NOT NULL,
	description VARCHAR NOT NULL,
	PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE types;
