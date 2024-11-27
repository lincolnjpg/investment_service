package services

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type assetIndexService struct {
	repository ports.AssetIndexRepository
}

func NewAssetIndexService(repository ports.AssetIndexRepository) assetIndexService {
	return assetIndexService{repository: repository}
}

func (service assetIndexService) CreateAssetIndex(ctx context.Context, input dtos.CreateAssetIndexInput) (dtos.CreateAssetIndexOutput, error) {
	assetIndex, err := service.repository.CreateAssetIndex(ctx, input)
	if err != nil {
		return dtos.CreateAssetIndexOutput{}, err
	}

	return dtos.CreateAssetIndexOutput{Id: assetIndex.Id}, nil
}

func (service assetIndexService) GetAssetIndexById(ctx context.Context, input dtos.GetAssetIndexByIdInput) (dtos.GetAssetIndexByIdOutput, error) {
	assetIndex, err := service.repository.GetAssetIndexById(ctx, input)
	if err != nil {
		return dtos.GetAssetIndexByIdOutput{}, err
	}

	return dtos.GetAssetIndexByIdOutput(assetIndex), nil
}

func (service assetIndexService) UpdateAssetIndexById(ctx context.Context, input dtos.UpdateAssetIndexByIdInput) (dtos.UpdateAssetIndexByIdOutput, error) {
	_, err := service.GetAssetIndexById(ctx, dtos.GetAssetIndexByIdInput{Id: input.Id})
	if err != nil {
		return dtos.UpdateAssetIndexByIdOutput{}, err
	}

	assetIndex, err := service.repository.UpdateAssetIndexById(ctx, input)
	if err != nil {
		return dtos.UpdateAssetIndexByIdOutput{}, err
	}

	return dtos.UpdateAssetIndexByIdOutput{Id: assetIndex.Id}, nil
}

func (service assetIndexService) DeleteAssetIndexById(ctx context.Context, input dtos.DeleteAssetIndexByIdInput) error {
	_, err := service.GetAssetIndexById(ctx, dtos.GetAssetIndexByIdInput(input))
	if err != nil {
		return err
	}

	err = service.repository.DeleteAssetIndexById(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
