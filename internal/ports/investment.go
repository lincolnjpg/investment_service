package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/entities"
)

type InvestmentService interface {
	CreateInvestment(ctx context.Context, input dtos.CreateInvestmentInput) (dtos.CreateInvestmentOutput, error)
	GetInvestmentById(ctx context.Context, input dtos.GetInvestmentByIdInput) (dtos.GetInvestmentByIdOutput, error)
	UpdateInvestmentById(ctx context.Context, input dtos.UpdateInvestmentByIdInput) (dtos.UpdateInvestmentByIdOutput, error)
	DeleteInvestmentById(ctx context.Context, input dtos.DeleteInvestmentByIdInput) error
}

type InvestmentRepository interface {
	CreateInvestment(ctx context.Context, input dtos.CreateInvestmentInput) (entities.Investment, error)
	GetInvestmentById(ctx context.Context, input dtos.GetInvestmentByIdInput) (entities.Investment, error)
	UpdateInvestmentById(ctx context.Context, input dtos.UpdateInvestmentByIdInput) (entities.Investment, error)
	DeleteInvestmentById(ctx context.Context, input dtos.DeleteInvestmentByIdInput) error
}
