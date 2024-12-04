-- +goose Up
CREATE TABLE assets (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	name VARCHAR(100) NOT NULL,
	type asset_types NOT NULL,
	asset_index_id UUID NULL,
	unit_price NUMERIC NOT NULL,
	rentability NUMERIC NOT NULL,
	ticker VARCHAR(10),
	due_date DATE,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
	updated_at TIMESTAMP WITH TIME ZONE NULL,
	PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE assets;
