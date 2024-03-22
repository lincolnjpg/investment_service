package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/domain"
)

type AssetTypeService interface {
	Create(ctx context.Context, input domain.CreateAssetTypeInput) (domain.CreateAssetTypeOutput, error)
}

type AssetTypeRepository interface {
	Create(ctx context.Context, input domain.CreateAssetTypeInput) (domain.AssetType, error)
}
