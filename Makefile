setup:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/onsi/ginkgo/v2/ginkgo
	go get github.com/onsi/gomega/...

graphql:
	go run github.com/99designs/gqlgen generate

run_rest:
	go run ./cmd/rest/main.go

run_graphql:
	go run ./cmd/graphql/main.go

migrations_dev_up:
	goose -dir migrations postgres "host=localhost port=5432 user=postgres password=example database=postgres sslmode=disable" up

migrations_dev_status:
	goose -dir migrations postgres "host=localhost port=5432 user=postgres password=example database=postgres sslmode=disable" status

migrations_test_up:
	goose -dir migrations postgres "host=localhost port=5433 user=postgres password=example database=postgres sslmode=disable" up

migrations_test_status:
	goose -dir migrations postgres "host=localhost port=5433 user=postgres password=example database=postgres sslmode=disable" status

migrations_dev_create:
	goose -dir migrations create $(name) sql