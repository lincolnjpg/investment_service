package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lincolnjpg/investment_service/internal/handlers"
	infra "github.com/lincolnjpg/investment_service/internal/infra/postgres"
	"github.com/lincolnjpg/investment_service/internal/repositories"
	"github.com/lincolnjpg/investment_service/internal/services"
)

func main() {
	db, err := infra.ConnectToDB("postgres://postgres:example@localhost:5432/postgres")
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
