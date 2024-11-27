package services

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/enum"
	customerror "github.com/lincolnjpg/investment_service/internal/error"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

const maxCdiRentability float64 = 150.0

type assetService struct {
	repository          ports.AssetRepository
	assetIndexesService ports.AssetIndexService
}

func NewAssetService(repository ports.AssetRepository, assetIndexesService ports.AssetIndexService) assetService {
	return assetService{
		repository:          repository,
		assetIndexesService: assetIndexesService,
	}
}

func (service assetService) CreateAsset(ctx context.Context, input dtos.CreateAssetInput) (dtos.CreateAssetOutput, error) {
	if input.AssetIndexId != uuid.Nil {
		assetIndex, err := service.assetIndexesService.GetAssetIndexById(ctx, dtos.GetAssetIndexByIdInput{Id: input.AssetIndexId})
		if err != nil {
			return dtos.CreateAssetOutput{}, err
		}

		if assetIndex.Acronym == enum.CdiAcronym && input.Rentability > maxCdiRentability {
			return dtos.CreateAssetOutput{}, customerror.NewAPIError("rentability of an investment indexed by CDI can not be greater than 150%", http.StatusBadRequest)
		}
	}

	asset, err := service.repository.CreateAsset(ctx, input)
	if err != nil {
		return dtos.CreateAssetOutput{}, err
	}

	return dtos.CreateAssetOutput{Id: asset.Id}, nil
}

func (service assetService) GetAssetById(ctx context.Context, input dtos.GetAssetByIdInput) (dtos.GetAssetByIdOutput, error) {
	asset, err := service.repository.GetAssetById(ctx, input)
	if err != nil {
		return dtos.GetAssetByIdOutput{}, err
	}

	return dtos.GetAssetByIdOutput(asset), nil
}

func (service assetService) UpdateAssetById(ctx context.Context, input dtos.UpdateAssetByIdInput) (dtos.UpdateAssetByIdOutput, error) {
	_, err := service.GetAssetById(ctx, dtos.GetAssetByIdInput{Id: input.Id})
	if err != nil {
		return dtos.UpdateAssetByIdOutput{}, err
	}

	asset, err := service.repository.UpdateAssetById(ctx, input)
	if err != nil {
		return dtos.UpdateAssetByIdOutput{}, err
	}

	return dtos.UpdateAssetByIdOutput{Id: asset.Id}, nil
}

func (service assetService) DeleteAssetById(ctx context.Context, input dtos.DeleteAssetByIdInput) error {
	_, err := service.GetAssetById(ctx, dtos.GetAssetByIdInput(input))
	if err != nil {
		return err
	}

	err = service.repository.DeleteAssetById(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
