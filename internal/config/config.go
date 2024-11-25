package config

import (
	"os"
	"strconv"
)

type Envs struct {
	PostgresHost     string
	PostgresUserName string
	PostgresPassword string
	PostgresDatabase string
	PostgresPort     int
	RestApiPort      string
	GraphQlPort      string
}

func ReadEnvsFromOS() Envs {
	postgresHost := "localhost"
	if host := os.Getenv("POSTGRES_HOST"); host != "" {
		postgresHost = host
	}

	postgresUsername := "postgres"
	if username := os.Getenv("POSTGRES_USERNAME"); username != "" {
		postgresUsername = username
	}

	postgresPassword := "example"
	if password := os.Getenv("POSTGRES_PASSWORD"); password != "" {
		postgresPassword = password
	}

	postgresDatabase := "postgres"
	if database := os.Getenv("POSTGRES_DATABASE"); database != "" {
		postgresDatabase = database
	}

	postgresPort := 5432
	if port := os.Getenv("POSTGRES_PORT"); port != "" {
		if value, err := strconv.Atoi(port); err == nil {
			postgresPort = value
		}
	}

	restApiPort := "1212"
	if port := os.Getenv("REST_API_PORT"); port != "" {
		restApiPort = port
	}

	graphQlPort := "8080"
	if port := os.Getenv("GRAPHQL_PORT"); port != "" {
		graphQlPort = port
	}

	return Envs{
		PostgresHost:     postgresHost,
		PostgresUserName: postgresUsername,
		PostgresPassword: postgresPassword,
		PostgresDatabase: postgresDatabase,
		PostgresPort:     postgresPort,
		RestApiPort:      restApiPort,
		GraphQlPort:      graphQlPort,
	}
}
