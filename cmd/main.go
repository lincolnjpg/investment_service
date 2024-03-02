package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func ReadEnvFromOS() Envs {
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envs := ReadEnvFromOS()

	dbConnParams := infra.DBConnParams{
		Host:     envs.PostgresHost,
		UserName: envs.PostgresUserName,
		Password: envs.PostgresUserName,
		Database: envs.PostgresDatabase,
		Port:     envs.PostgresPort,
	}

	db, err := infra.NewPostgres(dbConnParams)
	if err != nil {
		panic(err)
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/ping"))

	usersRouter := chi.NewRouter()
	usersRouter.Post("/", handlers.CreateUserHandle(userService))
	r.Mount("/users", usersRouter)

	http.ListenAndServe(fmt.Sprintf(":%s", envs.APIPort), r)
}
