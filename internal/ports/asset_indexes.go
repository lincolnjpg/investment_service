package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/entities"
)

type AssetIndexService interface {
	Create(ctx context.Context, input dtos.CreateAssetIndexInput) (dtos.CreateAssetIndexOutput, error)
	GetById(ctx context.Context, input dtos.GetAssetIndexByIdInput) (dtos.GetAssetIndexByIdOutput, error)
	UpdateById(ctx context.Context, input dtos.UpdateAssetIndexByIdInput) (dtos.UpdateAssetIndexByIdOutput, error)
	DeleteById(ctx context.Context, input dtos.DeleteAssetIndexByIdInput) error
}

type AssetIndexRepository interface {
	Create(ctx context.Context, input dtos.CreateAssetIndexInput) (entities.AssetIndex, error)
	GetById(ctx context.Context, input dtos.GetAssetIndexByIdInput) (entities.AssetIndex, error)
	UpdateById(ctx context.Context, input dtos.UpdateAssetIndexByIdInput) (entities.AssetIndex, error)
	DeleteById(ctx context.Context, input dtos.DeleteAssetIndexByIdInput) error
}
