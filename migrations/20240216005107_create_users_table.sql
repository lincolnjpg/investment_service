-- +goose Up
CREATE TABLE users (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	name VARCHAR NOT NULL,
	investor_profile investor_profiles NOT NULL,
	PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE users;
