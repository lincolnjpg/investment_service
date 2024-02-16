-- +goose Up
create extension "uuid-ossp";

-- +goose Down
drop extension "uuid-ossp";
