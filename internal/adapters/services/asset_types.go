package services

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type AssetTypeService struct {
	repo ports.AssetTypeRepository
}

func NewAssetTypeService(repo ports.AssetTypeRepository) AssetTypeService {
	return AssetTypeService{repo: repo}
}

func (s AssetTypeService) Create(ctx context.Context, input domain.CreateAssetTypeInput) (domain.CreateAssetTypeOutput, error) {
	assetType, err := s.repo.Create(ctx, input)
	if err != nil {
		return domain.CreateAssetTypeOutput{}, err
	}

	return domain.CreateAssetTypeOutput{Id: assetType.Id}, nil
}

func (s AssetTypeService) GetById(ctx context.Context, input domain.GetAssetTypeByIDInput) (domain.GetAssetTypeByIDOutput, error) {
	assetType, err := s.repo.GetById(ctx, input)
	if err != nil {
		return domain.GetAssetTypeByIDOutput{}, err
	}

	return domain.GetAssetTypeByIDOutput(assetType), nil
}

func (s AssetTypeService) UpdateById(ctx context.Context, input domain.UpdateAssetTypeByIdInput) (domain.UpdateAssetTypeByIdOutput, error) {
	_, err := s.GetById(ctx, domain.GetAssetTypeByIDInput{Id: input.Id})
	if err != nil {
		return domain.UpdateAssetTypeByIdOutput{}, err
	}

	assetType, err := s.repo.UpdateById(ctx, input)
	if err != nil {
		return domain.UpdateAssetTypeByIdOutput{}, err
	}

	return domain.UpdateAssetTypeByIdOutput{Id: assetType.Id}, nil
}
