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

func (service AssetIndexService) GetById(ctx context.Context, input domain.GetAssetIndexByIdInput) (domain.GetAssetIndexByIdOutput, error) {
	assetIndex, err := service.repository.GetById(ctx, input)
	if err != nil {
		return domain.GetAssetIndexByIdOutput{}, err
	}

	return domain.GetAssetIndexByIdOutput(assetIndex), nil
}

func (service AssetIndexService) UpdateById(ctx context.Context, input domain.UpdateAssetIndexByIdInput) (domain.UpdateAssetIndexByIdOutput, error) {
	_, err := service.GetById(ctx, domain.GetAssetIndexByIdInput{Id: input.Id})
	if err != nil {
		return domain.UpdateAssetIndexByIdOutput{}, err
	}

	assetIndex, err := service.repository.UpdateById(ctx, input)
	if err != nil {
		return domain.UpdateAssetIndexByIdOutput{}, err
	}

	return domain.UpdateAssetIndexByIdOutput{Id: assetIndex.Id}, nil
}
