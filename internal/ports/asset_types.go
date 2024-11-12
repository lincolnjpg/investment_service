package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/entities"
)

type AssetTypeService interface {
	Create(ctx context.Context, input dtos.CreateAssetTypeInput) (dtos.CreateAssetTypeOutput, error)
	GetById(ctx context.Context, input dtos.GetAssetTypeByIDInput) (dtos.GetAssetTypeByIDOutput, error)
	UpdateById(ctx context.Context, input dtos.UpdateAssetTypeByIdInput) (dtos.UpdateAssetTypeByIdOutput, error)
	DeleteById(ctx context.Context, input dtos.DeleteAssetTypeByIdInput) error
}

type AssetTypeRepository interface {
	Create(ctx context.Context, input dtos.CreateAssetTypeInput) (entities.AssetType, error)
	GetById(ctx context.Context, input dtos.GetAssetTypeByIDInput) (entities.AssetType, error)
	UpdateById(ctx context.Context, input dtos.UpdateAssetTypeByIdInput) (entities.AssetType, error)
	DeleteById(ctx context.Context, input dtos.DeleteAssetTypeByIdInput) error
}
