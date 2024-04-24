package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/domain"
)

type AssetService interface {
	Create(ctx context.Context, input domain.CreateAssetInput) (domain.CreateAssetOutput, error)
	GetById(ctx context.Context, input domain.GetAssetByIdInput) (domain.GetAssetByIdOutput, error)
	UpdateById(ctx context.Context, input domain.UpdateAssetByIdInput) (domain.UpdateAssetByIdOutput, error)
}

type AssetRepository interface {
	Create(ctx context.Context, input domain.CreateAssetInput) (domain.Asset, error)
	GetById(ctx context.Context, input domain.GetAssetByIdInput) (domain.Asset, error)
	UpdateById(ctx context.Context, input domain.UpdateAssetByIdInput) (domain.Asset, error)
}
