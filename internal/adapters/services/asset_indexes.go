package services

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type AssetIndexService struct {
	repository ports.AssetIndexRepository
}

func NewAssetIndexService(repository ports.AssetIndexRepository) AssetIndexService {
	return AssetIndexService{repository: repository}
}

func (service AssetIndexService) Create(ctx context.Context, input domain.CreateAssetIndexInput) (domain.CreateAssetIndexOutput, error) {
	assetIndex, err := service.repository.Create(ctx, input)
	if err != nil {
		return domain.CreateAssetIndexOutput{}, err
	}

	return domain.CreateAssetIndexOutput{Id: assetIndex.Id}, nil
}
