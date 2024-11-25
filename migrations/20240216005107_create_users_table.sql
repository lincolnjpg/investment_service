-- +goose Up
CREATE TABLE users (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	name VARCHAR NOT NULL,
	investor_profile investor_profiles NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
	updated_at TIMESTAMP WITH TIME ZONE NULL,
	PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE users;
