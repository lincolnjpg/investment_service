-- +goose Up
CREATE TABLE users_assets (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	user_id UUID NOT NULL,
	asset_id UUID NOT NULL,
	quantity SMALLINT NOT NULL,
	purchase_date DATE NOT NULL,
	status asset_status NOT NULL,
	message TEXT NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
	updated_at TIMESTAMP WITH TIME ZONE NULL,
	PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE users_assets;
