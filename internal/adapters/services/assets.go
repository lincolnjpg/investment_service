package services

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type AssetService struct {
	repository ports.AssetRepository
}

func NewAssetService(repository ports.AssetRepository) AssetService {
	return AssetService{repository: repository}
}

func (service AssetService) Create(ctx context.Context, input domain.CreateAssetInput) (domain.CreateAssetOutput, error) {
	asset, err := service.repository.Create(ctx, input)
	if err != nil {
		return domain.CreateAssetOutput{}, err
	}

	return domain.CreateAssetOutput{Id: asset.Id}, nil
}

func (service AssetService) GetById(ctx context.Context, input domain.GetAssetByIdInput) (domain.GetAssetByIdOutput, error) {
	asset, err := service.repository.GetById(ctx, input)
	if err != nil {
		return domain.GetAssetByIdOutput{}, err
	}

	return domain.GetAssetByIdOutput(asset), nil
}