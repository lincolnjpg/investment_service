package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/entities"
)

type AssetService interface {
	CreateAsset(ctx context.Context, input dtos.CreateAssetInput) (dtos.CreateAssetOutput, error)
	GetAssetById(ctx context.Context, input dtos.GetAssetByIdInput) (dtos.GetAssetByIdOutput, error)
	UpdateAssetById(ctx context.Context, input dtos.UpdateAssetByIdInput) (dtos.UpdateAssetByIdOutput, error)
	DeleteAssetById(ctx context.Context, input dtos.DeleteAssetByIdInput) error
}

type AssetRepository interface {
	CreateAsset(ctx context.Context, input dtos.CreateAssetInput) (entities.Asset, error)
	GetAssetById(ctx context.Context, input dtos.GetAssetByIdInput) (entities.Asset, error)
	UpdateAssetById(ctx context.Context, input dtos.UpdateAssetByIdInput) (entities.Asset, error)
	DeleteAssetById(ctx context.Context, input dtos.DeleteAssetByIdInput) error
}
