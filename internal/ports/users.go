package ports

import (
	"github.com/google/uuid"
	"github.com/lincolnjpg/investment_service/internal/domain"
)

type UserService interface {
	Create(input domain.CreateUserInput) (domain.CreateUserOutput, error)
	Update(input domain.UpdateUserInput) (domain.UpdateUserOutput, error)
	GetById(id uuid.UUID) (domain.GetUserByIdOutput, error)
	DeleteById(id uuid.UUID) error
}

type UserRepository interface {
	Create(input domain.CreateUserInput) (domain.User, error)
	Update(input domain.UpdateUserInput) (domain.User, error)
	GetById(id uuid.UUID) (domain.User, error)
	DeleteById(id uuid.UUID) error
}
