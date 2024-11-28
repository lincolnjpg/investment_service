package services

import (
	"context"
	"fmt"
	"log"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type userAssetService struct {
	repo                 ports.UserAssetRepository
	messageBrokerService ports.MessageBroker
}

func (s userAssetService) CreateUserAsset(ctx context.Context, input dtos.CreateUserAssetInput) (dtos.CreateUserAssetOutput, error) {
	message := fmt.Sprintf("Message for user id %s and asset id %s\n", input.UserId.String(), input.AssetId.String())
	err := s.messageBrokerService.Publish(message)
	if err != nil {
		log.Println(err)
	}
	return dtos.CreateUserAssetOutput{}, nil
}

func NewUserAssetService(repo ports.UserAssetRepository, messageBrokerService ports.MessageBroker) *userAssetService {
	return &userAssetService{
		repo:                 repo,
		messageBrokerService: messageBrokerService,
	}
}
