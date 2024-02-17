-- +goose Up
CREATE TABLE assets (
	id UUID NOT NULL DEFAULT uuid_generate_v4(),
	name VARCHAR(100) NOT NULL,
	type asset_types NOT NULL,
	unit_value NUMERIC NOT NULL,
	rentability NUMERIC NOT NULL,
	ticker VARCHAR(10),
	due_date DATE,
	PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE assets;
