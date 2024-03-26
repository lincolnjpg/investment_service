package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/domain"
)

type AssetIndexService interface {
	Create(ctx context.Context, input domain.CreateAssetIndexInput) (domain.CreateAssetIndexOutput, error)
	GetById(ctx context.Context, input domain.GetAssetIndexByIdInput) (domain.GetAssetIndexByIdOutput, error)
}

type AssetIndexRepository interface {
	Create(ctx context.Context, input domain.CreateAssetIndexInput) (domain.AssetIndex, error)
	GetById(ctx context.Context, input domain.GetAssetIndexByIdInput) (domain.AssetIndex, error)
}
