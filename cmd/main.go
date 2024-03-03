package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	router := mux.NewRouter()
	usersRouter := router.PathPrefix("/users").Subrouter() //use the path prefix "/users{_:/?}"" to match /users and /users/
	usersRouter.HandleFunc("", handlers.CreateUserHandle(ctx, userService)).Methods("POST")

	http.ListenAndServe(fmt.Sprintf(":%s", envs.APIPort), removeTrailingSlash(router))
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
