package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/domain"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s UserService) Create(input domain.CreateUserInput) (domain.CreateUserOutput, error) {
	fmt.Println("REPO OK")
	return domain.CreateUserOutput{}, nil
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
