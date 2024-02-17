-- +goose Up
CREATE TABLE types (
	id UUID NOT NULL DEFAULT uuid_generate_v4(),
	name VARCHAR(100) NOT NULL,
	description VARCHAR NOT NULL,
	PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE types;
