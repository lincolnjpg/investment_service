package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/entities"
)

type UserAssetService interface {
	CreateUserAsset(ctx context.Context, input dtos.CreateUserAssetInput) (dtos.CreateUserAssetOutput, error)
}

type UserAssetRepository interface {
	CreateUserAsset(ctx context.Context, input dtos.CreateUserAssetInput) (entities.UserAsset, error)
}
