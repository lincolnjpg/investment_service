package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/entities"
)

type UserService interface {
	Create(ctx context.Context, input dtos.CreateUserInput) (dtos.CreateUserOutput, error)
	UpdateById(ctx context.Context, input dtos.UpdateUserInput) (dtos.UpdateUserOutput, error)
	GetById(ctx context.Context, input dtos.GetUserByIdInput) (dtos.GetUserByIdOutput, error)
	DeleteById(ctx context.Context, input dtos.DeleteUserByIDInput) error
}

type UserRepository interface {
	Create(ctx context.Context, input dtos.CreateUserInput) (entities.User, error)
	UpdateById(ctx context.Context, input dtos.UpdateUserInput) (entities.User, error)
	GetById(ctx context.Context, input dtos.GetUserByIdInput) (entities.User, error)
	DeleteById(ctx context.Context, input dtos.DeleteUserByIDInput) error
}
