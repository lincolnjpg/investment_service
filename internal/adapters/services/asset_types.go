package services

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type assetTypeService struct {
	repository ports.AssetTypeRepository
}

func NewAssetTypeService(repository ports.AssetTypeRepository) assetTypeService {
	return assetTypeService{repository: repository}
}

func (service assetTypeService) Create(ctx context.Context, input domain.CreateAssetTypeInput) (domain.CreateAssetTypeOutput, error) {
	assetType, err := service.repository.Create(ctx, input)
	if err != nil {
		return domain.CreateAssetTypeOutput{}, err
	}

	return domain.CreateAssetTypeOutput{Id: assetType.Id}, nil
}

func (service assetTypeService) GetById(ctx context.Context, input domain.GetAssetTypeByIDInput) (domain.GetAssetTypeByIDOutput, error) {
	assetType, err := service.repository.GetById(ctx, input)
	if err != nil {
		return domain.GetAssetTypeByIDOutput{}, err
	}

	return domain.GetAssetTypeByIDOutput(assetType), nil
}

func (service assetTypeService) UpdateById(ctx context.Context, input domain.UpdateAssetTypeByIdInput) (domain.UpdateAssetTypeByIdOutput, error) {
	_, err := service.GetById(ctx, domain.GetAssetTypeByIDInput{Id: input.Id})
	if err != nil {
		return domain.UpdateAssetTypeByIdOutput{}, err
	}

	assetType, err := service.repository.UpdateById(ctx, input)
	if err != nil {
		return domain.UpdateAssetTypeByIdOutput{}, err
	}

	return domain.UpdateAssetTypeByIdOutput{Id: assetType.Id}, nil
}

func (service assetTypeService) DeleteById(ctx context.Context, input domain.DeleteAssetTypeByIdInput) error {
	_, err := service.GetById(ctx, domain.GetAssetTypeByIDInput(input))
	if err != nil {
		return err
	}

	err = service.repository.DeleteById(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
