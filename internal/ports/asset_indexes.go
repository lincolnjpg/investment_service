package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/domain"
)

type AssetIndexService interface {
	Create(ctx context.Context, input domain.CreateAssetIndexInput) (domain.CreateAssetIndexOutput, error)
}

type AssetIndexRepository interface {
	Create(ctx context.Context, input domain.CreateAssetIndexInput) (domain.AssetIndex, error)
}
