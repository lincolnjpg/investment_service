package services

import (
	"context"
	"encoding/json"
	"log"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/infra"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type userAssetService struct {
	repo         ports.UserAssetRepository
	producer     ports.Producer
	userService  ports.UserService
	assetService ports.AssetService
}

func (s userAssetService) CreateUserAsset(ctx context.Context, input dtos.CreateUserAssetInput) (dtos.CreateUserAssetOutput, error) {
	_, err := s.userService.GetUserById(ctx, dtos.GetUserByIdInput{Id: input.UserId})
	if err != nil {
		return dtos.CreateUserAssetOutput{}, err
	}

	asset, err := s.assetService.GetAssetById(ctx, dtos.GetAssetByIdInput{Id: input.AssetId})
	if err != nil {
		return dtos.CreateUserAssetOutput{}, err
	}

	userAsset, err := s.repo.CreateUserAsset(ctx, input)
	if err != nil {
		return dtos.CreateUserAssetOutput{}, err
	}

	m := infra.Message{
		UserAssetId: userAsset.Id,
		Ticker:      *asset.Ticker,
	}

	message, err := json.Marshal(m)
	if err != nil {
		return dtos.CreateUserAssetOutput{}, err
	}

	err = s.producer.Produce(message)
	if err != nil {
		log.Println(err)
	}

	return dtos.CreateUserAssetOutput{Id: userAsset.Id}, nil
}

func NewUserAssetService(repo ports.UserAssetRepository, producer ports.Producer, useuserService ports.UserService, assetsService ports.AssetService) *userAssetService {
	return &userAssetService{
		repo:         repo,
		producer:     producer,
		userService:  useuserService,
		assetService: assetsService,
	}
}
