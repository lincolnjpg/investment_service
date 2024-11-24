package services

import (
	"context"
	"net/http"

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

func (service assetService) Create(ctx context.Context, input dtos.CreateAssetInput) (dtos.CreateAssetOutput, error) {
	if input.AssetIndexId != nil {
		assetIndex, err := service.assetIndexesService.GetById(ctx, dtos.GetAssetIndexByIdInput{Id: *input.AssetIndexId})
		if err != nil {
			return dtos.CreateAssetOutput{}, err
		}

		if assetIndex.Acronym == enum.CdiAcronym && input.Rentability > maxCdiRentability {
			return dtos.CreateAssetOutput{}, customerror.NewAPIError("rentability of an investment indexed by CDI can not be greater than 150%", http.StatusBadRequest)
		}
	}

	asset, err := service.repository.Create(ctx, input)
	if err != nil {
		return dtos.CreateAssetOutput{}, err
	}

	return dtos.CreateAssetOutput{Id: asset.Id}, nil
}

func (service assetService) GetById(ctx context.Context, input dtos.GetAssetByIdInput) (dtos.GetAssetByIdOutput, error) {
	asset, err := service.repository.GetById(ctx, input)
	if err != nil {
		return dtos.GetAssetByIdOutput{}, err
	}

	return dtos.GetAssetByIdOutput(asset), nil
}

func (service assetService) UpdateById(ctx context.Context, input dtos.UpdateAssetByIdInput) (dtos.UpdateAssetByIdOutput, error) {
	_, err := service.GetById(ctx, dtos.GetAssetByIdInput{Id: input.Id})
	if err != nil {
		return dtos.UpdateAssetByIdOutput{}, err
	}

	asset, err := service.repository.UpdateById(ctx, input)
	if err != nil {
		return dtos.UpdateAssetByIdOutput{}, err
	}

	return dtos.UpdateAssetByIdOutput{Id: asset.Id}, nil
}

func (service assetService) DeleteById(ctx context.Context, input dtos.DeleteAssetByIdInput) error {
	_, err := service.GetById(ctx, dtos.GetAssetByIdInput(input))
	if err != nil {
		return err
	}

	err = service.repository.DeleteById(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
