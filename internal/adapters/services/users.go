package services

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/domain"
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

func (service userService) Create(ctx context.Context, input domain.CreateUserInput) (domain.CreateUserOutput, error) {
	user, err := service.repository.Create(ctx, input)
	if err != nil {
		return domain.CreateUserOutput{}, err
	}

	return domain.CreateUserOutput{Id: user.Id}, nil
}

func (service userService) UpdateById(ctx context.Context, input domain.UpdateUserInput) (domain.UpdateUserOutput, error) {
	_, err := service.GetById(ctx, domain.GetUserByIDInput{Id: input.Id})
	if err != nil {
		return domain.UpdateUserOutput{}, err
	}

	user, err := service.repository.UpdateById(ctx, input)
	if err != nil {
		return domain.UpdateUserOutput{}, err
	}

	return domain.UpdateUserOutput{Id: user.Id}, nil
}

func (service userService) GetById(ctx context.Context, input domain.GetUserByIDInput) (domain.GetUserByIdOutput, error) {
	user, err := service.repository.GetById(ctx, input)
	if err != nil {
		return domain.GetUserByIdOutput{}, err
	}

	return domain.GetUserByIdOutput(user), nil
}

func (service userService) DeleteById(ctx context.Context, input domain.DeleteUserByIDInput) error {
	_, err := service.GetById(ctx, domain.GetUserByIDInput(input))
	if err != nil {
		return err
	}

	err = service.repository.DeleteById(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
