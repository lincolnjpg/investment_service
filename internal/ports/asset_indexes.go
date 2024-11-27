package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/entities"
)

type AssetIndexService interface {
	CreateAssetIndex(ctx context.Context, input dtos.CreateAssetIndexInput) (dtos.CreateAssetIndexOutput, error)
	GetAssetIndexById(ctx context.Context, input dtos.GetAssetIndexByIdInput) (dtos.GetAssetIndexByIdOutput, error)
	UpdateAssetIndexById(ctx context.Context, input dtos.UpdateAssetIndexByIdInput) (dtos.UpdateAssetIndexByIdOutput, error)
	DeleteAssetIndexById(ctx context.Context, input dtos.DeleteAssetIndexByIdInput) error
}

type AssetIndexRepository interface {
	CreateAssetIndex(ctx context.Context, input dtos.CreateAssetIndexInput) (entities.AssetIndex, error)
	GetAssetIndexById(ctx context.Context, input dtos.GetAssetIndexByIdInput) (entities.AssetIndex, error)
	UpdateAssetIndexById(ctx context.Context, input dtos.UpdateAssetIndexByIdInput) (entities.AssetIndex, error)
	DeleteAssetIndexById(ctx context.Context, input dtos.DeleteAssetIndexByIdInput) error
}
