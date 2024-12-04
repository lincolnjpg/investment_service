package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/entities"
)

type InvestmentService interface {
	CreateInvestment(ctx context.Context, input dtos.CreateInvestmentInput) (dtos.CreateInvestmentOutput, error)
}

type InvestmentRepository interface {
	CreateInvestment(ctx context.Context, input dtos.CreateInvestmentInput) (entities.Investment, error)
}
