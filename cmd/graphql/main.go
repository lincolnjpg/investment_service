package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/lincolnjpg/investment_service/cmd/graphql/gqlgen"
	"github.com/lincolnjpg/investment_service/internal/adapters/repositories"
	"github.com/lincolnjpg/investment_service/internal/adapters/services"
	"github.com/lincolnjpg/investment_service/internal/infra"
)

const defaultPort = "8080"

type Envs struct {
	PostgresHost     string
	PostgresUserName string
	PostgresPassword string
	PostgresDatabase string
	PostgresPort     int
	APIPort          string
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

	apiPort := "1212"
	if port := os.Getenv("API_PORT"); port != "" {
		apiPort = port
	}

	return Envs{
		PostgresHost:     postgresHost,
		PostgresUserName: postgresUsername,
		PostgresPassword: postgresPassword,
		PostgresDatabase: postgresDatabase,
		PostgresPort:     postgresPort,
		APIPort:          apiPort,
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	godotenv.Load()

	envs := ReadEnvsFromOS()

	dbConnParams := infra.DBConnParams{
		Host:     envs.PostgresHost,
		UserName: envs.PostgresUserName,
		Password: envs.PostgresPassword,
		Database: envs.PostgresDatabase,
		Port:     envs.PostgresPort,
	}

	ctx := context.Background()

	db, err := infra.NewPostgres(dbConnParams)
	if err != nil {
		panic(err)
	}
	defer db.Close(ctx)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	srv := handler.NewDefaultServer(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: &gqlgen.Resolver{UserService: userService}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
