package services

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type userAssetService struct {
	repo                 ports.UserAssetRepository
	messageBrokerService ports.MessageBroker
	userService          ports.UserService
	assetService         ports.AssetService
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

	m := struct {
		AssetId uuid.UUID `json:"asset_id,omitempty"`
		Ticker  string    `json:"ticker,omitempty"`
	}{
		AssetId: userAsset.AssetId,
		Ticker:  *asset.Ticker,
	}
	message, err := json.Marshal(m)
	if err != nil {
		return dtos.CreateUserAssetOutput{}, err
	}

	err = s.messageBrokerService.Publish(message)
	if err != nil {
		log.Println(err)
	}

	return dtos.CreateUserAssetOutput{Id: userAsset.Id}, nil
}

func NewUserAssetService(repo ports.UserAssetRepository, messageBrokerService ports.MessageBroker, useuserService ports.UserService, assetsService ports.AssetService) *userAssetService {
	return &userAssetService{
		repo:                 repo,
		messageBrokerService: messageBrokerService,
		userService:          useuserService,
		assetService:         assetsService,
	}
}
