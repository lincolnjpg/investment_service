package main

import (
	"context"

	"github.com/joho/godotenv"

	"github.com/lincolnjpg/investment_service/internal/adapters/producers"
	"github.com/lincolnjpg/investment_service/internal/adapters/repositories"
	"github.com/lincolnjpg/investment_service/internal/adapters/services"
	"github.com/lincolnjpg/investment_service/internal/config"
	"github.com/lincolnjpg/investment_service/internal/infra"
	"github.com/lincolnjpg/investment_service/internal/ports"
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

	assetIndexesRepository := repositories.NewAssetIndexRepository(db)
	assetIndexesService := services.NewAssetIndexService(assetIndexesRepository)

	assetsRepository := repositories.NewAssetRepository(db)
	assetsService := services.NewAssetService(assetsRepository, assetIndexesService)

	investmentRepository := repositories.NewInvestmentRepository(db)
	producer := producers.RabbitMqProducer{}
	investmentService := services.NewInvestmentService(investmentRepository, producer, userService, assetsService)

	services := struct {
		ports.UserService
		ports.AssetIndexService
		ports.AssetService
		ports.InvestmentService
	}{
		UserService:       userService,
		AssetIndexService: assetIndexesService,
		AssetService:      assetsService,
		InvestmentService: investmentService,
	}

	restApi := NewRestApi(services, envs.RestApiPort)
	restApi.Run()
}
