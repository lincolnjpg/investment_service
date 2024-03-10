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
	user, err := s.repo.Create(ctx, input)
	if err != nil {
		return domain.CreateUserOutput{}, err
	}

	return domain.CreateUserOutput{Id: user.ID}, nil
}

func (s UserService) Update(ctx context.Context, input domain.UpdateUserInput) (domain.UpdateUserOutput, error) {
	user, err := s.repo.Update(ctx, input)
	if err != nil {
		return domain.UpdateUserOutput{}, err
	}

	return domain.UpdateUserOutput{ID: user.ID}, nil
}

func (s UserService) GetById(ctx context.Context, input domain.GetUserByIDInput) (domain.GetUserByIdOutput, error) {
	id, _ := uuid.Parse(input.ID)

	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return domain.GetUserByIdOutput{}, err
	}

	return domain.GetUserByIdOutput(user), nil
}

func (s UserService) DeleteById(ctx context.Context, id uuid.UUID) error {
	return nil
}
