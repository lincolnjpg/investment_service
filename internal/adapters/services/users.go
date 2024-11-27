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

func (service userService) CreateUser(ctx context.Context, input dtos.CreateUserInput) (dtos.CreateUserOutput, error) {
	user, err := service.repository.CreateUser(ctx, input)
	if err != nil {
		return dtos.CreateUserOutput{}, err
	}

	return dtos.CreateUserOutput{Id: user.Id}, nil
}

func (service userService) UpdateUserById(ctx context.Context, input dtos.UpdateUserInput) (dtos.UpdateUserOutput, error) {
	_, err := service.GetUserById(ctx, dtos.GetUserByIdInput{Id: input.Id})
	if err != nil {
		return dtos.UpdateUserOutput{}, err
	}

	user, err := service.repository.UpdateUserById(ctx, input)
	if err != nil {
		return dtos.UpdateUserOutput{}, err
	}

	return dtos.UpdateUserOutput{Id: user.Id}, nil
}

func (service userService) GetUserById(ctx context.Context, input dtos.GetUserByIdInput) (dtos.GetUserByIdOutput, error) {
	user, err := service.repository.GetUserById(ctx, input)
	if err != nil {
		return dtos.GetUserByIdOutput{}, err
	}

	return dtos.GetUserByIdOutput(user), nil
}

func (service userService) DeleteUserById(ctx context.Context, input dtos.DeleteUserByIdInput) error {
	_, err := service.GetUserById(ctx, dtos.GetUserByIdInput(input))
	if err != nil {
		return err
	}

	err = service.repository.DeleteUserById(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
