package services

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type AssetTypeService struct {
	repo ports.AssetTypeRepository
}

func NewAssetTypeService(repo ports.AssetTypeRepository) AssetTypeService {
	return AssetTypeService{repo: repo}
}

func (s AssetTypeService) Create(ctx context.Context, input domain.CreateAssetTypeInput) (domain.CreateAssetTypeOutput, error) {
	assetType, err := s.repo.Create(ctx, input)
	if err != nil {
		return domain.CreateAssetTypeOutput{}, err
	}

	return domain.CreateAssetTypeOutput{ID: assetType.ID}, nil
}
