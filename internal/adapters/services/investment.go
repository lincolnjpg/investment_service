package services

import (
	"context"
	"encoding/json"
	"log"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/infra"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type investmentService struct {
	repo         ports.InvestmentRepository
	producer     ports.Producer
	userService  ports.UserService
	assetService ports.AssetService
}

func (s investmentService) CreateInvestment(ctx context.Context, input dtos.CreateInvestmentInput) (dtos.CreateInvestmentOutput, error) {
	_, err := s.userService.GetUserById(ctx, dtos.GetUserByIdInput{Id: input.UserId})
	if err != nil {
		return dtos.CreateInvestmentOutput{}, err
	}

	asset, err := s.assetService.GetAssetById(ctx, dtos.GetAssetByIdInput{Id: input.AssetId})
	if err != nil {
		return dtos.CreateInvestmentOutput{}, err
	}

	investment, err := s.repo.CreateInvestment(ctx, input)
	if err != nil {
		return dtos.CreateInvestmentOutput{}, err
	}

	m := infra.Message{
		InvestmentId: investment.Id,
		Ticker:       *asset.Ticker,
	}

	message, err := json.Marshal(m)
	if err != nil {
		return dtos.CreateInvestmentOutput{}, err
	}

	err = s.producer.Produce(message)
	if err != nil {
		log.Println(err)
	}

	return dtos.CreateInvestmentOutput{Id: investment.Id}, nil
}

func (s investmentService) GetInvestmentById(ctx context.Context, input dtos.GetInvestmentByIdInput) (dtos.GetInvestmentByIdOutput, error) {
	investment, err := s.repo.GetInvestmentById(ctx, input)
	if err != nil {
		return dtos.GetInvestmentByIdOutput{}, err
	}

	return dtos.GetInvestmentByIdOutput(investment), nil
}

func (s investmentService) UpdateInvestmentById(ctx context.Context, input dtos.UpdateInvestmentByIdInput) (dtos.UpdateInvestmentByIdOutput, error) {
	_, err := s.GetInvestmentById(ctx, dtos.GetInvestmentByIdInput{Id: input.Id})
	if err != nil {
		return dtos.UpdateInvestmentByIdOutput{}, err
	}

	user, err := s.repo.UpdateInvestmentById(ctx, input)
	if err != nil {
		return dtos.UpdateInvestmentByIdOutput{}, err
	}

	return dtos.UpdateInvestmentByIdOutput{Id: user.Id}, nil
}

func NewInvestmentService(repo ports.InvestmentRepository, producer ports.Producer, useuserService ports.UserService, assetsService ports.AssetService) *investmentService {
	return &investmentService{
		repo:         repo,
		producer:     producer,
		userService:  useuserService,
		assetService: assetsService,
	}
}
