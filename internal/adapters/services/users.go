package services

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/ports"
)

type userService struct {
	repository ports.UserRepository
}

func NewUserService(repository ports.UserRepository) userService {
	return userService{
		repository: repository,
	}
}

func (service userService) Create(ctx context.Context, input dtos.CreateUserInput) (dtos.CreateUserOutput, error) {
	user, err := service.repository.Create(ctx, input)
	if err != nil {
		return dtos.CreateUserOutput{}, err
	}

	return dtos.CreateUserOutput{Id: user.Id}, nil
}

func (service userService) UpdateById(ctx context.Context, input dtos.UpdateUserInput) (dtos.UpdateUserOutput, error) {
	_, err := service.GetById(ctx, dtos.GetUserByIdInput{Id: input.Id})
	if err != nil {
		return dtos.UpdateUserOutput{}, err
	}

	user, err := service.repository.UpdateById(ctx, input)
	if err != nil {
		return dtos.UpdateUserOutput{}, err
	}

	return dtos.UpdateUserOutput{Id: user.Id}, nil
}

func (service userService) GetById(ctx context.Context, input dtos.GetUserByIdInput) (dtos.GetUserByIdOutput, error) {
	user, err := service.repository.GetById(ctx, input)
	if err != nil {
		return dtos.GetUserByIdOutput{}, err
	}

	return dtos.GetUserByIdOutput(user), nil
}

func (service userService) DeleteById(ctx context.Context, input dtos.DeleteUserByIdInput) error {
	_, err := service.GetById(ctx, dtos.GetUserByIdInput(input))
	if err != nil {
		return err
	}

	err = service.repository.DeleteById(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
