package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/lincolnjpg/investment_service/internal/handlers"
	infra "github.com/lincolnjpg/investment_service/internal/infra/postgres"
	"github.com/lincolnjpg/investment_service/internal/repositories"
	"github.com/lincolnjpg/investment_service/internal/services"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
	)

	db, err := infra.ConnectToDB(connString)
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

	http.ListenAndServe(":3000", r)
}
