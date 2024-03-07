package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddlewares "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/joho/godotenv"
	"github.com/lincolnjpg/investment_service/internal/adapters/repositories"
	"github.com/lincolnjpg/investment_service/internal/adapters/services"
	"github.com/lincolnjpg/investment_service/internal/handlers"
	"github.com/lincolnjpg/investment_service/internal/infra"
)

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

	postgresPassword := "postgres"
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
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	envs := ReadEnvsFromOS()

	dbConnParams := infra.DBConnParams{
		Host:     envs.PostgresHost,
		UserName: envs.PostgresUserName,
		Password: envs.PostgresUserName,
		Database: envs.PostgresDatabase,
		Port:     envs.PostgresPort,
	}

	ctx := context.Background()

	db, err := infra.NewPostgres(dbConnParams)
	if err != nil {
		panic(err)
	}
	defer db.Close(ctx)

	logger := httplog.NewLogger(
		"investment_service",
		httplog.Options{
			JSON:            true,
			LevelFieldName:  "severity",
			LogLevel:        slog.LevelDebug,
			Concise:         true,
			RequestHeaders:  true,
			TimeFieldFormat: time.RFC3339Nano,
		},
	)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	router := chi.NewRouter()
	router.Use(httplog.RequestLogger(logger))
	router.Use(chimiddlewares.Heartbeat("/ping"))

	usersRouter := chi.NewRouter()
	usersRouter.Post("/", handlers.CreateUserHandler(userService))
	usersRouter.Get("/{id}", handlers.GetUserByIDHandler(userService))
	router.Mount("/users", usersRouter)

	http.ListenAndServe(fmt.Sprintf(":%s", envs.APIPort), router)
}
