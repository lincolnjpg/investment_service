package services

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type assetTypeService struct {
	repository ports.AssetTypeRepository
}

func NewAssetTypeService(repository ports.AssetTypeRepository) assetTypeService {
	return assetTypeService{repository: repository}
}

func (service assetTypeService) Create(ctx context.Context, input dtos.CreateAssetTypeInput) (dtos.CreateAssetTypeOutput, error) {
	assetType, err := service.repository.Create(ctx, input)
	if err != nil {
		return dtos.CreateAssetTypeOutput{}, err
	}

	return dtos.CreateAssetTypeOutput{Id: assetType.Id}, nil
}

func (service assetTypeService) GetById(ctx context.Context, input dtos.GetAssetTypeByIDInput) (dtos.GetAssetTypeByIDOutput, error) {
	assetType, err := service.repository.GetById(ctx, input)
	if err != nil {
		return dtos.GetAssetTypeByIDOutput{}, err
	}

	return dtos.GetAssetTypeByIDOutput(assetType), nil
}

func (service assetTypeService) UpdateById(ctx context.Context, input dtos.UpdateAssetTypeByIdInput) (dtos.UpdateAssetTypeByIdOutput, error) {
	_, err := service.GetById(ctx, dtos.GetAssetTypeByIDInput{Id: input.Id})
	if err != nil {
		return dtos.UpdateAssetTypeByIdOutput{}, err
	}

	assetType, err := service.repository.UpdateById(ctx, input)
	if err != nil {
		return dtos.UpdateAssetTypeByIdOutput{}, err
	}

	return dtos.UpdateAssetTypeByIdOutput{Id: assetType.Id}, nil
}

func (service assetTypeService) DeleteById(ctx context.Context, input dtos.DeleteAssetTypeByIdInput) error {
	_, err := service.GetById(ctx, dtos.GetAssetTypeByIDInput(input))
	if err != nil {
		return err
	}

	err = service.repository.DeleteById(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
