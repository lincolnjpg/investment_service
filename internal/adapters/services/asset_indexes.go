package services

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type assetIndexService struct {
	repository ports.AssetIndexRepository
}

func NewAssetIndexService(repository ports.AssetIndexRepository) assetIndexService {
	return assetIndexService{repository: repository}
}

func (service assetIndexService) Create(ctx context.Context, input domain.CreateAssetIndexInput) (domain.CreateAssetIndexOutput, error) {
	assetIndex, err := service.repository.Create(ctx, input)
	if err != nil {
		return domain.CreateAssetIndexOutput{}, err
	}

	return domain.CreateAssetIndexOutput{Id: assetIndex.Id}, nil
}

func (service assetIndexService) GetById(ctx context.Context, input domain.GetAssetIndexByIdInput) (domain.GetAssetIndexByIdOutput, error) {
	assetIndex, err := service.repository.GetById(ctx, input)
	if err != nil {
		return domain.GetAssetIndexByIdOutput{}, err
	}

	return domain.GetAssetIndexByIdOutput(assetIndex), nil
}

func (service assetIndexService) UpdateById(ctx context.Context, input domain.UpdateAssetIndexByIdInput) (domain.UpdateAssetIndexByIdOutput, error) {
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

func (service assetIndexService) DeleteById(ctx context.Context, input domain.DeleteAssetIndexByIdInput) error {
	_, err := service.GetById(ctx, domain.GetAssetIndexByIdInput(input))
	if err != nil {
		return err
	}

	err = service.repository.DeleteById(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
