package services

import (
	"context"

	"github.com/go-chi/httplog/v2"
	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type UserService struct {
	logger *httplog.Logger
	repo   ports.UserRepository
}

func NewUserService(logger *httplog.Logger, repo ports.UserRepository) UserService {
	return UserService{
		logger: logger,
		repo:   repo,
	}
}

func (s UserService) Create(ctx context.Context, input domain.CreateUserInput) (domain.CreateUserOutput, error) {
	user, err := s.repo.Create(ctx, input)
	if err != nil {
		return domain.CreateUserOutput{}, err
	}

	return domain.CreateUserOutput{Id: user.Id}, nil
}

func (s UserService) Update(ctx context.Context, input domain.UpdateUserInput) (domain.UpdateUserOutput, error) {
	return domain.UpdateUserOutput{}, nil
}

func (s UserService) GetById(ctx context.Context, id uuid.UUID) (domain.GetUserByIdOutput, error) {
	return domain.GetUserByIdOutput{}, nil
}

func (s UserService) DeleteById(ctx context.Context, id uuid.UUID) error {
	return nil
}
