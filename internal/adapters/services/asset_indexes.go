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

func (service assetIndexService) Create(ctx context.Context, input dtos.CreateAssetIndexInput) (dtos.CreateAssetIndexOutput, error) {
	assetIndex, err := service.repository.Create(ctx, input)
	if err != nil {
		return dtos.CreateAssetIndexOutput{}, err
	}

	return dtos.CreateAssetIndexOutput{Id: assetIndex.Id}, nil
}

func (service assetIndexService) GetById(ctx context.Context, input dtos.GetAssetIndexByIdInput) (dtos.GetAssetIndexByIdOutput, error) {
	assetIndex, err := service.repository.GetById(ctx, input)
	if err != nil {
		return dtos.GetAssetIndexByIdOutput{}, err
	}

	return dtos.GetAssetIndexByIdOutput(assetIndex), nil
}

func (service assetIndexService) UpdateById(ctx context.Context, input dtos.UpdateAssetIndexByIdInput) (dtos.UpdateAssetIndexByIdOutput, error) {
	_, err := service.GetById(ctx, dtos.GetAssetIndexByIdInput{Id: input.Id})
	if err != nil {
		return dtos.UpdateAssetIndexByIdOutput{}, err
	}

	assetIndex, err := service.repository.UpdateById(ctx, input)
	if err != nil {
		return dtos.UpdateAssetIndexByIdOutput{}, err
	}

	return dtos.UpdateAssetIndexByIdOutput{Id: assetIndex.Id}, nil
}

func (service assetIndexService) DeleteById(ctx context.Context, input dtos.DeleteAssetIndexByIdInput) error {
	_, err := service.GetById(ctx, dtos.GetAssetIndexByIdInput(input))
	if err != nil {
		return err
	}

	err = service.repository.DeleteById(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
