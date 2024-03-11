package services

import (
	"context"

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
	_, err := s.GetById(ctx, domain.GetUserByIDInput{ID: input.ID})
	if err != nil {
		return domain.UpdateUserOutput{}, err
	}

	user, err := s.repo.Update(ctx, input)
	if err != nil {
		return domain.UpdateUserOutput{}, err
	}

	return domain.UpdateUserOutput{ID: user.ID}, nil
}

func (s UserService) GetById(ctx context.Context, input domain.GetUserByIDInput) (domain.GetUserByIdOutput, error) {
	user, err := s.repo.GetById(ctx, input)
	if err != nil {
		return domain.GetUserByIdOutput{}, err
	}

	return domain.GetUserByIdOutput(user), nil
}

func (s UserService) DeleteById(ctx context.Context, input domain.DeleteUserByIDInput) error {
	_, err := s.GetById(ctx, domain.GetUserByIDInput(input))
	if err != nil {
		return err
	}

	err = s.repo.DeleteById(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
