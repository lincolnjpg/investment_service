package services

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type UserService struct {
	ctx    context.Context
	logger *slog.Logger
	repo   ports.UserRepository
}

func NewUserService(ctx context.Context, logger *slog.Logger, repo ports.UserRepository) UserService {
	return UserService{
		ctx:    ctx,
		logger: logger,
		repo:   repo,
	}
}

func (s UserService) Create(input domain.CreateUserInput) (domain.CreateUserOutput, error) {
	user, err := s.repo.Create(input)
	if err != nil {
		return domain.CreateUserOutput{}, err
	}

	return domain.CreateUserOutput{Id: user.Id}, nil
}

func (s UserService) Update(input domain.UpdateUserInput) (domain.UpdateUserOutput, error) {
	return domain.UpdateUserOutput{}, nil
}

func (s UserService) GetById(id uuid.UUID) (domain.GetUserByIdOutput, error) {
	return domain.GetUserByIdOutput{}, nil
}

func (s UserService) DeleteById(id uuid.UUID) error {
	return nil
}
