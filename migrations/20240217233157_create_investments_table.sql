-- +goose Up
CREATE TABLE investments (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	user_id UUID NOT NULL,
	asset_id UUID NOT NULL,
	quantity SMALLINT NOT NULL,
	purchase_date DATE NOT NULL,
	status investment_status NOT NULL,
	message TEXT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
	updated_at TIMESTAMP WITH TIME ZONE NULL,
	PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE investments;
