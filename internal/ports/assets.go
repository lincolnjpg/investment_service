package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/domain"
)

type AssetService interface {
	Create(ctx context.Context, input domain.CreateAssetInput) (domain.CreateAssetOutput, error)
}

type AssetRepository interface {
	Create(ctx context.Context, input domain.CreateAssetInput) (domain.Asset, error)
}
