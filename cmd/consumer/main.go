package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/lincolnjpg/investment_service/internal/adapters/repositories"
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

	investmentRepository := repositories.NewInvestmentRepository(db)

	consumer := NewRabbitMqConsumer("https://query1.finance.yahoo.com/v8/finance/chart", investmentRepository)
	consumer.consume()
}
