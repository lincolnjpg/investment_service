-- +goose Up
CREATE TABLE assets (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	name VARCHAR(100) NOT NULL,
	asset_type_id UUID NOT NULL,
	asset_index_id UUID NULL,
	unit_price NUMERIC NOT NULL,
	rentability NUMERIC NOT NULL,
	ticker VARCHAR(10),
	due_date DATE,
	PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE assets;
