package infra

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectToDB(connString string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	return conn, err
}
