package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/domain"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s UserService) Create(ctx context.Context, input domain.CreateUserInput) (domain.CreateUserOutput, error) {
	return domain.CreateUserOutput{}, nil
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
