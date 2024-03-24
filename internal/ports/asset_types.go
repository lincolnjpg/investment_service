package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/domain"
)

type AssetTypeService interface {
	Create(ctx context.Context, input domain.CreateAssetTypeInput) (domain.CreateAssetTypeOutput, error)
	GetById(ctx context.Context, input domain.GetAssetTypeByIDInput) (domain.GetAssetTypeByIDOutput, error)
}

type AssetTypeRepository interface {
	Create(ctx context.Context, input domain.CreateAssetTypeInput) (domain.AssetType, error)
	GetById(ctx context.Context, input domain.GetAssetTypeByIDInput) (domain.AssetType, error)
}
