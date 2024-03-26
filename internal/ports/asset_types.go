package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/domain"
)

type AssetTypeService interface {
	Create(ctx context.Context, input domain.CreateAssetTypeInput) (domain.CreateAssetTypeOutput, error)
	GetById(ctx context.Context, input domain.GetAssetTypeByIDInput) (domain.GetAssetTypeByIDOutput, error)
	UpdateById(ctx context.Context, input domain.UpdateAssetTypeByIdInput) (domain.UpdateAssetTypeByIdOutput, error)
	DeleteById(ctx context.Context, input domain.DeleteAssetTypeByIdInput) error
}

type AssetTypeRepository interface {
	Create(ctx context.Context, input domain.CreateAssetTypeInput) (domain.AssetType, error)
	GetById(ctx context.Context, input domain.GetAssetTypeByIDInput) (domain.AssetType, error)
	UpdateById(ctx context.Context, input domain.UpdateAssetTypeByIdInput) (domain.AssetType, error)
	DeleteById(ctx context.Context, input domain.DeleteAssetTypeByIdInput) error
}
