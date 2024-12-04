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
	port string
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
	usersRouter.Post("/", handlers.CreateUserHandler(r.app))
	usersRouter.Get("/{id}", handlers.GetUserByIDHandler(r.app))
	usersRouter.Put("/{id}", handlers.UpateUserByIdHandler(r.app))
	usersRouter.Delete("/{id}", handlers.DeleteUserByIDHandler(r.app))

	assetIndexesRouter := chi.NewRouter()
	assetIndexesRouter.Post("/", handlers.CreateAssetIndexHandler(r.app))
	assetIndexesRouter.Get("/{id}", handlers.GetAssetIndexByIdHandler(r.app))
	assetIndexesRouter.Put("/{id}", handlers.UpdateAssetIndexByIdHandler(r.app))
	assetIndexesRouter.Delete("/{id}", handlers.DeleteAssetIndexByIDHandler(r.app))

	assetsRouter := chi.NewRouter()
	assetsRouter.Post("/", handlers.CreateAssetHandler(r.app))
	assetsRouter.Get("/{id}", handlers.GetAssetByIdHandler(r.app))
	assetsRouter.Put("/{id}", handlers.UpdateAssetByIdHandler(r.app))
	assetsRouter.Delete("/{id}", handlers.DeleteAssetByIdHandler(r.app))

	investmentRouter := chi.NewRouter()
	investmentRouter.Post("/", handlers.CreateInvestmentHandler(r.app))

	router.Mount("/users", usersRouter)
	router.Mount("/indexes", assetIndexesRouter)
	router.Mount("/assets", assetsRouter)
	router.Mount("/user-assets", investmentRouter)
	router.Mount("/debug", chimiddlewares.Profiler())

	http.ListenAndServe(fmt.Sprintf(":%s", r.port), router)
}

func NewRestApi(app ports.Application, port string) *restApi {
	return &restApi{
		app:  app,
		port: port,
	}
}
