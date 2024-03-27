package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/domain"
)

type AssetIndexService interface {
	Create(ctx context.Context, input domain.CreateAssetIndexInput) (domain.CreateAssetIndexOutput, error)
	GetById(ctx context.Context, input domain.GetAssetIndexByIdInput) (domain.GetAssetIndexByIdOutput, error)
	UpdateById(ctx context.Context, input domain.UpdateAssetIndexByIdInput) (domain.UpdateAssetIndexByIdOutput, error)
}

type AssetIndexRepository interface {
	Create(ctx context.Context, input domain.CreateAssetIndexInput) (domain.AssetIndex, error)
	GetById(ctx context.Context, input domain.GetAssetIndexByIdInput) (domain.AssetIndex, error)
	UpdateById(ctx context.Context, input domain.UpdateAssetIndexByIdInput) (domain.AssetIndex, error)
}
