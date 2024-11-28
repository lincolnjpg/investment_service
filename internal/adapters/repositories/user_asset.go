package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/entities"
)

type userAssetRepository struct {
	db *pgx.Conn
}

func (r userAssetRepository) CreateUserAsset(ctx context.Context, input dtos.CreateUserAssetInput) (entities.UserAsset, error) {
	return entities.UserAsset{}, nil
}

func NewUserAssetRepository(db *pgx.Conn) userAssetRepository {
	return userAssetRepository{db: db}
}
