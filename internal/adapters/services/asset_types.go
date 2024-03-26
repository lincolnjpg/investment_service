package services

import (
	"context"
	"net/http"

	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/infra"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type AssetTypeService struct {
	repository ports.AssetTypeRepository
}

func NewAssetTypeService(repository ports.AssetTypeRepository) AssetTypeService {
	return AssetTypeService{repository: repository}
}

func (service AssetTypeService) Create(ctx context.Context, input domain.CreateAssetTypeInput) (domain.CreateAssetTypeOutput, error) {
	if input.Class == domain.VARIABLE_INCOME && input.IndexId != nil {
		return domain.CreateAssetTypeOutput{}, infra.NewAPIError("this asset class can not be indexed", http.StatusBadRequest)
	}

	assetType, err := service.repository.Create(ctx, input)
	if err != nil {
		return domain.CreateAssetTypeOutput{}, err
	}

	return domain.CreateAssetTypeOutput{Id: assetType.Id}, nil
}

func (service AssetTypeService) GetById(ctx context.Context, input domain.GetAssetTypeByIDInput) (domain.GetAssetTypeByIDOutput, error) {
	assetType, err := service.repository.GetById(ctx, input)
	if err != nil {
		return domain.GetAssetTypeByIDOutput{}, err
	}

	return domain.GetAssetTypeByIDOutput(assetType), nil
}

func (service AssetTypeService) UpdateById(ctx context.Context, input domain.UpdateAssetTypeByIdInput) (domain.UpdateAssetTypeByIdOutput, error) {
	if input.Class == domain.VARIABLE_INCOME && input.IndexId != nil {
		return domain.UpdateAssetTypeByIdOutput{}, infra.NewAPIError("this asset class can not be indexed", http.StatusBadRequest)
	}

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

func (service AssetTypeService) DeleteById(ctx context.Context, input domain.DeleteAssetTypeByIdInput) error {
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
