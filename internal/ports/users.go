package ports

import (
	"context"

	"github.com/lincolnjpg/investment_service/internal/dtos"
	"github.com/lincolnjpg/investment_service/internal/entities"
)

type UserService interface {
	CreateUser(ctx context.Context, input dtos.CreateUserInput) (dtos.CreateUserOutput, error)
	UpdateUserById(ctx context.Context, input dtos.UpdateUserInput) (dtos.UpdateUserOutput, error)
	GetUserById(ctx context.Context, input dtos.GetUserByIdInput) (dtos.GetUserByIdOutput, error)
	DeleteUserById(ctx context.Context, input dtos.DeleteUserByIdInput) error
}

type UserRepository interface {
	CreateUser(ctx context.Context, input dtos.CreateUserInput) (entities.User, error)
	UpdateUserById(ctx context.Context, input dtos.UpdateUserInput) (entities.User, error)
	GetUserById(ctx context.Context, input dtos.GetUserByIdInput) (entities.User, error)
	DeleteUserById(ctx context.Context, input dtos.DeleteUserByIdInput) error
}
