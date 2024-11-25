package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddlewares "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/joho/godotenv"

	"github.com/lincolnjpg/investment_service/cmd/rest/handlers"
	"github.com/lincolnjpg/investment_service/internal/adapters/repositories"
	"github.com/lincolnjpg/investment_service/internal/adapters/services"
	"github.com/lincolnjpg/investment_service/internal/config"
	"github.com/lincolnjpg/investment_service/internal/infra"
)

func main() {
	godotenv.Load()

	envs := config.ReadEnvsFromOS()

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

	assetIndexesRepository := repositories.NewAssetIndexRepository(db)
	assetIndexesService := services.NewAssetIndexService(assetIndexesRepository)

	assetsRepository := repositories.NewAssetRepository(db)
	assetsService := services.NewAssetService(assetsRepository, assetIndexesService)

	router := chi.NewRouter()
	router.Use(httplog.RequestLogger(logger))
	router.Use(chimiddlewares.Heartbeat("/ping"))

	usersRouter := chi.NewRouter()
	usersRouter.Post("/", handlers.CreateUserHandler(userService))
	usersRouter.Get("/{id}", handlers.GetUserByIDHandler(userService))
	usersRouter.Put("/{id}", handlers.UpateUserByIdHandler(userService))
	usersRouter.Delete("/{id}", handlers.DeleteUserByIDHandler(userService))

	assetIndexesRouter := chi.NewRouter()
	assetIndexesRouter.Post("/", handlers.CreateAssetIndexHandler(assetIndexesService))
	assetIndexesRouter.Get("/{id}", handlers.GetAssetIndexByIdHandler(assetIndexesService))
	assetIndexesRouter.Put("/{id}", handlers.UpdateAssetIndexByIdHandler(assetIndexesService))
	assetIndexesRouter.Delete("/{id}", handlers.DeleteAssetIndexByIDHandler(assetIndexesService))

	assetsRouter := chi.NewRouter()
	assetsRouter.Post("/", handlers.CreateAssetHandler(assetsService))
	assetsRouter.Get("/{id}", handlers.GetAssetByIdHandler(assetsService))
	assetsRouter.Put("/{id}", handlers.UpdateAssetByIdHandler(assetsService))
	assetsRouter.Delete("/{id}", handlers.DeleteAssetByIdHandler(assetsService))

	router.Mount("/users", usersRouter)
	router.Mount("/indexes", assetIndexesRouter)
	router.Mount("/assets", assetsRouter)
	router.Mount("/debug", chimiddlewares.Profiler())

	http.ListenAndServe(fmt.Sprintf(":%s", envs.RestApiPort), router)
}
