package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/entities"
)

type AssetService interface {
	Create(ctx context.Context, input dtos.CreateAssetInput) (dtos.CreateAssetOutput, error)
	GetById(ctx context.Context, input dtos.GetAssetByIdInput) (dtos.GetAssetByIdOutput, error)
	UpdateById(ctx context.Context, input dtos.UpdateAssetByIdInput) (dtos.UpdateAssetByIdOutput, error)
	DeleteById(ctx context.Context, input dtos.DeleteAssetByIdInput) error
}

type AssetRepository interface {
	Create(ctx context.Context, input dtos.CreateAssetInput) (entities.Asset, error)
	GetById(ctx context.Context, input dtos.GetAssetByIdInput) (entities.Asset, error)
	UpdateById(ctx context.Context, input dtos.UpdateAssetByIdInput) (entities.Asset, error)
	DeleteById(ctx context.Context, input dtos.DeleteAssetByIdInput) error
}
