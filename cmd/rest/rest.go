package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddlewares "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/lincolnjpg/investment_service/cmd/rest/handlers"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type restApi struct {
	app  ports.Application
	port int
}

func (r restApi) Run() {
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

	router := chi.NewRouter()
	router.Use(httplog.RequestLogger(logger))
	router.Use(chimiddlewares.Heartbeat("/ping"))

	usersRouter := chi.NewRouter()
	usersRouter.Post("/", handlers.CreateUserHandler(r.app.UserService))
	usersRouter.Get("/{id}", handlers.GetUserByIDHandler(r.app.UserService))
	usersRouter.Put("/{id}", handlers.UpateUserByIdHandler(r.app.UserService))
	usersRouter.Delete("/{id}", handlers.DeleteUserByIDHandler(r.app.UserService))

	assetIndexesRouter := chi.NewRouter()
	assetIndexesRouter.Post("/", handlers.CreateAssetIndexHandler(r.app.AssetIndexService))
	assetIndexesRouter.Get("/{id}", handlers.GetAssetIndexByIdHandler(r.app.AssetIndexService))
	assetIndexesRouter.Put("/{id}", handlers.UpdateAssetIndexByIdHandler(r.app.AssetIndexService))
	assetIndexesRouter.Delete("/{id}", handlers.DeleteAssetIndexByIDHandler(r.app.AssetIndexService))

	assetsRouter := chi.NewRouter()
	assetsRouter.Post("/", handlers.CreateAssetHandler(r.app.AssetService))
	assetsRouter.Get("/{id}", handlers.GetAssetByIdHandler(r.app.AssetService))
	assetsRouter.Put("/{id}", handlers.UpdateAssetByIdHandler(r.app.AssetService))
	assetsRouter.Delete("/{id}", handlers.DeleteAssetByIdHandler(r.app.AssetService))

	router.Mount("/users", usersRouter)
	router.Mount("/indexes", assetIndexesRouter)
	router.Mount("/assets", assetsRouter)
	router.Mount("/debug", chimiddlewares.Profiler())

	http.ListenAndServe(fmt.Sprintf(":%d", r.port), router)
}

func NewRestApi(app ports.Application, port int) *restApi {
	return &restApi{
		app:  app,
		port: port,
	}
}
