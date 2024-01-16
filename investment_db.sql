CREATE TYPE investor_profiles AS ENUM ('conservador', 'moderado', 'arrojado');

CREATE TABLE users (
	id UUID NOT NULL DEFAULT uuid_generate_v4(),
	name VARCHAR NOT NULL,
	investor_profile investor_profiles NOT NULL,
	PRIMARY KEY (id)
);