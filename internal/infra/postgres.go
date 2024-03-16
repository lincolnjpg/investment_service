package infra

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type DBConnParams struct {
	Host     string
	UserName string
	Password string
	Database string
	Port     int
}

func NewPostgres(params DBConnParams) (*pgx.Conn, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		params.UserName,
		params.Password,
		params.Host,
		params.Port,
		params.Database,
	)

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn, err
}
