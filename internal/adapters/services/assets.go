package services

import (
	"context"
	"net/http"

	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/infra"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type AssetService struct {
	repository          ports.AssetRepository
	assetTypesService   ports.AssetTypeService
	assetIndexesService ports.AssetIndexService
}

func NewAssetService(repository ports.AssetRepository, assetTypesService ports.AssetTypeService, assetIndexesService ports.AssetIndexService) AssetService {
	return AssetService{
		repository:          repository,
		assetTypesService:   assetTypesService,
		assetIndexesService: assetIndexesService,
	}
}

func (service AssetService) Create(ctx context.Context, input domain.CreateAssetInput) (domain.CreateAssetOutput, error) {
	_, err := service.assetTypesService.GetById(ctx, domain.GetAssetTypeByIDInput{Id: input.AssetTypeId})
	if err != nil {
		return domain.CreateAssetOutput{}, err
	}

	assetIndex, err := service.assetIndexesService.GetById(ctx, domain.GetAssetIndexByIdInput{Id: *input.AssetIndexId})
	if err != nil {
		return domain.CreateAssetOutput{}, err
	}

	if assetIndex.Acronym == domain.CDI_ACRONYM && input.Rentability > domain.MAX_CDI_RENTABILITY {
		return domain.CreateAssetOutput{}, infra.NewAPIError("rentability of an investment indexed by CDI can not be greater than 150%", http.StatusBadRequest)
	}

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
