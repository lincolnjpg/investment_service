package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/lincolnjpg/investment_service/cmd/graphql/gqlgen"
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

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	resolver := gqlgen.Resolver{
		UserService: userService,
	}
	graphQl := NewGraphQl(resolver, "8080")
	graphQl.Run()
}
