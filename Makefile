setup:
	go install github.com/pressly/goose/v3/cmd/goose@latest

migration_up:
	goose -dir migrations postgres "host=localhost port=5432 user=postgres password=example database=postgres sslmode=disable" up

migration_status:
	goose -dir migrations postgres "host=localhost port=5432 user=postgres password=example database=postgres sslmode=disable" status

